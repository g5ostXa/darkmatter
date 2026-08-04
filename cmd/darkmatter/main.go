package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"charm.land/huh/v2"
	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/tree"
	"github.com/g5ostXa/darkmatter/internal/core"
	"github.com/g5ostXa/darkmatter/internal/getarch"
	ghosttp "github.com/g5ostXa/darkmatter/internal/ghosttp/cmd"
	"github.com/g5ostXa/darkmatter/internal/glyphs"
	"github.com/g5ostXa/darkmatter/internal/password_generator"
	"github.com/g5ostXa/darkmatter/internal/styles"
)

var (
	mainTitle = "d󰣇rkmatter 󱡕"
	version   = "dev"
)

func makeTree() {

	t := tree.Root(styles.TreeRootStyle.Render("○ Version")).
		Child(
			tree.New().
				Root(styles.TreeChildStyle.Render(version)),
		)
	lipgloss.Println(t)
}

func RenderHeader() {

	lipgloss.Println(styles.HeaderStyle.Render("", mainTitle, ""))
	makeTree()
}

func main() {

	showVersion := flag.Bool("version", false, "print current version")

	flag.Parse()
	if *showVersion {
		fmt.Println(mainTitle, strings.TrimSpace(version))
		os.Exit(0)
	}

	core.ClearScreen()
	core.TimeLogger.Info("Initializing...")

	core.ClearScreen()
	RenderHeader()

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
			core.Logger.Fatalf("Error running form: %v", err)
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

		core.ClearScreen()
		RenderHeader()

		fmt.Println()
	}
}
