package header

import (
	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/tree"
	"github.com/g5ostXa/darkmatter/internal/styles"
)

var (
	mainTitle     = "⋯ d󰣇rkm󰣇tter ⋯"
	latestVersion = "v0.1.1"
)

func makeTree() {

	t := tree.Root(styles.TreeRootStyle.Render("○ Version")).
		Child(
			tree.New().
				Root(styles.TreeChildStyle.Render(latestVersion)),
		)
	lipgloss.Println(t)
}

func Render() {

	lipgloss.Println(styles.HeaderStyle.Render("", mainTitle, ""))
	makeTree()
}
