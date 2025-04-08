# 🛡️ GSoC 2025: Improve KubeArmor Observability

This repository contains mock code, metric designs, and architecture for my GSoC 2025 proposal with CNCF and KubeArmor.

## ✨ Summary

This project enhances observability in KubeArmor by exposing security policy metrics directly from the controller via Prometheus.  
It removes the need for the legacy exporter, simplifying monitoring and integration with tools like Grafana.

## 🌎 Architecture

- 📡 gRPC APIs for policy metrics
- 📊 Metrics exposed directly from controller `/metrics` endpoint
- 🔗 Integration with Prometheus scrape endpoint
- 📈 Visualization using Grafana dashboard

![Architecture Diagram](path/to/architecture-diagram.png)

## 📇 Project Structure
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
