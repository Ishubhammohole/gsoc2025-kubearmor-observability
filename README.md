# GSoC 2025: Improve KubeArmor Observability

This repo contains mock code, metric designs, and architecture related to my GSoC 2025 proposal with CNCF and KubeArmor.

## ✨ Summary
This project enhances observability in KubeArmor by exposing security policy metrics directly from the controller via Prometheus. It removes the need for the legacy exporter, simplifying monitoring and integration with tools like Grafana.

## 🌎 Architecture
- gRPC APIs for policy metrics
- Metrics exposed directly from controller `/metrics`
- Integration with Prometheus scrape endpoint

![Architecture](./demo/architecture.png)

## 📇 Project Structure

## 🔥 Running the Metrics Demo

Start the server:
```bash
go run demo/main.go

# KubeArmor Observability Mock

This project demonstrates a simulated observability stack for KubeArmor.

## Components

- ✅ Go-based Prometheus mock server
- ✅ Prometheus instance scraping custom metrics
- ✅ Grafana dashboard visualizing metrics

## Setup

1. Run metrics server
   ```bash
   cd demo
   go run main.go
