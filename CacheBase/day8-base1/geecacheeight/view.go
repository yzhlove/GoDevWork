package geecacheeight

type ByteView struct {
	buf []byte
}

func (v ByteView) Len() int {
	return len(v.buf)
}

func (v ByteView) Bytes() []byte {
	return _copy(v.buf)
}

func (v ByteView) String() string {
	return string(v.buf)
}

func _copy(data []byte) []byte {
	if l := len(data); l > 0 {
		buf := make([]byte, l)
		copy(buf, data)
		return buf
	}
	return []byte{}
}
