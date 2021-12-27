package gracefully_shutdown

import (
	"os"
	"syscall"
	"testing"
	"time"

	"github.com/HunnTeRUS/infra-utils-go/configuration/logger/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGracefullyShutdownRun_receive_so_notify(t *testing.T) {

	router := gin.Default()
	addr := ":7878"
	log := mocks.NewLogInterfaceMock()
	gracefull := NewGracefullyShutdownInterface()
	channelError := make(chan error)

	gracefull.GracefullyShutdownRun(router, addr, log, channelError)
	time.Sleep(200 * time.Millisecond)

	syscall.Kill(os.Getpid(), syscall.SIGQUIT)
	err := <-channelError

	assert.Nil(t, err)
}

func TestGracefullyShutdownRun_error_trying_to_start_server(t *testing.T) {

	router := gin.Default()
	addr := "TEST_TEST"
	go router.Run(addr)
	log := mocks.NewLogInterfaceMock()
	gracefull := NewGracefullyShutdownInterface()
	channelError := make(chan error)

	go gracefull.GracefullyShutdownRun(router, addr, log, channelError)

	err := <-channelError

	assert.NotNil(t, err)
}
