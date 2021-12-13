//Package server implements a Start method to initialize prometheus, gracefully_shutdown and application health
package server

import (
	"net/http"

	"github.com/HunnTeRUS/infra-utils-go/configuration/log"
	"github.com/HunnTeRUS/infra-utils-go/gracefully_shutdown"
	"github.com/HunnTeRUS/infra-utils-go/health"
	"github.com/HunnTeRUS/infra-utils-go/prometheus_metrics"
)

//Server is a struct to receive a http server from the application origin to implements
//the gracefully_shutdown inside it
type Server struct {
	*http.Server
}

//Start is used to start all infra-utils-go methods and services
func Start(handler http.Handler, addr string, logger log.Logger, checkers ...health.HealthChecker) {
	go prometheus_metrics.PrometheusMetrics(logger)
	go health.HealthCheck(logger, checkers...)
	errC := gracefully_shutdown.GracefullyShutdownRun(handler, addr, logger)

	if err := <-errC; err != nil {
		logger.Error("Error tryng to shutdown server", err)
	}

	logger.Info("Exiting...")
}
