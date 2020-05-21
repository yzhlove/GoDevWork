package ziface

type DataPackInterface interface {
	GetHeadLen() uint32
	Pack(msg MessageInterface) ([]byte, error)
	Unpack(data []byte) (MessageInterface, error)
}
