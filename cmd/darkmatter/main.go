package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"charm.land/huh/v2"
	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/tree"
	"github.com/g5ostXa/darkmatter/internal/core"
	"github.com/g5ostXa/darkmatter/internal/getarch"
	ghosttp "github.com/g5ostXa/darkmatter/internal/ghosttp/cmd"
	"github.com/g5ostXa/darkmatter/internal/glyphs"
	"github.com/g5ostXa/darkmatter/internal/styles"
)

var (
	mainTitle     = "⋯ d󰣇rkm󰣇tter ⋯"
	latestVersion = "v0.1.6"
)

func makeTree() {

	t := tree.Root(styles.TreeRootStyle.Render("○ Version")).
		Child(
			tree.New().
				Root(styles.TreeChildStyle.Render(latestVersion)),
		)
	lipgloss.Println(t)
}

func RenderHeader() {

	lipgloss.Println(styles.HeaderStyle.Render("", mainTitle, ""))
	makeTree()
}

func main() {

	core.ClearScreen()

	core.TimeLogger.Info("Initializing...")
	time.Sleep(2 * time.Second)

	core.ClearScreen()
	RenderHeader()

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

		err := form.Run()
		if err != nil {
			log.Fatalf("Error running form: %v", err)
		}

		switch choice {
		case "opt1":
			ghosttp.Serve()
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
		RenderHeader()

		fmt.Println()
	}
}
