package partial

import (
	"fmt"
	"reflect"
	"strings"
)

type (
	rowList []string
	argList []any
)

func Partial(in any) (rowList, argList) {
	args := argList{}
	rows := rowList{}

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		panic(fmt.Errorf("ToMap only accepts structs; got %T", v))
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fv := v.Field(i)
		if fv.IsZero() {
			continue
		}
		fi := typ.Field(i)
		tagv := fi.Tag.Get("db")
		switch tagv {
		case "-":
			continue
		default:
			if strings.Contains(tagv, ",") {
				tagv = strings.TrimSpace(strings.Split(tagv, ",")[0])
			}
			if len(tagv) == 0 {
				tagv = fi.Name
			}
			if tagv == "id" {
				continue
			}
			rows = append(rows, fmt.Sprintf("`%s`", tagv))
			args = append(args, fv.Interface())
		}
	}
	return rows, args
}

func (m *argList) WithId(id int64) *argList {
	args := append(*m, id)
	return &args
}

func (m *rowList) StringWithPlaceHolder() string {
	return strings.Join(*m, "=?,") + "=?"
}
