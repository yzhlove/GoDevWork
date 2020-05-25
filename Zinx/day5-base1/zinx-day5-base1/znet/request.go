package znet

import "zinx-day5-base1/ziface"

type Request struct {
	conn ziface.ConnectionInterface
	msg  ziface.MessageInterface
}

func (r *Request) GetConn() ziface.ConnectionInterface {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

func (r *Request) GetMessageID() uint32 {
	return r.msg.GetMessageID()
}
