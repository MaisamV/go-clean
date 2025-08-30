//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-clean/internal/probes"
	probesHttp "github.com/go-clean/internal/probes/presentation/http"
	"github.com/go-clean/internal/swagger"
	swaggerHttp "github.com/go-clean/internal/swagger/presentation/http"
	"github.com/go-clean/platform"
	"github.com/go-clean/platform/config"
	"github.com/go-clean/platform/http"
	"github.com/go-clean/platform/logger"
	"github.com/google/wire"
)

// Application holds all the application dependencies
type Application struct {
	Config     *config.Config
	Logger     logger.Logger
	HTTPServer *http.Server
	Probes     *ProbesModule
	Swagger    *SwaggerModule
}

// ProbesModule holds all probes-related dependencies
type ProbesModule struct {
	PingHandler   *probesHttp.PingHandler
	HealthHandler *probesHttp.HealthHandler
}

// SwaggerModule holds all swagger-related dependencies
type SwaggerModule struct {
	DocsHandler *swaggerHttp.DocsHandler
}

// InitializeApplication creates and initializes the application with all dependencies
func InitializeApplication() (*Application, error) {
	wire.Build(
		// Platform providers
		platform.PlatformSet,

		// Internal module providers
		probes.ProbesSet,
		swagger.SwaggerSet,

		// Application structure providers
		ProvideProbesModule,
		ProvideSwaggerModule,
		ProvideApplication,
	)
	return &Application{}, nil
}

// ProvideProbesModule provides the probes module
func ProvideProbesModule(
	pingHandler *probesHttp.PingHandler,
	healthHandler *probesHttp.HealthHandler,
) *ProbesModule {
	return &ProbesModule{
		PingHandler:   pingHandler,
		HealthHandler: healthHandler,
	}
}

// ProvideSwaggerModule provides the swagger module
func ProvideSwaggerModule(
	docsHandler *swaggerHttp.DocsHandler,
) *SwaggerModule {
	return &SwaggerModule{
		DocsHandler: docsHandler,
	}
}

// ProvideApplication provides the main application structure
func ProvideApplication(
	config *config.Config,
	logger logger.Logger,
	httpServer *http.Server,
	probesModule *ProbesModule,
	swaggerModule *SwaggerModule,
) *Application {
	return &Application{
		Config:     config,
		Logger:     logger,
		HTTPServer: httpServer,
		Probes:     probesModule,
		Swagger:    swaggerModule,
	}
}
