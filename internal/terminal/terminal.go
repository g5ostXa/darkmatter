package terminal

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func ClearScreen() {

	if runtime.GOOS == "linux" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		_ = cmd.Run()
		return
	}
	fmt.Printf("\033[2J\033[H")
}
