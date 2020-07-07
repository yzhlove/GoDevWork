package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type RawProjectCfg struct {
	Name string       `json:"name"`
	Data BTProjectCfg `json:"data"`
	Path string       `json:"path"`
}

func LoadRawProjectCfg(path string) (*RawProjectCfg, bool) {

	var project RawProjectCfg
	f, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("LoadRawProjectCfg Fail:", err)
		return nil, false
	}
	if err := json.Unmarshal(f, &project); err != nil {
		fmt.Println("LoadRawProjectCfg Unmarshal Fail:", err)
		return nil, false
	}
	return &project, true
}
