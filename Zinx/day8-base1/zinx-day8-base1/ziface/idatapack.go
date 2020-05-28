package ziface

type DataPackInterface interface {
	MaxHead() uint32
	Pack(msg MsgInterface) ([]byte, error)
	Unpack(data []byte) (MsgInterface, error)
}
