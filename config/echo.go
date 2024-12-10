package config

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func NewEcho() (*echo.Echo, *logrus.Logger) {
	e := echo.New()
	log := logrus.New()

	return e, log
}

// GracefulShutdown handles the graceful shutdown of the Echo server
func GracefulShutdown(e *echo.Echo, log *logrus.Logger, shutdownTimeout time.Duration) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatalf("Error shutting down server: %v", err)
	}

	log.Println("Server gracefully stopped")
}
