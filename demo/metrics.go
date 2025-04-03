package main

import (
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "log"
    "net/http"
)

var (
    // Total policies applied per namespace
    policyTotal = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "kubearmor_policies_total",
            Help: "Total number of policies applied per namespace",
        },
        []string{"namespace"},
    )

    // Policy status gauge: 1 = active, 0 = inactive
    policyStatus = prometheus.NewGaugeVec(
        prometheus.GaugeOpts{
            Name: "kubearmor_policy_status",
            Help: "Status of policy (1=active, 0=inactive)",
        },
        []string{"policy_name", "namespace"},
    )
)

func init() {
    prometheus.MustRegister(policyTotal)
    prometheus.MustRegister(policyStatus)
}

func main() {
    // Simulate some test metrics
    policyTotal.WithLabelValues("default").Add(3)
    policyStatus.WithLabelValues("block-nginx", "default").Set(1)

    http.Handle("/metrics", promhttp.Handler())
    log.Println("Serving metrics at :9100/metrics")
    log.Fatal(http.ListenAndServe(":9100", nil))
}
