package utils

import (
	"reflect"
	"runtime"
)

// 获取函数名称的工具函数
func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
