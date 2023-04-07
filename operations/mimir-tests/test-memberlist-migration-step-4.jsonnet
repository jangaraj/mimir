local mimir = import 'mimir/mimir.libsonnet';

mimir {
  _config+:: {
    namespace: 'default',
    external_url: 'http://test',

    storage_backend: 'gcs',
    blocks_storage_bucket_name: 'blocks-bucket',

    ruler_enabled: true,
    ruler_storage_bucket_name: 'rules-bucket',

    alertmanager_enabled: true,
    alertmanager_storage_bucket_name: 'alerts-bucket',

    // Step 4: disable mirroring from primary (now memberlist) to secondary (now Consul) KV.
    multikv_migration_enabled: true,
    multikv_mirror_enabled: false,
    multikv_switch_primary_secondary: true,
  },
}
