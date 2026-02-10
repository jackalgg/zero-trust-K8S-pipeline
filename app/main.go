package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"runtime"
)

func main() {
	// Initialize standard logger
	logger := log.New(os.Stdout, "[SENTRY-POD] ", log.LstdFlags)

	logger.Printf("Initialization started. Arch: %s, OS: %s", runtime.GOARCH, runtime.GOOS)

	// Retrieve process identity
	uid := os.Getuid()
	gid := os.Getgid()

	logger.Printf("Process Security Context - UID: %d, GID: %d", uid, gid)

	// Validate non-root execution
	if uid == 0 {
		logger.Println("CRITICAL: Process running with administrative privileges (root).")
	} else {
		logger.Println("STATUS: Process running with restricted privileges (non-root).")
	}

	// Graceful shutdown handling
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	logger.Println("Service active. Awaiting signals...")

	sig := <-sigChan
	logger.Printf("Signal received: %v. Shutting down.", sig)
}