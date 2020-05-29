package ziface

type MsgImp interface {
	GetSize() uint32
	GetData() []byte
	GetID() uint32
	SetSize(size uint32)
	SetID(id uint32)
	SetData(data []byte)
}
