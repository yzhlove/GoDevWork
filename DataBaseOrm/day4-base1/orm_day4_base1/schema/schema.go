package schema

import (
	"go/ast"
	"orm_day4_base1/dialect"
	"reflect"
)

//object => database

type Field struct {
	Name string
	Type string
	Tag  string
}

type Schema struct {
	Model      interface{}
	Name       string
	Fields     []*Field
	FieldNames []string
	fieldMap   map[string]*Field
}

//去重
func (schema *Schema) GetField(name string) *Field {
	return schema.fieldMap[name]
}

func (schema *Schema) Build(dest interface{}) []interface{} {
	indirect := reflect.Indirect(reflect.ValueOf(dest))
	source := make([]interface{}, 0, len(schema.Fields))
	for _, field := range schema.Fields {
		source = append(source, indirect.FieldByName(field.Name).Interface())
	}
	return source
}

//object => database
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
				Type: d.GetDBType(reflect.Indirect(reflect.New(structField.Type))),
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
