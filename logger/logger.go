package logger

import (
	"fmt"
	"github.com/WHUCSStudy/StudyBot/setup"
	"log"
	"os"
	"runtime"
	"strconv"
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
	myLogger = log.New(os.Stdout, "[Default]", log.Lshortfile|log.Ldate|log.Ltime)
	logLevel = setup.Config.LogLevel
	logMap = map[string]int{
		"error":   0,
		"warning": 1,
		"info":    2,
		"debug":   3,
	}

	switch setup.Config.IsUnix {
	case "default":
		switch runtime.GOOS {
		case "windows":
			isUnix = false
		default:
		}
	case "true", "false":
		isUnix, _ = strconv.ParseBool(setup.Config.IsUnix)
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
