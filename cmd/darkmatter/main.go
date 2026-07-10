package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"charm.land/huh/v2"
	"charm.land/lipgloss/v2"
	gHosTTP "github.com/g5ostXa/darkmatter/gHosTTP/cmd"
	"github.com/g5ostXa/darkmatter/getarch"
	"github.com/g5ostXa/darkmatter/glyphs"
	"github.com/g5ostXa/darkmatter/header"
	"github.com/g5ostXa/darkmatter/internal/core"
	"github.com/g5ostXa/darkmatter/internal/styles"
)

func main() {

	core.ClearScreen()
	core.TimeLogger.Info("Initializing...")

	time.Sleep(2 * time.Second)

	core.ClearScreen()
	header.Render()

	time.Sleep(500 * time.Millisecond)
	fmt.Println()

	var choice string

	for {
		form := huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().
					Title("Main Menu").
					Description("Choose an option to execute, or exit the program.").
					Options(
						huh.NewOption("Initiate local HTTP server", "opt1"),
						huh.NewOption("Get latest archiso and sig", "opt2"),
						huh.NewOption("Glyphs menu", "opt3"),
						huh.NewOption("Exit", "exit"),
					).
					Value(&choice),
			),
		)

		// Run the form
		err := form.Run()
		if err != nil {
			log.Fatalf("Error running form: %v", err)
		}

		// Handle user selection
		switch choice {
		case "opt1":
			gHosTTP.Serve()
		case "opt2":
			getarch.Latest()
		case "opt3":
			glyphs.Pager()
		case "exit":
			lipgloss.Println(styles.CommonStyle.Render("\nExiting..."))
			os.Exit(0)
		}

		fmt.Print("\nPress Enter to return to the menu...")
		fmt.Scanln()

		core.ClearScreen()
		header.Render()

		fmt.Println()
	}
}
