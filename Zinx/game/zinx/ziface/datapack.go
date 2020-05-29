package ziface

type PackImp interface {
	HeadSize() uint32
	Pack(msg MsgImp) ([]byte, error)
	Unpack(data []byte) (MsgImp, error)
}
