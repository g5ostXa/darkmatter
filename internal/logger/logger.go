package logger

import (
	"io"
	"os"

	"github.com/charmbracelet/log"
)

var (
	Standard = log.NewWithOptions(os.Stderr, log.Options{
		ReportTimestamp: false,
		Prefix:          ":",
	})

	TimeStamped = log.NewWithOptions(os.Stderr, log.Options{
		ReportTimestamp: true,
		Prefix:          ":",
	})
)

func LogOutput(w io.Writer) {

	if w == nil {
		w = os.Stderr
	}
	Standard = log.NewWithOptions(os.Stderr, log.Options{
		ReportTimestamp: false,
		Prefix:          ":",
	})

	TimeStamped = log.NewWithOptions(os.Stderr, log.Options{
		ReportTimestamp: true,
		Prefix:          ":",
	})
}
