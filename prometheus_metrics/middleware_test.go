package prometheus_metrics

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestInitializeConfigurations(t *testing.T) {
	applicationName := "test"
	initializeConfigurations(applicationName)

	assert.NotEmpty(t, dimensions)
	assert.NotNil(t, reqTotal)
	assert.NotNil(t, reqDuration)
}

func TestHandler_test_metrics_successfully_saved(t *testing.T) {
	go NewPrometheusMetricsInterface().PrometheusMetrics(log, "metricsTest")
	router := gin.Default()

	router.GET("/testMetrics", Handler, handleEmptyPath)
	go router.Run(":9898")

	resp, err := http.Get("http://localhost:9898/testMetrics")
	if err != nil || resp.StatusCode != http.StatusOK {
		t.Fatal()
	}

	respMetrics, err := http.Get("http://localhost:8080/metrics")
	if err != nil || respMetrics.StatusCode != http.StatusOK {
		t.Fatal()
	}

	var metricsResponse string
	closer, _ := io.ReadAll(respMetrics.Body)

	// Getting a slice of bytes because string is so long, and then, on convert to a string variable
	// the last informations is hidden
	for i := 7500; i < len(closer); i++ {
		metricsResponse = metricsResponse + string(closer[i])
	}

	assert.True(t, strings.Contains(metricsResponse, `test_request_total{method="GET",path="/test",status="200"} 1`))
}

func handleEmptyPath(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}
