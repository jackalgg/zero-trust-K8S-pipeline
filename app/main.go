package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "[SENTRY-POD] ", log.LstdFlags)

	logger.Printf("Initialization started. Arch: %s, OS: %s", runtime.GOARCH, runtime.GOOS)

	uid := os.Getuid()
	gid := os.Getgid()
	logger.Printf("Process Security Context - UID: %d, GID: %d", uid, gid)

	if uid == 0 {
		logger.Println("CRITICAL: Process running with administrative privileges (root).")
	} else {
		logger.Println("STATUS: Process running with restricted privileges (non-root).")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logger.Printf("Request received: method=%s path=%s remote=%s", r.Method, r.URL.Path, r.RemoteAddr)
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "sentry-pod is running\narch=%s os=%s uid=%d gid=%d\n", runtime.GOARCH, runtime.GOOS, uid, gid)
	})

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok\n"))
	})

	server := &http.Server{
		Addr:              ":" + port,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		logger.Printf("HTTP server listening on port %s", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("HTTP server failed: %v", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigChan
	logger.Printf("Signal received: %v. Shutting down.", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Printf("Graceful shutdown failed: %v", err)
	} else {
		logger.Println("Shutdown complete.")
	}
}
