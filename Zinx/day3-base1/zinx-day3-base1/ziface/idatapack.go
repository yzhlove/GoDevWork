package ziface

type DataPackInterface interface {
	GetHeadLen() uint32
	Pack(MessageInterface) ([]byte, error)
	Unpack([]byte) (MessageInterface, error)
}
