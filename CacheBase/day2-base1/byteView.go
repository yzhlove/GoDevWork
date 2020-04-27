package geecachetwo

type ByteView struct {
	buf []byte
}

func (view ByteView) Len() int {
	return len(view.buf)
}

func (view ByteView) ByteSlice() []byte {
	return cloneBytes(view.buf)
}

func cloneBytes(buf []byte) []byte {
	newBuffer := make([]byte, len(buf))
	copy(newBuffer, buf)
	return newBuffer
}

func (view ByteView) String() string {
	return string(view.buf)
}
