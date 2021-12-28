//Package health implements the application health verifier to EKS knows if application is healthy
package health

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/HunnTeRUS/infra-utils-go/configuration/env"
	"github.com/HunnTeRUS/infra-utils-go/configuration/logger"
	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

//HealthChecker is a function that will return if the application is ok
type HealthChecker func() error

var (
	//HEALTH_CHECKER_ADDRESS receives the health endpoint port path to EKS to use
	HEALTH_CHECKER_ADDRESS = "HEALTH_CHECKER_ADDRESS"

	//HEALTH_CHECKER_PATH receives the health endpoint path to EKS to use
	HEALTH_CHECKER_PATH = "HEALTH_CHECKER_PATH"
)

//HealthInterface is an interface to implement the methods that health package implements
type HealthInterface interface {
	HealthCheck(
		logger logger.Logger,
		channelError chan<- error,
		checkers ...HealthChecker,
	)
}

//NewHealthHandler is used to return a instance of the required interface, so you can use
//health methods
func NewHealthHandler() HealthInterface {
	return &health{}
}

type health struct{}

//	HealthCheck is a function that will validate the healthChecker received by parameter
//	and implement the endpoint thats EKS is going to use to check if application is
//	healthy
func (ht *health) HealthCheck(
	logger logger.Logger,
	channelError chan<- error,
	checkers ...HealthChecker,
) {
	healthCheckerPath := env.Get(HEALTH_CHECKER_PATH, "/health")
	healthCheckerAdress := env.Get(HEALTH_CHECKER_ADDRESS, "4444")

	logger.Info(fmt.Sprintf("About to start health checker handler server on port %s and path %s",
		healthCheckerAdress,
		healthCheckerPath,
	))

	health := gin.Default()

	health.GET(healthCheckerPath, func(c *gin.Context) {
		healthCheckHandler(c, checkers...)
	})

	if err := health.Run(fmt.Sprintf(":%s", healthCheckerAdress)); err != nil {
		logger.Error(fmt.Sprintf("Error trying to execute healthChecker on port %s", healthCheckerAdress), err)
		channelError <- errors.New(fmt.Sprintf("Error tring to execute healthChecker on %s. Error: %v",
			healthCheckerAdress,
			err,
		))
	}
}

func healthCheckHandler(c *gin.Context, checkers ...HealthChecker) {
	c.Header("content-type", "application/json")
	for _, checker := range checkers {
		if checker == nil {
			continue
		}

		if err := checker(); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf(`{"healthy": false, "error": "%v"`, err))
			return
		}
	}

	c.String(http.StatusOK, `{"healthy": true}`)
}
