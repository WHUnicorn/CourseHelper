package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"syscall"
	"testCourse/conf"
)

const (
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorPurple = "\033[35m"

	colorReset = "\033[0m"
)

var myLogger *log.Logger
var logLevel string
var logMap map[string]int
var isUnix = true

func init() {
	isUnix, _ = strconv.ParseBool(conf.Config.IsUnix)
	//switch runtime.GOOS {
	//case "windows":
	//	isUnix = false
	//default:
	//}

	myLogger = log.New(os.Stdout, "[Default]", log.Lshortfile|log.Ldate|log.Ltime)
	logLevel = conf.Config.LogLevel
	logMap = map[string]int{
		"error":   0,
		"warning": 1,
		"info":    2,
		"debug":   3,
	}

}

// 重写 log 的Println 方法，修改调用堆栈的追踪深度，以便调试
func overridePrintln(l *log.Logger, isDisplay bool, v ...any) {
	if !isDisplay {
		return
	}
	err := l.Output(4, fmt.Sprintln(v...))
	if err != nil {
		return
	}
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

func colorPrint(color string, msg string, v ...any) {
	if isUnix {
		myLogger.SetPrefix(color + "[" + msg + "]" + colorReset)
		overridePrintln(myLogger, logMap[logLevel] >= logMap[msg], color+fmt.Sprint(v...)+colorReset)
		// 上一行参数传v表示整体当成数组传入，参数传v... 表示多个参数分别传入
	} else {
		winColorPrintln(msg, fmt.Sprint(v...))
	}
}

func colorPrintf(format string, color string, msg string, v ...any) {
	if isUnix {
		myLogger.SetPrefix(color + "[" + msg + "]" + colorReset)
		overridePrintln(myLogger, logMap[logLevel] >= logMap[msg], color+fmt.Sprintf(format, v...)+colorReset)
	} else {
		winColorPrintln(msg, fmt.Sprintf(format, v...))
	}
}

func Debug(v ...any) {
	colorPrint(colorPurple, "Debug", v...)
}

func Info(v ...any) {
	colorPrint(colorGreen, "Info", v...)
}

func Warning(v ...any) {
	colorPrint(colorYellow, "Warning", v...)
}

func Error(v ...any) {
	colorPrint(colorRed, "Error", v...)
	os.Exit(1)
}

// DebugF 带格式化的调试日志
func DebugF(format string, v ...any) {
	colorPrintf(format, colorPurple, "Debug", v...)
}

// InfoF 带格式化的信息日志
func InfoF(format string, v ...any) {
	colorPrintf(format, colorGreen, "Info", v...)
}

// WarningF 带格式化的警告日志
func WarningF(format string, v ...any) {
	colorPrintf(format, colorYellow, "Warning", v...)
}

// ErrorF 带格式化的错误日志
func ErrorF(format string, v ...any) {
	colorPrintf(format, colorRed, "Error", v...)
	os.Exit(1)
}
