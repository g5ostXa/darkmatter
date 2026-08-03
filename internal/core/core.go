package core

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
)

// Common logger
var Logger = log.NewWithOptions(os.Stderr, log.Options{
	ReportTimestamp: false,
	Prefix:          ":",
})

// Time-stamped logger:
var TimeLogger = log.NewWithOptions(os.Stderr, log.Options{
	ReportTimestamp: true,
	Prefix:          ":",
})

func ClearScreen() {

	fmt.Printf("\033[2J\033[H")
}
