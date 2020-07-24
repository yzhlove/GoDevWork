module go-kit-six

go 1.14

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-kit/kit v0.10.0
	github.com/golang/protobuf v1.4.2
	github.com/natefinch/lumberjack v2.0.0+incompatible
	go.uber.org/zap v1.15.0
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e
	google.golang.org/grpc v1.30.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
)


replace (
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)