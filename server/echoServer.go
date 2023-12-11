package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/guatom999/sharePrice/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type (
	server struct {
		app *echo.Echo
		db  *gorm.DB
		cfg *config.Config
	}
)

func NewEchoServer(app *echo.Echo, db *gorm.DB, cfg *config.Config) IServer {
	return &server{
		app: app,
		db:  db,
		cfg: cfg,
	}
}

func (s *server) Start() {

	s.app.Use(middleware.Logger())

	close := make(chan os.Signal, 1)
	signal.Notify(close, syscall.SIGINT, syscall.SIGTERM)

	if err := s.app.Start(fmt.Sprintf(":%d", s.cfg.App.Port)); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to shutdown:%v", err)
	}

}
