package main

import (
	"context"
	"log"
	"maribowman/signal-transmitter/app"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	server, err := app.InitServer()
	if err != nil {
		log.Fatalf("failed to initialize server: %v", err)
	}

	// graceful shutdown - https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/server.go
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to initialize server: %v", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	log.Println("shutting down server...")
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
	}
}
