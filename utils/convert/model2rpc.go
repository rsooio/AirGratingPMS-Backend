package convert

import (
	"database/sql"
	"reflect"
	"sync"
)

var skippingFields = map[string]bool{
	"CreateTime": true,
	"UpdateTime": true,
}

func Model2Rpc(model any, rpc any) error {
	rpcFields := rawFieldNames(rpc)
	modelFields := rawFieldNames(model)

	modelVal := reflect.ValueOf(model).Elem()
	rpcVal := reflect.ValueOf(rpc).Elem()

	wg := sync.WaitGroup{}
	for field := range modelFields {
		wg.Add(1)
		go func(field string) {
			if _, exist := skippingFields[field]; exist {
				wg.Done()
				return
			}

			_, exist := rpcFields[field]
			if !exist {
				wg.Done()
				return
			}

			typeName := modelFields[field]
			if typeName[:4] != "Null" {
				rpcVal.FieldByName(field).Set(modelVal.FieldByName(field))
			}

			switch typeName {
			case "NullFloat64":
				rpcVal.FieldByName(field).SetFloat(modelVal.FieldByName(field).Interface().(sql.NullFloat64).Float64)
			case "NullInt64":
				rpcVal.FieldByName(field).SetInt(modelVal.FieldByName(field).Interface().(sql.NullInt64).Int64)
			case "NullString":
				rpcVal.FieldByName(field).SetString(modelVal.FieldByName(field).Interface().(sql.NullString).String)
			}
			wg.Done()
		}(field)
	}

	// mr.ForEach(func(source chan<- interface{}) {
	// 	for field := range modelFields {
	// 		source <- field
	// 	}
	// }, func(item interface{}) {
	// 	field := item.(string)
	// 	if _, exist := skippingFields[field]; exist {
	// 		return
	// 	}

	// 	_, exist := rpcFields[field]
	// 	if !exist {
	// 		return
	// 	}

	// 	typeName := modelFields[field]
	// 	if typeName[:4] != "Null" {
	// 		rpcVal.FieldByName(field).Set(modelVal.FieldByName(field))
	// 	}

	// 	switch typeName {
	// 	case "NullFloat64":
	// 		rpcVal.FieldByName(field).SetFloat(modelVal.FieldByName(field).Interface().(sql.NullFloat64).Float64)
	// 	case "NullInt64":
	// 		rpcVal.FieldByName(field).SetInt(modelVal.FieldByName(field).Interface().(sql.NullInt64).Int64)
	// 	case "NullString":
	// 		rpcVal.FieldByName(field).SetString(modelVal.FieldByName(field).Interface().(sql.NullString).String)
	// 	}
	// })

	wg.Wait()
	return nil
}
