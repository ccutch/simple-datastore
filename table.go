package simpledb

import "reflect"

type Table struct {
	store  *Datastore
	name   string
	fields Schema
	data   []Node
}

type Node map[string]interface{}

func getName(f reflect.StructField) string {
	name := f.Tag.Get("simpledb")
	if name == "" {
		name = f.Name
	}
	return name
}

func getRootName(t reflect.Type, key string) string {
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		name := getName(f)
		if name == key {
			return f.Name
		}
	}
	return ""
}

func Nodify(s interface{}) Node {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	l := t.NumField()
	n := Node{}
	for i := 0; i < l; i++ {
		f := t.Field(i)
		name := getName(f)
		n[name] = v.FieldByName(f.Name).Interface()
	}
	return n
}

func (n Node) Parse(out interface{}) {
	v := reflect.ValueOf(out).Elem()
	t := v.Type()
	for key, value := range n {
		name := getRootName(t, key)
		if name != "" {
			v.FieldByName(name).Set(reflect.ValueOf(value))
		}
	}
}

func (t *Table) Insert(n Node) error {
	t.data = append(t.data, n)
	return nil
}
