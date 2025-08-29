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
	"os"
	"os/signal"
	"syscall"

	healthQuery "github.com/go-clean/internal/probes/application/query"
	pingQuery "github.com/go-clean/internal/probes/application/query"
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
	// Initialize logger first for early error reporting
	logger := logger.New()
	logger.Info().Msg("Starting Go Clean Architecture application")

	// Load configuration
	logger.Info().Msg("Loading application configuration")
	cfg, err := config.Load(logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to load configuration")
	}
	logger.Info().Str("environment", cfg.App.Environment).Str("version", cfg.App.Version).Msg("Configuration loaded successfully")

	// Initialize database connection
	db, err := database.NewConnection(cfg.Database, logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to database")
	}
	defer database.Close(db, logger)
	logger.Info().Msg("Database connection established")

	// Initialize Redis connection
	redisClient, err := platformRedis.NewClient(cfg.Redis, logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to Redis")
	}
	defer func() {
		if err := platformRedis.Close(redisClient, logger); err != nil {
			logger.Error().Err(err).Msg("Failed to close Redis connection")
		}
	}()
	logger.Info().Msg("Redis connection established")

	// Initialize ping module
	pingQueryHandler := pingQuery.NewPingQueryHandler(logger)
	pingHandler := pingHttp.NewPingHandler(logger, pingQueryHandler)

	// Initialize health module
	databaseChecker := healthInfra.NewDatabaseChecker(logger, db)
	redisChecker := healthInfra.NewRedisChecker(logger, redisClient)
	healthQueryHandler := healthQuery.NewGetHealthQueryHandler(logger, databaseChecker, redisChecker)
	healthService := healthQuery.NewHealthService(logger, healthQueryHandler)
	healthHandler := healthHttp.NewHealthHandler(logger, healthService)

	// Initialize swagger module
	swaggerConfig := infrastructure.SwaggerConfig{
		OpenApiFilePath: "./api/openapi.yaml",
		SwaggerFilePath: "./api/swagger.html",
	}
	swaggerLoader := infrastructure.NewSwaggerLoader(logger, swaggerConfig)
	if err := swaggerLoader.Init(); err != nil {
		logger.Fatal().Err(err).Msg("Failed to initialize swagger loader")
	}
	swaggerQueryHandler := swaggerQuery.NewSwaggerQueryHandler(logger, swaggerLoader)
	docsHandler := swaggerHttp.NewDocsHandler(logger, swaggerQueryHandler)

	// Initialize HTTP server
	server := http.NewServer(cfg.Server.Port, logger)
	app := server.GetApp()

	// Register routes
	pingHandler.RegisterRoutes(app)
	healthHandler.RegisterRoutes(app)
	docsHandler.RegisterRoutes(app, cfg.Swagger.Enabled)

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
