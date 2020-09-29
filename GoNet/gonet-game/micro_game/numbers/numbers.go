package numbers

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/tealeg/xlsx"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"sync"
)

//numbers 读取的表结构
//（null） 	字段1	字段2	字段3
//	记录1	值		值		值
//	记录2	值		值		值
//	记录3	值		值		值

type Option interface {
	GetInt(tbname string, rowname interface{}, fieldname string) int32
	GetFloat(tbname string, rowname interface{}, fieldname string) float64
	GetString(tbname string, rowname interface{}, fieldname string) string
	GetKeys(tbname string) []string
	IsField(tbname string, rowname interface{}, fieldname string) bool
	IsRecord(tbname string, rowname interface{}) bool
	IsTable(tbname string) bool
}

var (
	_dataConfig configs
)

type record struct {
	fields map[string]string
}

type table struct {
	records map[string]*record
	keys    []string
}

type numbers struct {
	tables map[string]*table
	name   string
}

type configs struct {
	numbers map[string]*numbers
	sync.RWMutex
}

func Init(file string) {
	_dataConfig.init(file)
}

func (c *configs) init(file string) {
	if isDir(file) {
		err := filepath.Walk(file, func(p string, info os.FileInfo, err error) error {
			//如果info 为nil有可能是扫描路径存在问题，直接return nil 继续扫描
			if info == nil {
				log.Error("info is nil ,check path:", p)
				return nil
			}
			//跳过文件夹扫描
			if !info.IsDir() {
				if x, err := xlsx.OpenFile(p); err != nil {
					return err
				} else {
					number := &numbers{tables: make(map[string]*table), name: path.Base(p)}
					number.parse(path.Base(p), x.Sheets)
					SetNumbers(number)
				}
			}
			return nil
		})
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	log.Fatal("path must is dir:", file)
}

func isDir(file string) bool {
	if stat, err := os.Stat(file); err != nil {
		return false
	} else {
		return stat.IsDir()
	}
}

func Numbers(name string) Option {
	_dataConfig.RLock()
	defer _dataConfig.RUnlock()

	if n, ok := _dataConfig.numbers[name]; ok {
		return Option(n)
	}
	panic(fmt.Sprintf("numbers not exists %v", name))
}

func SetNumbers(n *numbers) {
	_dataConfig.Lock()
	defer _dataConfig.Unlock()
	if _dataConfig.numbers == nil {
		_dataConfig.numbers = make(map[string]*numbers)
	}
	_dataConfig.numbers[n.name] = n
}

func (n *numbers) parse(xmlname string, sheets []*xlsx.Sheet) {
	var name string
	defer func() {
		if x := recover(); x != nil {
			log.WithField("errMsg", fmt.Sprintf("xls %v sheet %v panic %v", xmlname, name, x)).
				WithField("err", x).Error()
		}
	}()

	for _, sheet := range sheets {
		log.Info("parse sheet ", sheet.Name)
		if len(sheet.Rows) > 0 {
			header := sheet.Rows[0]
			for i := 0; i < len(sheet.Rows); i++ {
				row := sheet.Rows[i]
				for j := 0; j < len(row.Cells); j++ {
					name := row.Cells[0].String()
					field := header.Cells[j].String()
					value := row.Cells[j].String()
					n.set(sheet.Name, name, field, value)
				}
			}
		}
		n.dump(sheet.Name)
	}

}

func (n *numbers) set(name, row, field, value string) {
	t, ok := n.tables[name]
	if !ok {
		t = new(table)
		t.records = make(map[string]*record)
		n.tables[name] = t
	}
	rec, ok := t.records[row]
	if !ok {
		rec = new(record)
		rec.fields = make(map[string]string)
		t.records[row] = rec
	}
	rec.fields[field] = value
}

func (n *numbers) dump(name string) {
	if t, ok := n.tables[name]; ok {
		for k := range t.records {
			t.keys = append(t.keys, k)
		}
	}
	panic("table not found by name:" + name)
}

func (n *numbers) pack(name, row, field string) string {
	if t, ok := n.tables[name]; ok {
		if r, ok := t.records[row]; ok {
			if v, ok := r.fields[field]; ok {
				return v
			}
		}
	}
	return ""
}

func (n *numbers) GetInt(name string, row interface{}, field string) int32 {
	var res string
	if res = n.pack(name, fmt.Sprint(row), field); len(res) == 0 {
		return 0
	}
	if i, err := strconv.ParseFloat(res, 64); err != nil {
		panic(fmt.Sprintf("integer parse err: %v %v %v %v \n", name, row, field, err))
	} else {
		return int32(i)
	}
}

func (n *numbers) GetFloat(name string, row interface{}, field string) float64 {
	var res string
	if res = n.pack(name, fmt.Sprint(row), field); len(res) == 0 {
		return 0
	}
	if f, err := strconv.ParseFloat(res, 64); err != nil {
		panic(fmt.Sprintf("float parse err:%v %v %v %v \n", name, row, field, err))
	} else {
		return f
	}
}

func (n *numbers) GetString(name string, row interface{}, field string) string {
	return n.pack(name, fmt.Sprint(row), field)
}

func (n *numbers) GetKeys(name string) []string {
	if t, ok := n.tables[name]; ok {
		return t.keys
	}
	return nil
}

func (n *numbers) Count(name string) int32 {
	if t, ok := n.tables[name]; ok {
		return int32(len(t.records))
	}
	return 0
}

func (n *numbers) IsTable(name string) (ok bool) {
	_, ok = n.tables[name]
	return
}

func (n *numbers) IsRecord(name string, row interface{}) (ok bool) {
	var t *table
	if t, ok = n.tables[name]; !ok {
		return
	}
	_, ok = t.records[fmt.Sprint(row)]
	return
}

func (n *numbers) IsField(name string, row interface{}, field string) (ok bool) {
	var t *table
	if t, ok = n.tables[name]; !ok {
		return
	}
	var r *record
	if r, ok = t.records[fmt.Sprint(row)]; !ok {
		return
	}
	_, ok = r.fields[field]
	return
}
