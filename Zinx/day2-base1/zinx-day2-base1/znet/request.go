package znet

import "zinx-day2-base1/ziface"

type Request struct {
	conn ziface.ConnectionInterface
	data []byte
}

func (r *Request) GetConn() ziface.ConnectionInterface {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.data
}

