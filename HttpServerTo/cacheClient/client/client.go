package client

const (
	RedisClient = iota
	HttpClient
	TcpClient
)

type Message struct {
	Name  string
	Key   string
	Value string
	Error error
}

type Client interface {
	Run(*Message)
	PipeLineRun([] *Message)
}

func New(c int, server string) Client {

	switch c {
	case RedisClient:

	case HttpClient:

	case TcpClient:
		return newTcpClient(server)
	}
	panic("unknown client type")
}
