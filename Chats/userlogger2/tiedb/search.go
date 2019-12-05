package tiedb

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/HouzuoGuo/tiedot/db"
	"time"
	"yo-star.com/nekopara/manager/logger/userlogger/base"
)

type LogQuery interface {
	Query(tie *TieReader, cond base.LogCondMessage) (interface{}, error)
	GetQueryName() string
}

//List 查询接口列表
var List []LogQuery

func init() {
	List = append(List, new(BaseQuery), new(TsQuery), new(OperatorEventTsQuery))
}

//查询所有，仅用于测试
type BaseQuery struct{}

func (BaseQuery) GetQueryName() string {
	return "all"
}

func (b BaseQuery) Query(tie *TieReader, cond base.LogCondMessage) (interface{}, error) {
	docs := make([]interface{}, 0, tie.Col.ApproxDocCount()>>1)
	tie.Col.ForEachDoc(func(_ int, doc []byte) bool {
		var data interface{}
		if err := json.Unmarshal(doc, &data); err != nil {
			fmt.Println(b.GetQueryName(), " Query Err:", err)
			return false
		}
		docs = append(docs, data)
		return true
	})
	return docs, nil
}

//按时间查询 (用户时间条件的测试)
type TsQuery struct{}

func (TsQuery) GetQueryName() string {
	return "time"
}

func (t TsQuery) Query(tie *TieReader, cond base.LogCondMessage) (interface{}, error) {

	type C struct {
		Mints int64 `json:"start"`
		Maxts int64 `json:"end"`
		Limit int   `json:"limit"`
	}

	var conds struct {
		C `json:"cond"`
	}

	if err := json.Unmarshal([]byte(cond), &conds); err != nil {
		return nil, err
	}

	if conds.Mints == 0 {
		return nil, errors.New("start time is not null")
	}

	if conds.Maxts == 0 {
		conds.Maxts = time.Now().Unix()
	}

	if conds.Limit == 0 {
		conds.Limit = 100
	}

	min := time.Unix(conds.Mints, 0)
	max := time.Unix(conds.Maxts, 0)

	//不是同一年或者不是同一个月
	if min.Year() != max.Year() || min.Month() != max.Month() {
		return nil, errors.New("must be the same year and month")
	}

	var cs []interface{}

	cs = append(cs, map[string]interface{}{
		"int-from": min.YearDay(),
		"int-to":   max.YearDay(),
		"in":       []interface{}{"day"},
	})

	cs = append(cs, map[string]interface{}{
		"eq": min.Year(),
		"in": []interface{}{"year"},
	})

	search := map[string]interface{}{
		"n": cs,
	}

	ids := make(map[int]struct{}, conds.Limit)
	if err := db.EvalQuery(search, tie.Col, &ids); err != nil {
		fmt.Println(t.GetQueryName(), " query err:", err)
		return nil, err
	}

	docs := make([]interface{}, 0, len(ids))
	for id := range ids {
		if doc, err := tie.Col.Read(id); err != nil {
			return nil, err
		} else {
			if _, ok := doc["ts"]; ok {
				if t, ok := doc["ts"].(float64); ok {
					fmt.Printf("s => %v n => %v \n", t, int64(t))
				}
			}
			docs = append(docs, doc)
		}
	}
	return docs, nil
}

type OperatorEventTsQuery struct{}

func (OperatorEventTsQuery) GetQueryName() string {
	return "complex"
}

func (t OperatorEventTsQuery) Query(tie *TieReader, cond base.LogCondMessage) (interface{}, error) {

	type C struct {
		Operator string `json:"operator"` //操作人员
		Mints    int64  `json:"start"`    //开始时间
		Maxts    int64  `json:"end"`      //结束时间
		Event    string `json:"event"`    //事件
		Limit    int    `json:"limit"`    //查询条数
	}

	var cds struct {
		C `json:"cond"`
	}

	if err := json.Unmarshal([]byte(cond), &cds); err != nil {
		return nil, err
	}

	if cds.Operator == "" && cds.Mints == 0 && cds.Event == "" {
		return nil, errors.New(t.GetQueryName() + " must satisfy one (operator,ts,event)")
	}
	if cds.Maxts == 0 {
		cds.Maxts = time.Now().Unix()
	}
	if cds.Limit == 0 {
		cds.Limit = 1000
	}

	var search []interface{}

	if cds.Operator != "" {
		search = append(search, map[string]interface{}{
			"eq":    cds.Operator,
			"in":    []interface{}{"operator"},
			"limit": cds.Limit,
		})
	}

	if cds.Mints != 0 {

		start := time.Unix(cds.Mints, 0)
		end := time.Unix(cds.Maxts, 0)

		if start.Year() != end.Year() || start.Month() != end.Month() {
			return nil, errors.New("must be the same year or month")
		}

		search = append(search, map[string]interface{}{
			"int-from": start.YearDay(),
			"int-to":   end.YearDay(),
			"in":       []interface{}{"day"},
			"limit":    cds.Limit,
		})
	}

	if cds.Event != "" {
		search = append(search, map[string]interface{}{
			"eq":    cds.Event,
			"in":    []interface{}{"event"},
			"limit": cds.Limit,
		})
	}

	ids := make(map[int]struct{}, cds.Limit)
	if err := db.EvalQuery(map[string]interface{}{"n": search}, tie.Col, &ids); err != nil {
		fmt.Println(t.GetQueryName(), " query err:", err)
		return nil, err
	}
	docs := make([]interface{}, 0, len(ids))
	for id := range ids {
		if doc, err := tie.Col.Read(id); err != nil {
			return nil, err
		} else {
			if cds.Mints != 0 && cds.Maxts != 0 {
				if _, ok := doc["ts"]; ok {
					if t, ok := doc["ts"].(float64); ok {
						if int64(t) >= cds.Mints && int64(t) <= cds.Maxts {
							docs = append(docs, doc)
						}
					}
				}
			} else {
				docs = append(docs, doc)
			}
		}
	}
	return docs, nil
}
