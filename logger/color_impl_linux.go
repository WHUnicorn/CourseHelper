package logger

import (
	"strings"
)

const (
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorPurple = "\033[35m"

	colorReset = "\033[0m"
)

func beforeOut(levelStr string, outStr *string) {
	switch strings.ToLower(levelStr) {
	case "error":
		*outStr = colorRed + *outStr + colorReset
	case "warning":
		*outStr = colorYellow + *outStr + colorReset
	case "info":
		*outStr = colorGreen + *outStr + colorReset
	case "debug":
		*outStr = colorPurple + *outStr + colorReset
	default:
		*outStr = colorPurple + *outStr + colorReset
	}
}

func afterOut() {
	// do nothing
}
