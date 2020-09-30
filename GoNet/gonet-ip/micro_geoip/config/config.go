package config

const (
	LANG    = "en"
	AREA    = "zh-CN"
	SERVICE = "[GEOIP]"
)

type Config struct {
	Path    string
	Host    string
	Streams uint32
}
