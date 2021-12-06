package prometheus_metrics

import (
	"github.com/HunnTeRUS/infra-utils-go/configuration/log"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	MetricsAddr = ":3333"
	MetricsPath = "/metrics"
)

func PrometheusMetrics(logger log.Logger) {
	logger.Info("Abount to start prometheus handler server")

	prom := promhttp.Handler()
	m := gin.Default()
	m.GET(MetricsPath, func(c *gin.Context) {
		prom.ServeHTTP(c.Writer, c.Request)
	})

	if err := m.Run(MetricsAddr); err != nil {
		logger.Error("ERrro", err)
		panic(err)
	}
}
