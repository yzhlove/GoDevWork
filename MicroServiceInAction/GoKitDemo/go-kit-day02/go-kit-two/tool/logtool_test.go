package tool

import (
	"path/filepath"
	"testing"
)

func TestPath(t *testing.T) {
	pth, err := filepath.Abs(filepath.Dir(filepath.Join(".")))
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(pth)
}

func TestPathFile(t *testing.T) {
	t.Log(filepath.Join("."))
	t.Log(filepath.Dir(filepath.Join(".")))
	t.Log(filepath.Abs(filepath.Dir(filepath.Join("."))))
}
