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
	"path/filepath"
	"syscall"

	healthQuery "github.com/go-clean/internal/probes/application/query"
	healthInfra "github.com/go-clean/internal/probes/infrastructure"
	healthHttp "github.com/go-clean/internal/probes/presentation/http"
	pingHttp "github.com/go-clean/internal/probes/presentation/http"
	swaggerQuery "github.com/go-clean/internal/swagger/application/query"
	"github.com/go-clean/internal/swagger/infrastructure"
	swaggerHttp "github.com/go-clean/internal/swagger/presentation/http"
	"github.com/go-clean/platform/config"
	"github.com/go-clean/platform/database"
	"github.com/go-clean/platform/http"
	"github.com/go-clean/platform/logger"
	platformRedis "github.com/go-clean/platform/redis"
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

	// Initialize database connection
	db, err := database.NewConnection(cfg.Database)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to database")
	}
	defer database.Close(db)
	logger.Info().Msg("Database connection established")

	// Initialize Redis connection
	redisClient, err := platformRedis.NewClient(cfg.Redis)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to Redis")
	}
	defer func() {
		if err := platformRedis.Close(redisClient); err != nil {
			logger.Error().Err(err).Msg("Failed to close Redis connection")
		}
	}()
	logger.Info().Msg("Redis connection established")

	// Initialize dependencies
	pingQueryHandler := healthQuery.NewPingQueryHandler()
	pingHandler := pingHttp.NewPingHandler(pingQueryHandler)

	// Initialize health module
	databaseChecker := healthInfra.NewDatabaseChecker(db)
	redisChecker := healthInfra.NewRedisChecker(redisClient)
	healthQueryHandler := healthQuery.NewGetHealthQueryHandler(databaseChecker, redisChecker)
	healthService := healthQuery.NewHealthService(healthQueryHandler)
	healthHandler := healthHttp.NewHealthHandler(healthService)

	// Initialize swagger module
	swaggerAdapter := infrastructure.NewSwaggerLoader(
		filepath.Join("./api", "openapi.yaml"),
		filepath.Join("./api", "swagger.html"),
	)
	if err := swaggerAdapter.Init(); err != nil {
		logger.Fatal().Err(err).Msg("Failed to initialize swagger adapter")
	}
	swaggerQueryHandler := swaggerQuery.NewSwaggerQueryHandler(swaggerAdapter)
	swaggerHandler := swaggerHttp.NewDocsHandler(swaggerQueryHandler)

	// Initialize HTTP server
	server := http.NewServer(cfg.Server.Port)
	app := server.GetApp()

	// Register routes
	pingHandler.RegisterRoutes(app)
	healthHandler.RegisterRoutes(app)
	swaggerHandler.RegisterRoutes(app)

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
