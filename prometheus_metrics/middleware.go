//Package prometheus_metrics implements the service and the configuration to run prometheus server
package prometheus_metrics

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	dimensions  *[]string = &[]string{""}
	reqTotal    *prometheus.CounterVec
	reqDuration *prometheus.HistogramVec
)

//Handler implements the endpoint middleware to intercept all requests
//and save the data
func Handler(c *gin.Context) {
	start := time.Now()

	c.Next()
	status := strconv.Itoa(c.Writer.Status())
	elapsed := float64(time.Since(start)) / float64(time.Millisecond)
	path := c.FullPath()
	if path == "" {
		path = "NOT_FOUND"
	}
	reqTotal.WithLabelValues(status, c.Request.Method, path).Inc()
	reqDuration.WithLabelValues(status, c.Request.Method, path).Observe(elapsed)
}

func initializeConfigurations(applicationName string) {
	*dimensions = []string{"status", "method", "path"}
	reqTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name:      "request_total",
		Help:      "Total number of requests handled",
		Subsystem: applicationName,
	}, *dimensions)
	reqDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:      "request_duration_milliseconds",
		Help:      "Request latencies in milliseconds",
		Subsystem: applicationName,
	}, *dimensions)
}
