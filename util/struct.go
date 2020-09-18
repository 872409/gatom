package util

import (
	"fmt"
	"reflect"
)

// func CopyStruct(src, dst interface{}) {
// 	sval := reflect.ValueOf(src).Elem()
// 	dval := reflect.ValueOf(dst).Elem()
//
// 	for i := 0; i < sval.NumField(); i++ {
// 		value := sval.Field(i)
// 		name := sval.Type().Field(i).Name
//
// 		dvalue := dval.FieldByName(name)
// 		if dvalue.IsValid() == false {
// 			continue
// 		}
// 		dvalue.Set(value)
// 	}
// }


// 用src的所有字段覆盖dst
// 如果fields不为空, 表示用b的特定字段覆盖a的
// dst应该为结构体指针
func CopyStruct(src interface{}, dst interface{}, fields ...string) (err error) {
	srcT := reflect.TypeOf(src)
	srcV := reflect.ValueOf(src)

	dstT := reflect.TypeOf(dst)
	dstV := reflect.ValueOf(dst)

	// 简单判断下
	if dstT.Kind() != reflect.Ptr {
		err = fmt.Errorf("a must be a struct pointer")
		return
	}
	dstV = reflect.ValueOf(dstV.Interface())

	// 要复制哪些字段
	_fields := make([]string, 0)
	if len(fields) > 0 {
		_fields = fields
	} else {
		for i := 0; i < srcV.NumField(); i++ {
			_fields = append(_fields, srcT.Field(i).Name)
		}
	}

	if len(_fields) == 0 {
		fmt.Println("no fields to copy")
		return
	}

	// 复制
	for i := 0; i < len(_fields); i++ {
		name := _fields[i]
		dstF := dstV.Elem().FieldByName(name)
		srcValue := srcV.FieldByName(name)

		// a中有同名的字段并且类型一致才复制
		if dstF.IsValid() && dstF.Kind() == srcValue.Kind() {
			dstF.Set(srcValue)
		} else {
			fmt.Printf("no such field or different kind, fieldName: %s\n", name)
		}
	}
	return
}
