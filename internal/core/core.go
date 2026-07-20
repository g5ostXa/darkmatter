package core

import (
	"fmt"
	"os"
	"os/exec"

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

	clr := exec.Command("clear")

	output, err := clr.Output()
	if err != nil {
		fmt.Print("\033[2J\033[H")
	}
	fmt.Println(string(output))
}
