package health

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/HunnTeRUS/infra-utils-go/configuration/logger/mocks"
	"github.com/stretchr/testify/assert"
)

type HealthResponseStruct struct {
	Healthy bool  `json:"healthy"`
	Err     error `json:"error"`
}

func TestHealthCheck_error_trying_to_start_health_endpoint(t *testing.T) {
	portWrongTestValue := "TEST"
	os.Setenv(HEALTH_CHECKER_ADDRESS, portWrongTestValue)
	defer os.Clearenv()
	channelError := make(chan error, 1)

	heandleHealth := NewHealthHandler()

	healthChecker := func() error {
		return nil
	}

	heandleHealth.HealthCheck(mocks.NewLogInterfaceMock(), channelError, healthChecker)

	err := channelError

	assert.NotNil(t, err)
}

func TestHealthCheck_start_successfully_and_is_application_healthy(t *testing.T) {
	portWrongTestValue := "4545"
	os.Setenv(HEALTH_CHECKER_ADDRESS, portWrongTestValue)
	heandleHealth := NewHealthHandler()
	channelError := make(chan error, 1)
	healthChecker := func() error {
		return nil
	}

	go heandleHealth.HealthCheck(mocks.NewLogInterfaceMock(), channelError, healthChecker)
	w := httptest.NewRecorder()
	healthResponse := HealthResponseStruct{}

	resp, err := http.Get("http://localhost:4545/health")

	if err != nil || w.Code != http.StatusOK {
		t.Fail()
	}

	closer, _ := io.ReadAll(resp.Body)

	err = json.Unmarshal(closer, &healthResponse)
	assert.True(t, healthResponse.Healthy)
}

func TestHealthCheck_checker_is_nil(t *testing.T) {
	portWrongTestValue := "4546"
	os.Setenv(HEALTH_CHECKER_ADDRESS, portWrongTestValue)
	heandleHealth := NewHealthHandler()
	channelError := make(chan error, 1)
	w := httptest.NewRecorder()
	go heandleHealth.HealthCheck(mocks.NewLogInterfaceMock(), channelError, nil)
	healthResponse := HealthResponseStruct{}

	resp, err := http.Get("http://localhost:4546/health")

	if err != nil || w.Code != http.StatusOK {
		t.Fail()
	}

	closer, _ := io.ReadAll(resp.Body)

	err = json.Unmarshal(closer, &healthResponse)

	assert.True(t, healthResponse.Healthy)
}

func TestHealthCheck_application_is_not_healthy(t *testing.T) {
	portWrongTestValue := "4547"
	os.Setenv(HEALTH_CHECKER_ADDRESS, portWrongTestValue)
	heandleHealth := NewHealthHandler()
	channelError := make(chan error, 1)
	healthChecker := func() error {
		return errors.New("Error test")
	}

	go heandleHealth.HealthCheck(mocks.NewLogInterfaceMock(), channelError, healthChecker)
	w := httptest.NewRecorder()
	healthResponse := HealthResponseStruct{}

	resp, err := http.Get("http://localhost:4547/health")

	if err != nil || w.Code != http.StatusOK {
		t.Fail()
	}

	closer, _ := io.ReadAll(resp.Body)

	err = json.Unmarshal(closer, &healthResponse)
	assert.False(t, healthResponse.Healthy)
}
