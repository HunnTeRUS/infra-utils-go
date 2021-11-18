package server

import (
	"net/http"

	"github.com/HunnTeRUS/infra-utils-go/gracefully_shutdown"
	"github.com/HunnTeRUS/infra-utils-go/health"
	"github.com/HunnTeRUS/infra-utils-go/logger"
	"github.com/HunnTeRUS/infra-utils-go/prometheus_metrics"
)

type Server struct {
	*http.Server
}

func Start(handler http.Handler, addr string, checkers ...health.HealthChecker) {
	go prometheus_metrics.PrometheusMetrics()
	go health.HealthCheck(checkers...)

	errC := gracefully_shutdown.GracefullyShutdownRun(handler, addr)

	if err := <-errC; err != nil {
		logger.Info("Erro")
	}

	logger.Info("Exiting")
}
