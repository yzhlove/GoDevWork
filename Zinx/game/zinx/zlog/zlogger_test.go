package zlog

import (
	"bytes"
	"testing"
)

func Test_Stuff(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})
	stuff(buf, 6, 4)
	t.Log(buf.String())
}

func Test_Out(t *testing.T) {
	Debug("what are youd doing?")
	Info("what are youd doing?")
	Warn("what are youd doing?")
	Error("what are youd doing?")
	Panic("what are youd doing?")
	Fatal("what are youd doing?")
}


