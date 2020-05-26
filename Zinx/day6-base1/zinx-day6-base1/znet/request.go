package znet

import "zinx-day6-base1/ziface"

type Req struct {
	conn ziface.ConnectionInterface
	msg  ziface.MsgInterface
}

func (r *Req) GetConn() ziface.ConnectionInterface {
	return r.conn
}

func (r *Req) GetData() []byte {
	return r.msg.GetData()
}

func (r *Req) GetMsgID() uint32 {
	return r.msg.GetMsgID()
}
