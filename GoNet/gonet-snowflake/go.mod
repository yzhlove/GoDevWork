module gonet_micro_snowflake

go 1.14

require (
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/google/uuid v1.1.2 // indirect
	go.uber.org/zap v1.16.0 // indirect
	google.golang.org/grpc v1.31.1
	micro_snowflake v0.0.0
)

replace (
	micro_snowflake => ./micro_snowflake
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
	github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0
)