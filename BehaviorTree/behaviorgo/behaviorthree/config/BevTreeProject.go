package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type BTProjectCfg struct {
	ID     string      `json:"id"`
	Select string      `json:"selectedTree"`
	Scope  string      `json:"scope"`
	Trees  []BTTreeCfg `json:"trees"`
}

func LoadProjectCfg(path string) (*BTProjectCfg, bool) {
	var project BTProjectCfg
	f, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("LoadProjectCfg Fail: ", err)
		return nil, false
	}
	if err := json.Unmarshal(f, &project); err != nil {
		fmt.Println("LoadProjectCfg Unmarshal Fail:", err)
		return nil, false
	}
	return &project, true
}
