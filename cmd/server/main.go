package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Jayvir101/git-cms/internal/api"
	"github.com/Jayvir101/git-cms/internal/config"
)

func main() {
	// Load application configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	// Initialize router
	router := api.NewRouter()

	// Server address expects ":<port>"
	addr := fmt.Sprintf(":%s", cfg.Port)

	// Configure HTTP server
	server := &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("starting server on %s", addr)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server failed to start: %v", err)
	}
}
