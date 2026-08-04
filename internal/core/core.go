package core

import (
	"fmt"
	"os"

	"charm.land/huh/v2"
	"charm.land/lipgloss/v2"
	"github.com/g5ostXa/darkmatter/internal/getarch"
	ghosttp "github.com/g5ostXa/darkmatter/internal/ghosttp/cmd"
	"github.com/g5ostXa/darkmatter/internal/glyphs"
	"github.com/g5ostXa/darkmatter/internal/logger"
	"github.com/g5ostXa/darkmatter/internal/password_generator"
	"github.com/g5ostXa/darkmatter/internal/styles"
)

func Menu() {

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
						huh.NewOption("Password generator", "opt4"),
						huh.NewOption("Exit", "exit"),
					).
					Value(&choice),
			),
		)

		err := form.Run()

		if err != nil {
			logger.Standard.Fatalf("Error running form: %v", err)
		}

		switch choice {
		case "opt1":
			ghosttp.Serve()
		case "opt2":
			getarch.Latest()
		case "opt3":
			glyphs.Pager()
		case "opt4":
			password_generator.Gen()
		case "exit":
			lipgloss.Println(styles.CommonStyle.Render("\nExiting..."))
			os.Exit(0)
		}

		fmt.Print("\nPress Enter to return to the menu...")
		fmt.Scanln()
	}
}
