package utils

import "runtime"

// GetRuntimeStack 获取堆栈信息
func GetRuntimeStack() string {
	buf := make([]byte, 1<<16)
	n := runtime.Stack(buf, false)
	return string(buf[:n])
}
