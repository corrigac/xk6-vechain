# xk6-vechain

This repository contains the xk6 extension for interacting with the VeChain blockchain.

## Usage

To build the executable:

```bash
XK6_RACE_DETECTOR=1 xk6 build --with github.com/darrenvechain/xk6-vechain=. --with github.com/grafana/xk6-dashboard@latest --with github.com/grafana/xk6-output-influxdb@latest
```

Start thor solo:

```
docker run -p 8669:8669 ghcr.io/vechain/thor:release-galactica-latest solo --api-addr 0.0.0.0:8669
```

```
docker run -p 8669:8669 ghcr.io/vechain/thor:release-galactica-latest solo --api-addr 0.0.0.0:8669 --api-priority-fees-percentage 2
```

**Optional**: Start Grafana + InfluxDB:

```bash
docker compose up -d --wait
```

To run the tests:

```bash
K6_INFLUXDB_ORGANIZATION=vechain \
K6_INFLUXDB_BUCKET=vechain \
K6_INFLUXDB_TOKEN=admin-token \
./k6 run -o xk6-influxdb=http://localhost:8086 --quiet test.js
```
