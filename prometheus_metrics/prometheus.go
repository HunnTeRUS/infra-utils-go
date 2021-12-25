//Package prometheus_metrics implements the service and the configuration to run prometheus server
package prometheus_metrics

import (
	"fmt"

	"github.com/HunnTeRUS/infra-utils-go/configuration/env"
	"github.com/HunnTeRUS/infra-utils-go/configuration/logger"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

var (
	//METRICS_ADDRESS gets the endpoint port from env to implements an endpoint inside the
	//application that will return all prometheus data saved
	METRICS_ADDRESS = "METRICS_ADDRESS"
	//METRICS_PATH gets the endpoint path from env to implements an endpoint inside the
	//application that will return all prometheus data saved
	METRICS_PATH = "METRICS_PATH"
)

//PrometheusMetricsInterface declares prometheus metrics functions
type PrometheusMetricsInterface interface {
	PrometheusMetrics(logger logger.Logger)
}

type prometheusService struct{}

//NewPrometheusMetricsInterface returns a instance of NewPrometheusMetricsInterface
//so you can call prometheus functions
func NewPrometheusMetricsInterface() PrometheusMetricsInterface {
	return &prometheusService{}
}

//PrometheusMetrics implements the server and endpoint to return all prometheus data
//saved about the project
func (prm *prometheusService) PrometheusMetrics(logger logger.Logger) {
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
