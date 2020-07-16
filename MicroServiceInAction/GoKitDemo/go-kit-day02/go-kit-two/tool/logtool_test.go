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
