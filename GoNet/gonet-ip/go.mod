module gonet_micro_geoip

go 1.14

require (
	github.com/sirupsen/logrus v1.7.0
	golang.org/x/net v0.0.0-20190311183353-d8887717615a
	google.golang.org/grpc v1.32.0
	micro_geoip v0.0.0
)

replace micro_geoip => ./micro_geoip
