//go:build !linux

package logger

import (
	"fmt"
	"strings"
	"syscall"
)

type colorLoggerImpl int

func colorPrint(color string, msg string, v ...any) {
	winColorPrintln(msg, fmt.Sprint(v...))

}
func colorPrintf(format string, color string, msg string, v ...any) {
	winColorPrintln(msg, fmt.Sprintf(format, v...))
}

// windows 操作系统下调用系统api实现cmd着色
func winColorPrintln(msg string, v ...any) {

	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	proc := kernel32.NewProc("SetConsoleTextAttribute")
	switch strings.ToLower(msg) {
	case "error":
		_, _, _ = proc.Call(uintptr(syscall.Stdout), uintptr(4))
	case "warning":
		_, _, _ = proc.Call(uintptr(syscall.Stdout), uintptr(6))
	case "info":
		_, _, _ = proc.Call(uintptr(syscall.Stdout), uintptr(2))
	case "debug":
		_, _, _ = proc.Call(uintptr(syscall.Stdout), uintptr(5))
	default:
	}

	myLogger.SetPrefix("[" + msg + "]")
	overridePrintln(myLogger, logMap[logLevel] >= logMap[msg], v...)

	_, _, _ = proc.Call(uintptr(syscall.Stdout), uintptr(7))
}
