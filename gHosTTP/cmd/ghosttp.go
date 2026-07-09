package gHosTTP

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/tree"
	"github.com/g5ostXa/darkmatter/internal/core"
	"github.com/g5ostXa/darkmatter/internal/styles"
)

// HTTP service status message
var (
	serviceServeMsg = "→ Serving static files from"
	serviceAddrMsg  = "→ Server is running at"

	serviceServeDir = "./static/ and ./static/website"
	serviceAddrUrl  = "http://localhost:8080 / 127.0.0.1:8080"
)

func makeServeTree() {

	t := tree.Root(styles.TreeRootStyle.Render(serviceServeMsg)).
		Child(
			tree.New().
				Root(styles.HttpChildStyle.Render(serviceServeDir)),
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

func Serve() {

	fs := http.FileServer(http.Dir("./static"))

	mux := http.NewServeMux()
	mux.Handle("/", fs)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {

		makeServeTree()
		fmt.Println()

		makeAddrTree()
		fmt.Println()

		lipgloss.Println(styles.LegendStyle.Render("[CTRL + C to stop server]"))

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			core.TimeLogger.Fatal("Listen / Serve error...")
		}
	}()

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// TODO: Add error message
	if err := srv.Shutdown(ctx); err != nil {
		core.TimeLogger.Fatal("")
	}

	core.TimeLogger.Info("Server terminated successfully!")
}
