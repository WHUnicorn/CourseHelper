package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

type MyWriter struct {
}

func writeToFile(b []byte) {
	file, err := os.OpenFile("logs/text.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeDir|os.ModePerm)
	if err != nil {
		fmt.Println("writeToFileErr: ", err)
		return
	}
	_, err = file.Write(b)
	if err != nil {
		fmt.Println("writeToFileErr: ", err)
		return
	}
}

// Write 实现 io.Writer 接口的 Write 方法
func (cw *MyWriter) Write(p []byte) (n int, err error) {
	// 外层函数已加锁

	// 输出前
	compile := regexp.MustCompile(`(^\[\S+])`)
	levelStr := strings.TrimFunc(compile.FindString(string(p)), func(r rune) bool {
		return r == '[' || r == ']'
	})
	outStr := string(p)
	writeToFile(p)
	beforeOut(levelStr, &outStr)

	// 调用底层输出
	n, err = os.Stdout.Write([]byte(outStr))

	// 输出后
	afterOut()

	return n, err
}

type MyFormatter struct{}

// Format 格式化日志条目
func (f *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {

	var filePath string
	var line int
	if entry.Level == logrus.TraceLevel {
		// 经测试，gorm 的追踪栈深度为11
		_, filePath, line, _ = runtime.Caller(11) // 调整栈深度
	} else {
		_, filePath, line, _ = runtime.Caller(7) // 调整栈深度
		if strings.Contains(filePath, "utils/logger/logger.go") {
			_, filePath, line, _ = runtime.Caller(8) // 调整栈深度
		}
	}

	//functionName := runtime.FuncForPC(pc).Name()
	//fmt.Printf(functionName)

	rootPath, _ := os.Getwd()
	//fmt.Printf(rootPath)
	// 去掉根路径和最前面的 /
	file := filePath[len(rootPath)+1:]

	// 构建日志格式
	level := strings.ToUpper(entry.Level.String())
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	message := entry.Message

	// 构建输出字符串
	formatted := "[" + level + "] " + timestamp + " " + file + ":" + strconv.Itoa(line) + ": " + message + "\n"
	return []byte(formatted), nil
}

var myLogger = logrus.New()

// 仅显示时间，但写入文件
func init() {
	myLogger.SetFormatter(&MyFormatter{})
	myLogger.SetOutput(&MyWriter{})
	myLogger.SetLevel(logrus.TraceLevel)
}

func Debug(v ...any) {
	myLogger.Debug(v...)
}

func Info(v ...any) {
	myLogger.Info(v...)
}

func Warning(v ...any) {
	myLogger.Warning(v...)
}

func Error(v ...any) {
	myLogger.Error(v...)
	//os.Exit(1)
}

func Trace(v ...any) {
	myLogger.Trace(v...)
}

// DebugF 带格式化的调试日志
func DebugF(format string, v ...any) {
	myLogger.Debugf(format, v...)
}

// InfoF 带格式化的信息日志
func InfoF(format string, v ...any) {
	myLogger.Infof(format, v...)

}

// WarningF 带格式化的警告日志
func WarningF(format string, v ...any) {
	myLogger.Warningf(format, v...)
}

// ErrorF 带格式化的错误日志
func ErrorF(format string, v ...any) {
	myLogger.Errorf(format, v...)
	//os.Exit(1)
}

// TraceF 带格式化的追踪日志
func TraceF(format string, v ...any) {
	myLogger.Tracef(format, v...)
}
