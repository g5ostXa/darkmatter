package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/tree"
	"github.com/g5ostXa/darkmatter/internal/core"
	"github.com/g5ostXa/darkmatter/internal/logger"
	"github.com/g5ostXa/darkmatter/internal/styles"
	"github.com/g5ostXa/darkmatter/internal/terminal"
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

	fmt.Println()
}

func main() {

	showVersion := flag.Bool("version", false, "print current version")

	flag.Parse()
	if *showVersion {
		fmt.Println(mainTitle, strings.TrimSpace(version))
		os.Exit(0)
	}

	terminal.ClearScreen()
	logger.TimeLogger.Info("Initializing...")

	terminal.ClearScreen()
	RenderHeader()

	core.Menu()
	terminal.ClearScreen()

	RenderHeader()
}
