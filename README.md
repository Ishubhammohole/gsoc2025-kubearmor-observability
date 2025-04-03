# GSoC 2025: Improve KubeArmor Observability

This repo contains mock code, metric designs, and architecture related to my GSoC 2025 proposal with CNCF and KubeArmor.

## Summary
The project enhances observability in KubeArmor by exposing security policy metrics directly from the controller via Prometheus. It removes the need for the legacy exporter, simplifying monitoring and integration with tools like Grafana.

## Architecture

- gRPC APIs for policy metrics
- Metrics exposed directly from controller `/metrics`
- Integration with Prometheus scrape endpoint

## Project Structure

- `demo/`: Sample Go code for exposing new metrics
- `proto/`: Sample `.proto` changes for `WatchPolicies`
- `grafana/`: Example JSON panels for policy metrics
>>>>>>> 64cdb39 (Add working metrics server with Prometheus for GSoC demo)
