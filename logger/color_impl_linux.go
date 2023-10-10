//go:build !linux

package logger

func colorPrint(color string, msg string, v ...any) {
	myLogger.SetPrefix(color + "[" + msg + "]" + colorReset)
	overridePrintln(myLogger, logMap[logLevel] >= logMap[msg], color+fmt.Sprint(v...)+colorReset)
	// 上一行参数传v表示整体当成数组传入，参数传v... 表示多个参数分别传入
}
func colorPrintf(format string, color string, msg string, v ...any) {
	myLogger.SetPrefix(color + "[" + msg + "]" + colorReset)
	overridePrintln(myLogger, logMap[logLevel] >= logMap[msg], color+fmt.Sprintf(format, v...)+colorReset)
}
