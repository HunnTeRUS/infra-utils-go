package health

import (
	"fmt"
	"net/http"

	"github.com/HunnTeRUS/infra-utils-go/configuration/env"
	"github.com/HunnTeRUS/infra-utils-go/configuration/log"
	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

type HealthChecker func() error

var (
	HEALTH_CHECKER_ADDRESS = "HEALTH_CHECKER_ADDRESS"
	HEALTH_CHECKER_PATH    = "HEALTH_CHECKER_PATH"
)

func HealthCheck(logger log.Logger, checkers ...HealthChecker) {
	healthCheckerPath := env.Get(HEALTH_CHECKER_PATH, "/health")
	healthCheckerAdress := env.Get(HEALTH_CHECKER_ADDRESS, "4444")

	health := gin.Default()

	health.GET(healthCheckerPath, func(c *gin.Context) {
		HealthCheckHandler(c, checkers...)
	})

	if err := health.Run(fmt.Sprintf(":%s", healthCheckerAdress)); err != nil {
		logger.Error(fmt.Sprintf("Error trying to execute healthChecker on port %s", healthCheckerAdress), err)
		panic(err)
	}
}

func HealthCheckHandler(c *gin.Context, checkers ...HealthChecker) {
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
