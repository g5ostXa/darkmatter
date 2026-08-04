package logger

import (
	"io"
	"os"

	"github.com/charmbracelet/log"
)

var (
	Logger = log.NewWithOptions(os.Stderr, log.Options{
		ReportTimestamp: false,
		Prefix:          ":",
	})

	TimeLogger = log.NewWithOptions(os.Stderr, log.Options{
		ReportTimestamp: true,
		Prefix:          ":",
	})
)

func LogOutput(w io.Writer) {

	if w == nil {
		w = os.Stderr
	}
	Logger = log.NewWithOptions(os.Stderr, log.Options{
		ReportTimestamp: false,
		Prefix:          ":",
	})

	TimeLogger = log.NewWithOptions(os.Stderr, log.Options{
		ReportTimestamp: true,
		Prefix:          ":",
	})
}
