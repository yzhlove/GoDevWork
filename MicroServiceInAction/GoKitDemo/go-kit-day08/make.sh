# 启动etcd
ETCD_API=3 ./etcd

# 启动prometheus
docker run -d -p 9090:9090 -v /Users/yostar/promtheus/prometheus.yml:/etc/prometheus/prometheus.yml prom/prometheus

# 启动Jaeger
docker run -d -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 -p5775:5775/udp -p6831:6831/udp -p6832:6832/udp -p5778:5778 -p16686:16686 -p14268:14268 -p9411:9411 jaegertracing/all-in-one:latest

# 启动Hystrix UI
docker run --name hystrix-dashboard -d -p 8081:9002 mlabouardy/hystrix-dashboard:latest

