package gracefully_shutdown

import (
	"fmt"
	"os"
	"syscall"
	"testing"

	"github.com/HunnTeRUS/infra-utils-go/configuration/logger/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGracefullyShutdownRun_receive_so_notify(t *testing.T) {

	router := gin.Default()
	addr := ":7878"
	log := mocks.NewLogInterfaceMock()
	gracefull := NewGracefullyShutdownInterface()

	errChannel := gracefull.GracefullyShutdownRun(router, addr, log)

	fmt.Print("test")
	syscall.Kill(os.Getpid(), syscall.SIGQUIT)

	err := <-errChannel
	assert.Nil(t, err)
}

func TestGracefullyShutdownRun_error_trying_to_start_server(t *testing.T) {
	router := gin.Default()
	addr := ":7877"
	go router.Run(addr)
	log := mocks.NewLogInterfaceMock()
	gracefull := NewGracefullyShutdownInterface()

	channelError := gracefull.GracefullyShutdownRun(router, addr, log)

	err := <-channelError
	syscall.Kill(os.Getpid(), syscall.SIGQUIT)

	assert.NotNil(t, err)
}
