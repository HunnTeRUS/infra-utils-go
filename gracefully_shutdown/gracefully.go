//Package gracefully_shutdown implements the application shutdown
package gracefully_shutdown

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/HunnTeRUS/infra-utils-go/configuration/env"
	"github.com/HunnTeRUS/infra-utils-go/configuration/logger"
)

const (
	//GRACEFULLY_SHUTDOWN_TIMER receives a timeout to wait for current requests get finished
	GRACEFULLY_SHUTDOWN_TIMER = "GRACEFULLY_SHUTDOWN_TIMER"
	//FALLBACK_GRACEFULLY_SHUTDOWN sets a timeout fallback value to set if GRACEFULLY_SHUTDOWN_TIMER is not set
	FALLBACK_GRACEFULLY_SHUTDOWN = 5 * time.Second
)

//GracefullyShutdownInterface declare gracefully_shutdown functions
type GracefullyShutdownInterface interface {
	GracefullyShutdownRun(
		handler http.Handler,
		addr string,
		logger logger.Logger,
		chanError chan<- error,
	)
}

//NewGracefullyShutdownInterface returns a instance of gracefully_shutdown interface,
//so you can call GracefullyShutdownRun with this instance
func NewGracefullyShutdownInterface() GracefullyShutdownInterface {
	return &gracefully{}
}

type gracefully struct{}

//GracefullyShutdownRun implements the gracefully shutdown to the application, waits
//for a SO signal to shutdown the application
func (gcf *gracefully) GracefullyShutdownRun(
	handler http.Handler,
	addr string,
	logger logger.Logger,
	chanError chan<- error) {

	logger.Info(fmt.Sprintf("About to start application on port %s",
		addr,
	))

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", addr),
		Handler: handler,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-ctx.Done()

		logger.Info("Shutdown signal received")

		shutdownTimeout := env.GetDuration(GRACEFULLY_SHUTDOWN_TIMER, FALLBACK_GRACEFULLY_SHUTDOWN)
		ctxTimeout, cancel := context.WithTimeout(context.Background(), shutdownTimeout)

		defer func() {
			stop()
			cancel()
			close(chanError)
		}()

		srv.Shutdown(ctxTimeout)

		logger.Info("Shutdown completed")
	}()

	go func() {
		logger.Info("Listening and serving")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			chanError <- errors.New(fmt.Sprintf("Error tring to run application on %s. Error: %v",
				addr,
				err,
			))
			return
		}
	}()
}
