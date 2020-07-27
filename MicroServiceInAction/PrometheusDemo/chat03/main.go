package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"math/rand"
	"net/http"
)

type ClusterManager struct {
	Zone         string
	OOMCountDesc *prometheus.Desc
	RAMUsageDesc *prometheus.Desc
}

func (c *ClusterManager) Really() (oomHost map[string]int, ramHost map[string]float64) {
	oomHost = map[string]int{
		"foo.example.org": int(rand.Int31n(1000)),
		"bar.example.org": int(rand.Int31n(1000)),
	}
	ramHost = map[string]float64{
		"foo.example.org": rand.Float64() * 100,
		"bar.example.org": rand.Float64() * 100,
	}
	return
}

//传递指标
func (c *ClusterManager) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.OOMCountDesc
	ch <- c.RAMUsageDesc
}

//绑定指标
func (c *ClusterManager) Collect(ch chan<- prometheus.Metric) {
	oom, ram := c.Really()
	for host, count := range oom {
		ch <- prometheus.MustNewConstMetric(
			c.OOMCountDesc, prometheus.CounterValue, float64(count), host)
	}
	for host, usage := range ram {
		ch <- prometheus.MustNewConstMetric(
			c.RAMUsageDesc, prometheus.GaugeValue, usage, host)
	}
}

func NewClusterManager(zone string) *ClusterManager {
	return &ClusterManager{
		OOMCountDesc: prometheus.NewDesc(
			"clustermanager_oom_crashes_total",
			"错误计数",
			[]string{"host"},
			prometheus.Labels{"zone": zone}),
		RAMUsageDesc: prometheus.NewDesc(
			"clustermanager_ram_usage_bytes",
			"内存使用信息",
			[]string{"host"},
			prometheus.Labels{"zone": zone}),
	}
}

func main() {

	workerDB := NewClusterManager("db")
	workerCA := NewClusterManager("ca")

	reg := prometheus.NewPedanticRegistry()
	reg.MustRegister(workerDB)
	reg.MustRegister(workerCA)

	gatherers := prometheus.Gatherers{
		prometheus.DefaultGatherer,
		reg,
	}
	h := promhttp.HandlerFor(gatherers, promhttp.HandlerOpts{
		ErrorLog:      log.NewErrorLogger(),
		ErrorHandling: promhttp.ContinueOnError,
	})
	http.HandleFunc("/metrics", func(writer http.ResponseWriter, request *http.Request) {
		h.ServeHTTP(writer, request)
	})
	if err := http.ListenAndServe(":1234", nil); err != nil {
		panic(err)
	}
}
