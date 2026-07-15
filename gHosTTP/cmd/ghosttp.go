package gHosTTP

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"charm.land/huh/v2"
	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/tree"
	"github.com/g5ostXa/darkmatter/internal/core"
	"github.com/g5ostXa/darkmatter/internal/styles"
)

// HTTP service status message
var (
	serviceServeMsg = "→ Serving static files from"
	serviceAddrMsg  = "→ Server is running at"
	serviceAddrUrl  = "http://localhost:8080 / 127.0.0.1:8080"
)

func makeServeTree(path string) {

	t := tree.Root(styles.TreeRootStyle.Render(serviceServeMsg)).
		Child(
			tree.New().
				Root(styles.HttpChildStyle.Render(path)),
		)
	lipgloss.Println(t)
}

func makeAddrTree() {

	t := tree.Root(styles.TreeRootStyle.Render(serviceAddrMsg)).
		Child(
			tree.New().
				Root(styles.HttpChildStyle.Render(serviceAddrUrl)),
		)
	lipgloss.Println(t)
}

func getWebsitePath() (string, error) {

	var dirPath string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Enter website directory path").
				Description("Provide the full path to your website files").
				Placeholder("/home/user/my-website").
				Value(&dirPath).
				Validate(func(s string) error {
					info, err := os.Stat(s)
					if err != nil {
						return fmt.Errorf("path does not exist: %v", err)
					}
					if !info.IsDir() {
						return fmt.Errorf("path is not a directory")
					}
					return nil
				}),
		),
	)

	if err := form.Run(); err != nil {
		return "", err
	}

	return dirPath, nil
}

func Serve() {

	dirPath, err := getWebsitePath()
	if err != nil {
		core.TimeLogger.Fatal("Failed to get directory path")
	}

	fs := http.FileServer(http.Dir(dirPath))

	mux := http.NewServeMux()
	mux.Handle("/", fs)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {

		makeServeTree(dirPath)
		fmt.Println()

		makeAddrTree()
		fmt.Println()

		lipgloss.Println(styles.LegendStyle.Render("[CTRL + C to stop server]"))

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			core.TimeLogger.Fatal(":: Listen / Serve error...")
		}
	}()

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		core.TimeLogger.Fatal(":: Failed to terminate server...")
	}

	core.TimeLogger.Info(":: Server terminated successfully!")
}
