package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	str := "/Users/yostar/Perforce/NovaMacminiMasterPoj/GameDataTables/tables/AVG/a.xlsx"
	pathRef(str)
}

func pathRef(path string) string {
	if index := strings.Index(path, "tables"); index != -1 {
		a := path[:index]
		fmt.Println(a)
		fmt.Println(filepath.Dir(path))
		fmt.Println(filepath.Base(path))
		fmt.Println(filepath.Ext(path))
	}
	return ""
}
