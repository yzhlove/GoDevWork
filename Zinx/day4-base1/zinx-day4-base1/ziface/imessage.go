package ziface

type MessageInterface interface {
	GetDataLen() uint32
	GetMessageID() uint32
	GetData() []byte

	SetDataLen(length uint32)
	SetMessageID(id uint32)
	SetData(data []byte)
}
