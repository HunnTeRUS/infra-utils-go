package health

import (
	"net/http"

	"github.com/HunnTeRUS/infra-utils-go/logger"
	"github.com/gin-gonic/gin"
)

type HealthChecker func() error

const (
	HealthAddr = ":4444"
	HealthPath = "/health"
)

func HealthCheck(checkers ...HealthChecker) {
	health := gin.Default()

	health.GET(HealthPath, func(c *gin.Context) {
		HealthCheckHandler(c, checkers...)
	})

	if err := health.Run(HealthAddr); err != nil {
		logger.Error("Test", err)
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
			c.String(http.StatusInternalServerError, "health")
			return
		}
	}
	c.String(http.StatusOK, "aaaaa")
}
