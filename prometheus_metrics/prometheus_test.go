package prometheus_metrics

import (
	"os"
	"testing"

	"github.com/HunnTeRUS/infra-utils-go/configuration/logger/mocks"
	"github.com/stretchr/testify/assert"
)

var (
	log = mocks.NewLogInterfaceMock()
)

func TestPrometheusMetrics_error_trying_to_start_server(t *testing.T) {
	os.Setenv(METRICS_ADDRESS, "TEST")
	defer os.Clearenv()
	channelError := make(chan error, 1)

	prometheusHandle := NewPrometheusMetricsInterface()
	go prometheusHandle.PrometheusMetrics(log, channelError, "test")

	err := <-channelError
	assert.NotNil(t, err)
}
