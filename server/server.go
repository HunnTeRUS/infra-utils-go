package server

import (
	"net/http"

	"github.com/HunnTeRUS/infra-utils-go/configuration/log"
	"github.com/HunnTeRUS/infra-utils-go/gracefully_shutdown"
	"github.com/HunnTeRUS/infra-utils-go/health"
	"github.com/HunnTeRUS/infra-utils-go/prometheus_metrics"
)

type Server struct {
	*http.Server
}

func Start(handler http.Handler, addr string, logger log.Logger, checkers ...health.HealthChecker) {
	go prometheus_metrics.PrometheusMetrics(logger)
	go health.HealthCheck(logger, checkers...)

	errC := gracefully_shutdown.GracefullyShutdownRun(handler, addr, logger)

	if err := <-errC; err != nil {
		logger.Error("Error tryng to shutdown server", err)
	}

	logger.Info("Exiting...")
}
