package config

var (
	Listen               = ":53000"
	MaxConcurrentStreams = uint32(1024)
	RedisHost            = "127.0.0.1:6379"
	RedisDb              = 0
)
