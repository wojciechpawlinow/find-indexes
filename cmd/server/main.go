package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/wojciechpawlinow/find-indexes/internal/config"
	"github.com/wojciechpawlinow/find-indexes/internal/infrastructure/container"
	"github.com/wojciechpawlinow/find-indexes/internal/infrastructure/httpserver"
	"github.com/wojciechpawlinow/find-indexes/pkg/logger"
)

func main() {

	// create config from file or fallback to defaults
	cfg, err := config.New()
	if err != nil {
		fmt.Println("loaded default configuration")
	}

	// create logger
	logger.Setup(cfg)

	// build dependencies
	ctn := container.New()

	// create and run HTTP server
	s := httpserver.Run(cfg, ctn)

	// wait for interrupt signal to gracefully shut down the server with a timeout of 10 seconds
	// use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, syscall.SIGTERM)

	<-quit
	logger.Info("initiating graceful shutdown...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// gracefully shut down the server
	if err := s.Shutdown(ctx); err != nil {
		logger.Fatalf("failed shutting down the server: %w", err)
	}
}
