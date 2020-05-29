package znet

import "zinx/ziface"

type Req struct {
	conn ziface.ConnImp
	msg  ziface.MsgImp
}

func (r *Req) GetConn() ziface.ConnImp {
	return r.conn
}

func (r *Req) GetMsgData() []byte {
	return r.msg.GetData()
}

func (r *Req) GetMsgID() uint32 {
	return r.msg.GetID()
}
