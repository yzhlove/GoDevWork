package znet

type Message struct {
	ID     uint32
	Length uint32
	Data   []byte
}

func NewMessagePackage(id uint32, data []byte) *Message {
	return &Message{
		ID:     id,
		Length: uint32(len(data)),
		Data:   data,
	}
}

func (m Message) GetDataLen() uint32 {
	return m.Length
}

func (m Message) GetMessageID() uint32 {
	return m.ID
}

func (m Message) GetData() []byte {
	return m.Data
}

func (m *Message) SetDataLen(length uint32) {
	m.Length = length
}

func (m *Message) SetMessageID(id uint32) {
	m.ID = id
}

func (m *Message) SetData(data []byte) {
	m.Data = data
}
