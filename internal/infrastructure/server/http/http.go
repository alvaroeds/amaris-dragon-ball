package http

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// Server is a base server configuration.
type server struct {
	*http.Server
}

// NewServer initialize a new server with configuration.
func NewServer(listening string, mux http.Handler) *server {
	s := &http.Server{
		Addr:         ":" + listening,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &server{s}
}

// Start runs ListenAndServe on the http.Server with graceful shutdown.
func (srv *server) Start() {
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("❌ Could not listen on %s due to: %v\n", srv.Addr, err)
		}
	}()
	fmt.Printf("🚀 Dragon Ball API server is ready to handle requests on %s\n", srv.Addr)
	srv.gracefulShutdown()
}

func (srv *server) gracefulShutdown() {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	fmt.Printf("🔄 Server is shutting down: %s\n", sig.String())

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	srv.SetKeepAlivesEnabled(false)
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("❌ Could not gracefully shutdown the server: %v\n", err)
	}
	fmt.Println("🛑 Server stopped")
}
