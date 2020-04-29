package geecachefive

type ByteView struct {
	buffer []byte
}

func (v ByteView) Len() int {
	return len(v.buffer)
}

func (v ByteView) GetBytes() []byte {
	return _copy(v.buffer)
}

func (v ByteView) String() string {
	return string(v.buffer)
}

func _copy(data []byte) []byte {
	if length := len(data); length > 0 {
		buffer := make([]byte, length)
		copy(buffer, data)
		return buffer
	}
	return []byte{}
}
