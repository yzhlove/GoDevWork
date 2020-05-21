package znet

type Message struct {
	ID      uint32
	DataLen uint32
	Data    []byte
}

func (m *Message) GetDataLen() uint32 {
	return m.DataLen
}

func (m *Message) GetMsgID() uint32 {
	return m.ID
}

func (m *Message) GetData() []byte {
	return m.Data
}

func (m *Message) SetMsgID(u uint32) {
	m.ID = u
}

func (m *Message) SetData(bytes []byte) {
	m.Data = bytes
}

func (m *Message) SetDataLen(u uint32) {
	m.DataLen = u
}

func NewMessagePackage(id uint32, data []byte) *Message {
	return &Message{
		ID:      id,
		DataLen: uint32(len(data)),
		Data:    data,
	}
}
