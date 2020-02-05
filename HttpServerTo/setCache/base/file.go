package base

import "os"

type FileWriteInterface interface {
	Writer(string)
	Stop()
}

type FileWrite struct {
	f *os.File
}

func NewFileWrite() *FileWrite {
	path := "test_file.data"
	t, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		panic("create file writer err: " + err.Error())
	}
	return &FileWrite{f: t}
}

func (fw *FileWrite) Writer(str string) {
	if fw.f != nil {
		_, _ = fw.f.Write([]byte(str))
	}
}

func (fw *FileWrite) Stop() {
	if fw.f != nil {
		_ = fw.f.Close()
	}
}
