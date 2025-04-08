package main

import (
    "net/http"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
    policiesApplied = prometheus.NewGauge(
        prometheus.GaugeOpts{
            Name: "kubearmor_policies_applied_total",
            Help: "Total number of KubeArmor policies applied",
        },
    )

    alertsTriggered = prometheus.NewCounter(
        prometheus.CounterOpts{
            Name: "kubearmor_alerts_triggered_total",
            Help: "Total number of KubeArmor security alerts triggered",
        },
    )
)

func init() {
    prometheus.MustRegister(policiesApplied)
    prometheus.MustRegister(alertsTriggered)
}

func main() {
    policiesApplied.Set(5)
    alertsTriggered.Add(12)

    http.Handle("/metrics", promhttp.Handler())
    http.ListenAndServe(":2112", nil)
}
