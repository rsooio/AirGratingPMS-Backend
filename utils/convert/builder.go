package convert

import (
	"fmt"
	"reflect"
)

func rawFieldNames(in interface{}, postgresSql ...bool) map[string]string {
	out := map[string]string{}
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// we only accept structs
	if v.Kind() != reflect.Struct {
		panic(fmt.Errorf("ToMap only accepts structs; got %T", v))
	}

	typ := v.Type()

	for i := 0; i < v.NumField(); i++ {
		// gets us a StructField
		fi := typ.Field(i)
		out[fi.Name] = fi.Type.Name()
	}

	return out
}
