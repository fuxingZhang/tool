package utils

import (
	"reflect"
	"runtime"
	"strings"
)

// GetFuncName 获取函数名
func GetFuncName(fn interface{}) string {
	name := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	i := strings.LastIndex(name, ".") + 1
	return name[i:]
}

// GetFuncNameWithPath 获取函数名，带路径
func GetFuncNameWithPath(fn interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
}
