## duckdns-updater

![](https://github.com/vladimir-kotikov/duckdns-updater/workflows/Build%20image/badge.svg)


A simple Go program wrapped in Docker container that updates DuckDns domains
provided via REST API.

### Configuration

The updater is configured using 2 environment variables:

  - `DUCKDNS_DOMAINS` - comma separated list of domains to update A records for
  - `DUCKDNS_TOKEN` - API token, can be obtained on https://www.duckdns.org/

Update repeats every 5 minutes (not configurable, sorry)
