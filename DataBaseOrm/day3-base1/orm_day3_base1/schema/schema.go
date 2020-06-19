package schema

import (
	"go/ast"
	"orm_day3_base1/dialect"
	"reflect"
)

type Field struct {
	Name string //字段名
	Type string //字段类型
	Tag  string //字段标签
}

type Schema struct {
	Model      interface{} //source object 对应database的表
	Name       string      //object name 对应database table name
	Fields     []*Field    //对应database 字段具体
	FieldNames []string    //对应数据库的字段
	fieldMap   map[string]*Field
}

func (schema *Schema) GetField(name string) *Field {
	return schema.fieldMap[name]
}

func (schema *Schema) RecordValues(dest interface{}) []interface{} {
	indirect := reflect.Indirect(reflect.ValueOf(dest))
	var source []interface{}
	for _, field := range schema.Fields {
		source = append(source, indirect.FieldByName(field.Name).Interface())
	}
	return source
}

func Parse(dest interface{}, d dialect.Dialect) *Schema {
	indirect := reflect.Indirect(reflect.ValueOf(dest)).Type()
	schema := &Schema{
		Model:    dest,
		Name:     indirect.Name(),
		fieldMap: make(map[string]*Field, indirect.NumField()),
	}
	for i := 0; i < indirect.NumField(); i++ {
		structField := indirect.Field(i)
		if !structField.Anonymous && ast.IsExported(structField.Name) {
			field := &Field{
				Name: structField.Name,
				Type: d.GetDataBaseType(reflect.Indirect(reflect.New(structField.Type))),
			}
			if value, ok := structField.Tag.Lookup("geeorm"); ok {
				field.Tag = value
			}
			schema.Fields = append(schema.Fields, field)
			schema.FieldNames = append(schema.FieldNames, structField.Name)
			schema.fieldMap[structField.Name] = field
		}
	}
	return schema
}
