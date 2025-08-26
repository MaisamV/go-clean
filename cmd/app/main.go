package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"

	"github.com/go-clean/internal/ping/application"
	pingHttp "github.com/go-clean/internal/ping/infrastructure/http"
	"github.com/go-clean/platform/config"
	"github.com/go-clean/platform/http"
	"github.com/go-clean/platform/logger"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize logger
	logger := logger.New()
	logger.Info().Msg("Starting Go Clean Architecture application")

	// Initialize dependencies
	pingService := application.NewPingService()
	pingHandler := pingHttp.NewPingHandler(pingService)

	// Initialize HTTP server
	server := http.NewServer(cfg.Server.Port)
	app := server.GetApp()

	// Register routes
	pingHandler.RegisterRoutes(app)

	// Add a root endpoint for basic info
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"service": "go-clean-architecture",
			"version": "1.0.0",
			"status":  "running",
		})
	})

	// Start server in a goroutine
	go func() {
		logger.Info().Str("port", cfg.Server.Port).Msg("Starting HTTP server")
		if err := server.Start(); err != nil {
			logger.Fatal().Err(err).Msg("Failed to start server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info().Msg("Shutting down server...")

	// Gracefully shutdown the server
	if err := server.Shutdown(); err != nil {
		logger.Error().Err(err).Msg("Server forced to shutdown")
	}

	logger.Info().Msg("Server exited")
}