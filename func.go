package tool

import (
	"reflect"
	"runtime"
	"strings"
)

// GetFuncName get function name
func GetFuncName(fn interface{}) string {
	name := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	i := strings.LastIndex(name, ".") + 1
	return name[i:]
}

// GetFuncNameWithPath get function name with filepath
func GetFuncNameWithPath(fn interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
}
