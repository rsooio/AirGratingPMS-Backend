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

func UpdateKeys2x1_12(arg1, arg2, search1, search2 int64, f1 func(arg int64) string, f2 func(arg1, arg2 int64) string) *[]string {
	var keys []string

	if search1 == 0 {
		search1 = arg1
	}
	if search2 == 0 {
		search2 = arg2
	}

	if search1 != arg1 {
		keys = append(keys, f1(arg1))
		keys = append(keys, f1(search1))
	}

	if search2 != arg2 {
		keys = append(keys, f2(arg1, arg2))
		keys = append(keys, f2(search1, search2))
	}

	return &keys
}

func UpdateKeys1x12(cst, arg, search int64, f func(arg1, arg2 int64) string) *[]string {
	var keys []string

	if search == 0 {
		search = arg
	}

	if search != arg {
		keys = append(keys, f(cst, arg))
		keys = append(keys, f(cst, search))
	}

	return &keys
}

func UpdateKeys1x1(arg, search int64, f func(arg int64) string) *[]string {
	var keys []string

	if search == 0 {
		search = arg
	}

	if search != arg {
		keys = append(keys, f(arg))
		keys = append(keys, f(search))
	}

	return &keys
}
