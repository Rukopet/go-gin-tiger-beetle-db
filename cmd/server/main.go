package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"finance-api/internal/config"
	"finance-api/internal/handlers"
	"finance-api/internal/services"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)

	cfg := config.Load()

	accountService := services.NewAccountService()
	transferService := services.NewTransferService()
	statsService := services.NewStatsService()

	handlersInstance := handlers.NewHandlers(
		accountService,
		transferService,
		statsService,
	)

	router := handlers.SetupRouter(handlersInstance)

	address := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	server := &http.Server{
		Addr:    address,
		Handler: router,
	}

	go func() {
		logrus.WithField("address", address).Info("Starting server")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.WithError(err).Fatal("Failed to start server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logrus.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logrus.WithError(err).Fatal("Server forced to shutdown")
	}

	logrus.Info("Server exited")
}
