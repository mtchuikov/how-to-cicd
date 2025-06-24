package main

import (
	"context"
	"log/slog"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/generateUser", generateUserV1Handler)

	server := http.Server{
		Addr:         "127.0.0.1:7835",
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	slog.Info("listening http server on 127.0.0.1:7835...")

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			slog.Error(
				"failed to start http server",
				slog.String("err", err.Error()),
			)
			return
		}
	}()

	<-ctx.Done()
}
