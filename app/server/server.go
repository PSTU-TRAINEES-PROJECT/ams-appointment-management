package server

import (
	"ams-appointment-management/app/common/logger"
	"ams-appointment-management/app/config"
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

type Server struct {
	framework *echo.Echo
	config    *config.Config
}

func New(config config.Config, framework *echo.Echo) *Server {
	corsConfig := middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.PATCH, echo.HEAD, echo.OPTIONS},
		AllowHeaders:     []string{"*"},
		AllowCredentials: false,
	}

	framework.Use(middleware.CORSWithConfig(corsConfig))
	return &Server{
		framework: framework,
		config:    &config,
	}
}

func (s *Server) Start() {
	e := s.framework
	// start http server
	go func() {
		e.Start(":" + strconv.Itoa(s.config.App.ServerPort))
	}()

	// graceful shutdown
	s.GracefulShutdown()
}

// GracefulShutdown server will gracefully shut down within 5 sec
func (s *Server) GracefulShutdown() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT)
	<-ch
	logger.Info("shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = s.framework.Shutdown(ctx)
	logger.Info("server shutdowns gracefully")
}
