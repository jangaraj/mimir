FROM golang:1.20.1
ENV CGO_ENABLED=0
RUN go install github.com/go-delve/delve/cmd/dlv@v1.7.3

FROM alpine:3.17.2

RUN     mkdir /mimir
WORKDIR /mimir
COPY     ./mimir ./
COPY --from=0 /go/bin/dlv ./
