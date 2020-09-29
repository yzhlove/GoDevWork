module gonet_micro_game

go 1.14

require (
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/google/uuid v1.1.2 // indirect
	github.com/sirupsen/logrus v1.6.0
	go.uber.org/zap v1.16.0 // indirect
	google.golang.org/grpc v1.32.0
	micro_game v0.0.0
)

replace micro_game => ./micro_game
