package dgsys

import "runtime"

func IsMacOS() bool {
	return runtime.GOOS == "darwin"
}

func IsWindows() bool {
	return runtime.GOOS == "windows"
}

func IsLinux() bool {
	return runtime.GOOS == "linux"
}
