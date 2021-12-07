package prometheus_metrics

import (
	"fmt"

	"github.com/HunnTeRUS/infra-utils-go/configuration/env"
	"github.com/HunnTeRUS/infra-utils-go/configuration/log"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

var (
	METRICS_ADDRESS = "METRICS_ADDRESS"
	METRICS_PATH    = "METRICS_PATH"
)

func PrometheusMetrics(logger log.Logger) {
	metricsPath := env.Get(METRICS_PATH, "/metrics")
	metricsAdress := env.Get(METRICS_ADDRESS, "8080")

	logger.Info("About to start prometheus handler server")

	prom := promhttp.Handler()
	m := gin.Default()
	m.GET(metricsPath, func(c *gin.Context) {
		prom.ServeHTTP(c.Writer, c.Request)
	})

	if err := m.Run(fmt.Sprintf(":%s", metricsAdress)); err != nil {
		logger.Error(fmt.Sprintf("Error tring to execute promethus on %s", metricsAdress), err)
		panic(err)
	}
}
