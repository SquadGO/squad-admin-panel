package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func New(r http.Handler) {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("BACKEND_PORT")),
		Handler: r,
	}

	go func() {
		slog.Info("Starting server", "port", server.Addr)

		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			slog.Error("HTTP server:", slog.Any("err", err))
		}

		slog.Info("Stopped serving new connections.")
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("HTTP shutdown:", slog.Any("err", err))
	}

	slog.Info("Graceful shutdown complete.")
}
