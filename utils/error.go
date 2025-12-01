package utils

import (
	"runtime"
	"strings"
)

// GetRuntimeStack 获取堆栈信息
func GetRuntimeStack() string {
	buf := make([]byte, 1<<16)
	n := runtime.Stack(buf, false)
	return string(buf[:n])
}

// GetSimpleRuntimeStack 获取精简堆栈信息
func GetSimpleRuntimeStack() string {
	stack := GetRuntimeStack()

	var filteredLines []string
	lines := strings.Split(stack, "\n")
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine != "" && !strings.HasPrefix(trimmedLine, "/go/pkg") && !strings.HasPrefix(trimmedLine, "github.com") {
			filteredLines = append(filteredLines, line)
		}
	}

	if len(filteredLines) == 0 {
		return ""
	}

	return strings.Join(filteredLines, "\n")
}
