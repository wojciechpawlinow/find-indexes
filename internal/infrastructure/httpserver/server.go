package httpserver

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/sarulabs/di"

	"github.com/wojciechpawlinow/find-indexes/internal/config"
	"github.com/wojciechpawlinow/find-indexes/internal/infrastructure/httpserver/handlers"
)

type Server struct {
	*http.Server
}

// Run is a Server constructor that starts the HTTP server in a goroutine and enables routing
func Run(cfg config.Provider, ctn di.Container, errChan chan error) *Server {

	// get handlers
	indexHandler := ctn.Get("http-index").(*handlers.IndexHTTPHandler) // no interface here, exact dependency

	// define routes
	router := gin.Default()
	router.GET("/index/:value", indexHandler.FindIndex)

	s := &Server{
		&http.Server{
			Addr:    fmt.Sprintf(":%s", cfg.GetString("port")),
			Handler: router,
		},
	}

	go func() {
		if err := s.Server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errChan <- err
		}
	}()

	_, _ = color.New(color.FgHiGreen).Printf("\n=> an HTTP server listening at: %s\n\n", cfg.GetString("port"))

	return s
}

// Shutdown is a Shutdown function overload
func (srv *Server) Shutdown(ctx context.Context) error {
	return srv.Server.Shutdown(ctx)
}
