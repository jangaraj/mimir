// SPDX-License-Identifier: AGPL-3.0-only
// Provenance-includes-location: https://github.com/cortexproject/cortex/blob/master/pkg/alertmanager/alertmanager_client.go
// Provenance-includes-license: Apache-2.0
// Provenance-includes-copyright: The Cortex Authors.

package alertmanager

import (
	"flag"
	"time"

	"github.com/go-kit/log"
	"github.com/grafana/dskit/crypto/tls"
	"github.com/grafana/dskit/grpcclient"
	"github.com/grafana/dskit/ring/client"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"

	"github.com/grafana/mimir/pkg/alertmanager/alertmanagerpb"
)

// ClientsPool is the interface used to get the client from the pool for a specified address.
type ClientsPool interface {
	// GetClientFor returns the alertmanager client for the given address.
	GetClientFor(addr string) (Client, error)
}

// Client is the interface that should be implemented by any client used to read/write data to an alertmanager via GRPC.
type Client interface {
	alertmanagerpb.AlertmanagerClient

	// RemoteAddress returns the address of the remote alertmanager and is used to uniquely
	// identify an alertmanager instance.
	RemoteAddress() string
}

// ClientConfig is the configuration struct for the alertmanager client.
type ClientConfig struct {
	RemoteTimeout time.Duration    `yaml:"remote_timeout"`
	TLSEnabled    bool             `yaml:"tls_enabled" category:"advanced"`
	TLS           tls.ClientConfig `yaml:",inline"`
}

// RegisterFlagsWithPrefix registers flags with prefix.
func (cfg *ClientConfig) RegisterFlagsWithPrefix(prefix string, f *flag.FlagSet) {
	f.BoolVar(&cfg.TLSEnabled, prefix+".tls-enabled", cfg.TLSEnabled, "Enable TLS in the GRPC client. This flag needs to be enabled when any other TLS flag is set. If set to false, insecure connection to gRPC server will be used.")
	f.DurationVar(&cfg.RemoteTimeout, prefix+".remote-timeout", 2*time.Second, "Timeout for downstream alertmanagers.")
	cfg.TLS.RegisterFlagsWithPrefix(prefix, f)
}

// alertmanagerClientsPool is a pool of alertmanager clients.
type alertmanagerClientsPool struct {
	pool *client.Pool
}

// newAlertmanagerClientsPool creates a new pool of alertmanager clients.
func newAlertmanagerClientsPool(discovery client.PoolServiceDiscovery, amClientCfg ClientConfig, logger log.Logger, reg prometheus.Registerer) ClientsPool {
	// We prefer sensible defaults instead of exposing further config options.
	grpcCfg := grpcclient.Config{
		MaxRecvMsgSize:      16 * 1024 * 1024, // 16MiB.
		MaxSendMsgSize:      4 * 1024 * 1024,  // 4MiB.
		GRPCCompression:     "",               // No compression.
		RateLimit:           0,                // No rate limit.
		RateLimitBurst:      0,                // No burst of rate limit.
		BackoffOnRatelimits: false,            // No backoffs for rate limiting.
		TLSEnabled:          amClientCfg.TLSEnabled,
		TLS:                 amClientCfg.TLS,
	}

	requestDuration := promauto.With(reg).NewHistogramVec(prometheus.HistogramOpts{
		Name: "cortex_alertmanager_distributor_client_request_duration_seconds",
		Help: "Time spent executing requests from an alertmanager to another alertmanager.",
		// Buckets
		// 1 0.008s  (8ms)
		// 2 0.032s  (32ms)
		// 3 0.128s  (128ms)
		// 4 0.512s  (512ms)
		// 5 2.048s  (2048ms)
		// 6 8.192s  (8192ms)
		// 7 32.768s (32768ms)
		// 8 +Inf
		Buckets: prometheus.ExponentialBuckets(0.008, 4, 7),
	}, []string{"operation", "status_code"})

	factory := func(addr string) (client.PoolClient, error) {
		return dialAlertmanagerClient(grpcCfg, addr, requestDuration)
	}

	poolCfg := client.PoolConfig{
		CheckInterval:      time.Minute,
		HealthCheckEnabled: true,
		HealthCheckTimeout: 10 * time.Second,
	}

	clientsCount := promauto.With(reg).NewGauge(prometheus.GaugeOpts{
		Namespace: "cortex",
		Name:      "alertmanager_distributor_clients",
		Help:      "The current number of alertmanager distributor clients in the pool.",
	})

	return &alertmanagerClientsPool{pool: client.NewPool("alertmanager", poolCfg, discovery, factory, clientsCount, logger)}
}

// GetClientFor returns the alertmanager client for the specified address, creating it if one did not already exist.
func (f *alertmanagerClientsPool) GetClientFor(addr string) (Client, error) {
	c, err := f.pool.GetClientFor(addr)
	if err != nil {
		return nil, err
	}
	return c.(Client), nil
}

// dialAlertmanagerClient establishes a GRPC connection to an alertmanager that is aware of the the health of the server
// and collects observations of request durations.
func dialAlertmanagerClient(cfg grpcclient.Config, addr string, requestDuration *prometheus.HistogramVec) (*alertmanagerClient, error) {
	opts, err := cfg.DialOption(grpcclient.Instrument(requestDuration))
	if err != nil {
		return nil, err
	}
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to dial alertmanager %s", addr)
	}

	return &alertmanagerClient{
		AlertmanagerClient: alertmanagerpb.NewAlertmanagerClient(conn),
		HealthClient:       grpc_health_v1.NewHealthClient(conn),
		conn:               conn,
	}, nil
}

// alertmanagerClient is a GRPC client of an alertmanager.
type alertmanagerClient struct {
	alertmanagerpb.AlertmanagerClient
	grpc_health_v1.HealthClient
	conn *grpc.ClientConn
}

// Close closes the client's GRPC connection.
func (c *alertmanagerClient) Close() error {
	return c.conn.Close()
}

// String implements the Stringer interface.
// It returns the remote address of the alertmanager server which is unique for each client.
func (c *alertmanagerClient) String() string {
	return c.RemoteAddress()
}

// RemoteAddress implements the Client interface.
// It returns the address of the alertmanager server which is unique for each client.
func (c *alertmanagerClient) RemoteAddress() string {
	return c.conn.Target()
}
