//Package server implements a Start method to initialize prometheus, gracefully_shutdown and application health
package server

import (
	"net/http"

	"github.com/HunnTeRUS/infra-utils-go/configuration/logger"
	"github.com/HunnTeRUS/infra-utils-go/gracefully_shutdown"
	"github.com/HunnTeRUS/infra-utils-go/health"
	"github.com/HunnTeRUS/infra-utils-go/prometheus_metrics"
)

// ServerInterface implements the method Start that will be used to initialize infra-utils method
type ServerInterface interface {
	Start(
		handler http.Handler,
		addr string,
		logger logger.Logger,
		applicationName string,
		checkers ...health.HealthChecker,
	) <-chan error
}

type server struct {
}

// NewServerInterface returns a instance of ServerInterface
func NewServerInterface() ServerInterface {
	return &server{}
}

//Start is used to start all infra-utils-go methods and services
func (s *server) Start(
	handler http.Handler,
	addr string,
	logger logger.Logger,
	applicationName string,
	checkers ...health.HealthChecker) <-chan error {

	channelError := make(chan error, 3)
	prometheusHandler := prometheus_metrics.NewPrometheusMetricsInterface()
	healthHandler := health.NewHealthHandler()
	gracefullyHandler := gracefully_shutdown.NewGracefullyShutdownInterface()

	go prometheusHandler.PrometheusMetrics(logger, channelError, applicationName)
	go healthHandler.HealthCheck(logger, channelError, checkers...)
	go gracefullyHandler.GracefullyShutdownRun(handler, addr, logger, channelError)

	return channelError
}
