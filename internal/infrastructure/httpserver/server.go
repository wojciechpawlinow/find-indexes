package httpserver

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/wojciechpawlinow/find-indexes/internal/config"
	"github.com/wojciechpawlinow/find-indexes/pkg/logger"
)

type Server struct {
	*http.Server
}

func Run(c config.Provider) *Server {
	router := gin.Default()
	router.GET("/find/:value", func(c *gin.Context) {
		v := c.Param("value")
		c.String(http.StatusOK, "%s", v)
	})

	s := &Server{
		&http.Server{
			Addr:    fmt.Sprintf(":%s", c.GetString("port")),
			Handler: router,
		},
	}

	go func() {
		if err := s.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("shutting down the server: %s", err)
		}
	}()
	_, _ = fmt.Printf("\nlistening at localhost:%s\n", c.GetString("port"))

	return s
}

// Shutdown is a Shutdown function overload in order to trigger a different Shutdown hooks
func (srv *Server) Shutdown(ctx context.Context) error {
	if err := srv.Server.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}
