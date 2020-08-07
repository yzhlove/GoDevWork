module go-kit-eight

go 1.14

require (
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-kit/kit v0.10.0
	github.com/golang/protobuf v1.4.2
	github.com/natefinch/lumberjack v2.0.0+incompatible
	go.uber.org/zap v1.15.0
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0
	google.golang.org/grpc v1.31.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
)

replace (
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)