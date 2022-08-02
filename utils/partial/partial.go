package partial

import (
	"fmt"
	"reflect"
	"strings"
)

func Partial(in any) ([]string, []any) {
	args := []any{}
	rows := []string{}

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

func RowsWithPlaceHolder(rows []string) string {
	return strings.Join(rows, "=?,") + "=?"
}
