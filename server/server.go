//Package server implements a Start method to initialize prometheus, gracefully_shutdown and application health
package server

import (
	"log"
	"net/http"

	"github.com/HunnTeRUS/infra-utils-go/configuration/logger"
	"github.com/HunnTeRUS/infra-utils-go/gracefully_shutdown"
	"github.com/HunnTeRUS/infra-utils-go/health"
	"github.com/HunnTeRUS/infra-utils-go/prometheus_metrics"
)

type ServerInterface interface {
	Start(handler http.Handler, addr string, logger log.Logger, checkers ...health.HealthChecker)
}

type Server struct {
}

//Start is used to start all infra-utils-go methods and services
func (s *Server) Start(
	handler http.Handler,
	addr string,
	logger logger.Logger,
	errChannel chan<- error,
	checkers ...health.HealthChecker) {

	prometheusHandler := prometheus_metrics.NewPrometheusMetricsInterface()
	healthHandler := health.NewHealthHandler()
	gracefullyHandler := gracefully_shutdown.NewGracefullyShutdownInterface()

	go prometheusHandler.PrometheusMetrics(logger)
	go healthHandler.HealthCheck(logger, checkers...)
	errC := gracefullyHandler.GracefullyShutdownRun(handler, addr, logger)

	if err := <-errC; err != nil {
		logger.Error("Error tryng to shutdown server", err)
		errChannel <- err
		return
	}

	logger.Info("Exiting...")
}
