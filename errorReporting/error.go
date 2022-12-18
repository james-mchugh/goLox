// Error reporting utilities

package errorReporting

import (
	"fmt"
	"os"
)

var errorOccurred = false

type logLevel = string

func Error(line int, message string) {
	report(line, "", message, "Error")
	errorOccurred = true
}

func ErrorOccurred() bool {
	return errorOccurred
}

func report(line int, s string, message string, level logLevel) {
	_, _ = fmt.Fprintf(os.Stderr, "[line %d] %s %s: %s\n", line, s, message, level)
}
