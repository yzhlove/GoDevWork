module zinx-game-example

go 1.14

require (
	github.com/golang/protobuf v1.3.4
	snowflake v0.0.0
	zinx v0.0.0
)

replace (
	github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0
	snowflake => ./snowflake
	zinx => ./zinx
)
