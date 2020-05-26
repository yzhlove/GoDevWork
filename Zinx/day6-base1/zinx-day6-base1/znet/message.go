package znet

type Msg struct {
	ID   uint32
	Size uint32
	Data []byte
}

func NewMsgPackage(id uint32, data []byte) *Msg {
	return &Msg{
		ID:   id,
		Size: uint32(len(data)),
		Data: data,
	}
}

func (m *Msg) GetDataSize() uint32 {
	return m.Size
}

func (m *Msg) GetMsgID() uint32 {
	return m.ID
}

func (m *Msg) GetData() []byte {
	return m.Data
}

func (m *Msg) SetDataSize(size uint32) {
	m.Size = size
}

func (m *Msg) SetMsgID(id uint32) {
	m.ID = id
}

func (m *Msg) SetData(data []byte) {
	m.Data = data
}
