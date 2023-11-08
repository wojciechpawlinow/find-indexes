package httpserver

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sarulabs/di"

	"github.com/wojciechpawlinow/find-indexes/internal/config"
	"github.com/wojciechpawlinow/find-indexes/internal/infrastructure/httpserver/handlers"
	"github.com/wojciechpawlinow/find-indexes/pkg/logger"
)

type Server struct {
	*http.Server
}

func Run(cfg config.Provider, ctn di.Container) *Server {

	indexHTTPHandler := ctn.Get("http-index").(*handlers.IndexFinderHTTPHandler)

	router := gin.Default()
	router.GET("/find/:value", indexHTTPHandler.FindIndex)

	s := &Server{
		&http.Server{
			Addr:    fmt.Sprintf(":%s", cfg.GetString("port")),
			Handler: router,
		},
	}

	go func() {
		if err := s.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("shutting down the server: %s", err)
		}
	}()
	_, _ = fmt.Printf("\nlistening at localhost:%s\n", cfg.GetString("port"))

	return s
}

// Shutdown is a Shutdown function overload in order to trigger a different Shutdown hooks
func (srv *Server) Shutdown(ctx context.Context) error {
	if err := srv.Server.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}
