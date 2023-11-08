package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fatih/color"

	"github.com/wojciechpawlinow/find-indexes/internal/config"
	"github.com/wojciechpawlinow/find-indexes/internal/infrastructure/container"
	"github.com/wojciechpawlinow/find-indexes/internal/infrastructure/httpserver"
	"github.com/wojciechpawlinow/find-indexes/pkg/logger"
)

func main() {

	// load configuration from a file or fallback to defaults
	cfg, err := config.Load()
	if err != nil {
		_, _ = color.New(color.FgYellow).Printf("\n=> no configuration file found, using defaults\n\n")
	}

	// create logger
	logger.Setup(cfg)

	// build dependencies
	ctn := container.New()

	errChan := make(chan error, 1)

	// create and run HTTP server
	s := httpserver.Run(cfg, ctn, errChan)

	// wait for interrupt signal to gracefully shut down the server with a timeout of 10 seconds
	// use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, syscall.SIGTERM)

	select {
	case <-quit:
		logger.Info("initiating graceful shutdown...")
	case err = <-errChan:
		logger.Error("server error: ", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// gracefully shut down the server
	if err = s.Shutdown(ctx); err != nil {
		logger.Error("server shutdown failed: ", err)
	}
}
