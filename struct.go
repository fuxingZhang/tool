package utils

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// TrimStruct 去除结构体内字符串两边的空格
func TrimStruct(src interface{}) error {
	srcVal := reflect.ValueOf(src)
	if srcVal.Kind() != reflect.Ptr {
		return errors.New("paramter must be ptr")
	}
	val := srcVal.Elem()

	if val.Kind() != reflect.Struct {
		return errors.New("paramter must be struct")
	}

	var trim func(v reflect.Value)
	trim = func(v reflect.Value) {
		for i := 0; i < v.NumField(); i++ {
			filed := v.Field(i)
			switch filed.Kind() {
			case reflect.String:
				filed.SetString(strings.TrimSpace(filed.String()))
			case reflect.Struct:
				trim(filed)
			}
		}
	}

	trim(val)
	return nil
}

// CopyStruct 结构体拷贝
func CopyStruct(src, dst interface{}) error {
	s := reflect.Indirect(reflect.ValueOf(src))
	dstVal := reflect.ValueOf(dst)
	if dstVal.Kind() != reflect.Ptr {
		return errors.New("dst must be ptr")
	}
	d := dstVal.Elem()
	// t := reflect.TypeOf(src)

	for i := 0; i < s.NumField(); i++ {
		value := reflect.Indirect(s).Field(i)
		name := s.Type().Field(i).Name
		// name := t.Field(i).Name

		if value.Kind() == reflect.Ptr {
			value = value.Elem()
		}

		v := d.FieldByName(name)
		if v.Kind() == reflect.Ptr {
			v.Set(value.Addr())
			continue
		}
		if value.Kind() != v.Kind() {
			return fmt.Errorf("filed type unexpect, src %+v, dest %+v", src, dst)
		}
		if v.IsValid() && v.CanSet() {
			v.Set(value)
		}
	}
	return nil
}
