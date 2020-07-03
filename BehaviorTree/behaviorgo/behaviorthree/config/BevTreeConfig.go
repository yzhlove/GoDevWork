package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type BTNodeCfg struct {
	Id          string                 `json:"id"`
	Name        string                 `json:"name"`
	Category    string                 `json:"category"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Children    []string               `json:"children"`
	Child       string                 `json:"child"`
	Parameters  map[string]interface{} `json:"parameters"`
	Properties  map[string]interface{} `json:"properties"`
}

func (this *BTNodeCfg) GetProperty(name string) float64 {
	if v, ok := this.Properties[name]; ok {
		if f64, ok := v.(float64); ok {
			return f64
		}
	}
	panic("GetProperty err,no value:" + name)
}

func (this *BTNodeCfg) GetPropertyAsInt(name string) int {
	return int(this.GetProperty(name))
}

func (this *BTNodeCfg) GetPropertyAsInt64(name string) int64 {
	return int64(this.GetProperty(name))
}

func (this *BTNodeCfg) GetPropertyAsBool(name string) bool {
	if v, ok := this.Properties[name]; ok {
		if b, ok := v.(bool); !ok {
			if str, ok := v.(string); !ok {
				fmt.Println("GetProperty err, format not bool:", name, v)
				panic("GetProperty err, format bot bool " + name)
			} else {
				return strings.ToLower(strings.TrimSpace(str)) == "true"
			}
		} else {
			return b
		}
	}
	return false
}

func (this *BTNodeCfg) GetPropertyAsString(name string) string {
	if v, ok := this.Properties[name]; ok {
		if str, ok := v.(string); ok {
			return str
		}
	}
	panic("GetProperty err,no value:" + name)
}

type BTTreeCfg struct {
	ID          string                 `json:"id"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Root        string                 `json:"root"`
	Properties  map[string]interface{} `json:"properties"`
	Nodes       map[string]BTNodeCfg   `json:"nodes"`
}

func LoadTreeCfg(path string) (*BTTreeCfg, bool) {
	var tree BTTreeCfg
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("fail:", err)
		return nil, false
	}
	if err := json.Unmarshal(file, &tree); err != nil {
		fmt.Println("fail, unmarshal:", err, len(file))
		return nil, false
	}
	return &tree, true
}
