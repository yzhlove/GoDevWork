package znet

type Msg struct {
	ID   uint32
	Size uint32
	Data []byte
}

func NewMsg(id uint32, data []byte) *Msg {
	return &Msg{
		ID:   id,
		Size: uint32(len(data)),
		Data: data,
	}
}

func (m *Msg) GetID() uint32 {
	return m.ID
}

func (m *Msg) GetSize() uint32 {
	return m.Size
}

func (m *Msg) GetData() []byte {
	return m.Data
}

func (m *Msg) SetID(id uint32) {
	m.ID = id
}

func (m *Msg) SetSize(size uint32) {
	m.Size = size
}

func (m *Msg) SetData(data []byte) {
	m.Data = data
}
