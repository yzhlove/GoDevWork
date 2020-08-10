module go-kit-nine-example

go 1.14

require (
	github.com/go-kit/kit v0.10.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.1-0.20190118093823-f849b5445de4
	github.com/prometheus/client_golang v1.3.0
	go-kit-nine v0.0.0
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0
	google.golang.org/grpc v1.31.0
)

replace (
	go-kit-nine => ./go-kit-nine
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)
