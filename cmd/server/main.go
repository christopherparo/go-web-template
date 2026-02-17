package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/christopherparo/go-web-template/internal/server"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := server.New(port)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := srv.Start(ctx); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
