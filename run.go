package qstnnr

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"sync"

	"github.com/mateopresacastro/qstnnr/pkg/qservice"
	"github.com/mateopresacastro/qstnnr/pkg/server"
	"github.com/mateopresacastro/qstnnr/pkg/store"
	"google.golang.org/grpc"
)

const DefaultPort = "5974"

func Run(
	ctx context.Context,
	getenv func(string) string,
	output io.Writer,
) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	data := getInitialData()
	store, err := store.NewInMemory(data)
	if err != nil {
		return err
	}

	service := qservice.New(store)
	logger := slog.New(slog.NewJSONHandler(output, &slog.HandlerOptions{
		Level: parseLogLevel(getenv("LOG_LEVEL")),
	}))

	cfg := &server.Config{
		Logger:  logger,
		Service: service,
	}

	server, err := server.New(cfg)
	if err != nil {
		return fmt.Errorf("creating server: %w", err)
	}

	port := getenv("PORT")
	if port == "" {
		port = DefaultPort
	}

	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return fmt.Errorf("listening on port %s: %w", port, err)
	}

	go func() {
		logger.Info("listening", "port", getenv("PORT"))
		if err := server.Serve(ln); err != nil && err != grpc.ErrServerStopped {
			fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		logger.Info("shutting down server")
		server.GracefulStop()
	}()

	wg.Wait()
	return nil
}

func parseLogLevel(level string) slog.Level {
	switch level {
	case "DEBUG":
		return slog.LevelDebug
	case "INFO":
		return slog.LevelInfo
	case "WARN":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
