package logger

import (
	"strings"
	"syscall"
)

func beforeOut(levelStr string, _ *string) {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	proc := kernel32.NewProc("SetConsoleTextAttribute")
	switch strings.ToLower(levelStr) {
	case "error":
		_, _, _ = proc.Call(uintptr(syscall.Stdout), uintptr(4))
	case "warning":
		_, _, _ = proc.Call(uintptr(syscall.Stdout), uintptr(6))
	case "info":
		_, _, _ = proc.Call(uintptr(syscall.Stdout), uintptr(2))
	case "debug":
		_, _, _ = proc.Call(uintptr(syscall.Stdout), uintptr(5))
	default:
		_, _, _ = proc.Call(uintptr(syscall.Stdout), uintptr(5))
	}
}

func afterOut() {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	proc := kernel32.NewProc("SetConsoleTextAttribute")
	_, _, _ = proc.Call(uintptr(syscall.Stdout), uintptr(7))
}
