package server

import (
	"os"
	"testing"

	"github.com/HunnTeRUS/infra-utils-go/configuration/logger/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestStart_getting_error_from_services(t *testing.T) {

	handler := gin.Default()
	os.Clearenv()
	gin.SetMode(gin.TestMode)
	logs := mocks.NewLogInterfaceMock()

	applicationName := "testeMetrics"

	serverHandler := NewServerInterface()
	channelError := serverHandler.Start(
		handler,
		"WRONG_PORT_VALUE",
		logs,
		applicationName,
		func() error {
			return nil
		},
	)

	err := <-channelError

	assert.NotNil(t, err)
}
