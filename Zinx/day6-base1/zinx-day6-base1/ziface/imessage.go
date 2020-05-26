package ziface

type MsgInterface interface {
	GetDataSize() uint32
	GetMsgID() uint32
	GetData() []byte
	SetDataSize(size uint32)
	SetMsgID(id uint32)
	SetData(data []byte)
}
