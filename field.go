package simpledb

import "reflect"

type Field string
type Schema map[string]Field

func (f Field) Parse(v interface{}) interface{} {
	return ""
}

func Schemafy(s interface{}) Schema {
	sc := Schema{}
	t := reflect.TypeOf(s)
	l := t.NumField()
	for i := 0; i < l; i++ {
		f := t.Field(i)
		name := getName(f)
		tp := f.Tag.Get("relatesTo")
		if tp == "" {
			tp = f.Type.String()
		} else {
			tp = "_" + tp
		}
		sc[name] = Field(tp)
	}
	return sc
}
