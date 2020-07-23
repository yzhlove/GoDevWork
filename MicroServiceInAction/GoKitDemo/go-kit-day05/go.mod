module go-kit-five-example

go 1.14

require (
	github.com/go-kit/kit v0.10.0
	go-kit-five v0.0.0
	go.uber.org/zap v1.15.0
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0
	google.golang.org/grpc v1.30.0
)

replace go-kit-five => ./go-kit-five
