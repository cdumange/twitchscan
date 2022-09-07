package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/cdumange/twitchscan/ui"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", 8000),
		Handler: nil,
	}

	go func(ctx context.Context) {
		<-ctx.Done()
		server.Shutdown(context.Background())
	}(ctx)

	app.Route("/", new(ui.Front))
	app.RunWhenOnBrowser()

	http.Handle("/", &app.Handler{
		Name:        "test",
		Description: "first test",
		Styles: []string{
			"/web/css/global.css",
		},
	})

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
