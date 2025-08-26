// @title Go Clean Architecture API
// @version 1.0
// @description A clean architecture implementation in Go with comprehensive API documentation
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email support@example.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /
// @schemes http https
package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"

	_ "github.com/go-clean/docs" // Import generated docs
	"github.com/go-clean/internal/ping/application"
	pingHttp "github.com/go-clean/internal/ping/infrastructure/http"
	docsApp "github.com/go-clean/internal/docs/application"
	"github.com/go-clean/internal/docs/infrastructure"
	docsHttp "github.com/go-clean/internal/docs/infrastructure/http"
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

	// Initialize docs module
	docsAdapter := infrastructure.NewDocsAdapter("./api")
	docsService := docsApp.NewDocsService(docsAdapter)
	docsHandler := docsHttp.NewDocsHandler(docsService)

	// Initialize HTTP server
	server := http.NewServer(cfg.Server.Port)
	app := server.GetApp()

	// Register routes
	pingHandler.RegisterRoutes(app)
	docsHandler.RegisterRoutes(app)
	
	// Register health endpoint
	app.Get("/health", pingHttp.Health)
	
	// Register Swagger UI route
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

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