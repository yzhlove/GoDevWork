package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"math/rand"
	"net/http"
	"time"
)

var httpReqCount = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_request_count",
		Help: "HTTP请求数",
	}, []string{"endpoint"})

var orderNum = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "order_total",
		Help: "订单数量",
	})

var httpReqTime = prometheus.NewSummaryVec(
	prometheus.SummaryOpts{
		Name: "http_request_time",
		Help: "HTTP请求用时",
	}, []string{"endpoint"})

func init() {
	prometheus.MustRegister(httpReqCount, orderNum, httpReqTime)
}

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/handle", handle)
	if err := http.ListenAndServe(":1234", nil); err != nil {
		panic(err)
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	httpReqCount.WithLabelValues(r.URL.Path).Inc()
	start := time.Now()
	if t := rand.Intn(10); t&t-1 == 0 {
		orderNum.Desc()
		time.Sleep(100 * time.Millisecond)
	} else {
		orderNum.Inc()
		time.Sleep(50 * time.Millisecond)
	}
	httpReqTime.WithLabelValues(r.URL.Path).Observe(float64(time.Since(start).Milliseconds() / 1e6))
	w.Write([]byte("ok"))
}
