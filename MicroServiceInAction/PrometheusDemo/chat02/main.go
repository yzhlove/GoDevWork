package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

//Counter累加指标
//Gauge测量指标
//Summary概略图
//Histogram直方图

//cpu温度
var cpuTemp = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "cpu_temperature_celsius",
	Help: "Current temperature of the CPU",
})

//磁盘失败次数
var hdFailures = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "hd_errors_total",
	Help: "Number of hard-disk errors",
}, []string{"device"})

func init() {
	prometheus.MustRegister(cpuTemp, hdFailures)
}

func main() {
	cpuTemp.Set(65.3)
	hdFailures.With(prometheus.Labels{"device": "/dev/sda"}).Inc()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":1235", nil))

}
