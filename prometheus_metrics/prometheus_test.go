package prometheus_metrics

import (
	"net/http"
	"os"
	"testing"

	"github.com/HunnTeRUS/infra-utils-go/configuration/logger/mocks"
	"github.com/stretchr/testify/assert"
)

var (
	log = mocks.NewLogInterfaceMock()
)

func TestNewPrometheusMetricsInterface(t *testing.T) {
	prometheusHandle := NewPrometheusMetricsInterface()

	assert.NotNil(t, prometheusHandle)
}

func TestPrometheusMetrics_error_trying_to_start_server(t *testing.T) {

	os.Setenv(METRICS_ADDRESS, "TEST")
	defer os.Clearenv()

	prometheusHandle := NewPrometheusMetricsInterface()
	assert.Panics(t, func() {
		prometheusHandle.PrometheusMetrics(log, "test")
	})
}

func TestPrometheusMetrics_server_started_successfully(t *testing.T) {
	os.Setenv(METRICS_ADDRESS, "7777")
	defer os.Clearenv()
	prometheusHandle := NewPrometheusMetricsInterface()
	go prometheusHandle.PrometheusMetrics(log, "test")

	resp, err := http.Get("http://localhost:7777/metrics")

	assert.Equal(t, err, nil)
	assert.NotNil(t, resp)
}
