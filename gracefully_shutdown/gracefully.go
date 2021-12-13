//Package gracefully_shutdown implements the application shutdown
package gracefully_shutdown

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/HunnTeRUS/infra-utils-go/configuration/env"
	"github.com/HunnTeRUS/infra-utils-go/configuration/log"
)

const (
	//GRACEFULLY_SHUTDOWN_TIMER receives a timeout to wait for current requests get finished
	GRACEFULLY_SHUTDOWN_TIMER = "GRACEFULLY_SHUTDOWN_TIMER"
	//FALLBACK_GRACEFULLY_SHUTDOWN sets a timeout fallback value to set if GRACEFULLY_SHUTDOWN_TIMER is not set
	FALLBACK_GRACEFULLY_SHUTDOWN = 5 * time.Second
)

//GracefullyShutdownRun implements the gracefully shutdown to the application, waits
//for a SO signal to shutdown the application
func GracefullyShutdownRun(handler http.Handler, addr string, logger log.Logger) <-chan error {
	errC := make(chan error, 1)

	srv := &http.Server{
		Addr:    addr,
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
			close(errC)
		}()

		if err := srv.Shutdown(ctxTimeout); err != nil {
			errC <- err
		}

		logger.Info("Shutdown completed")
	}()

	go func() {
		logger.Info("Listening and serving")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errC <- err
		}
	}()

	return errC
}
