# Infra utils packages

<img align="right" width="159px" src="https://i0.wp.com/cdn-images-1.medium.com/max/1200/1*lSUb1T4YW1td0UskwsGZ1w.gif?w=1920&ssl=1">

[![Build Status](https://github.com/gin-gonic/gin/workflows/Run%20Tests/badge.svg?branch=master)](https://github.com/HunnTeRUS/infra-utils-go/actions?query=branch%3Amain)
[![codecov](https://codecov.io/gh/HunnTeRUS/infra-utils-go/branch/main/graph/badge.svg?token=5WANMQY5NA)](https://codecov.io/gh/HunnTeRUS/infra-utils-go)

Infra utils is a couple of package that implements some required features when you are using kubernetes with golang projects, just like Prometheus metrics, application health endpoint and gracefully shutdown.

## Installation

To install infra-utils-go, you need to install Go and set your Go workspace first.

1. The first need [Go](https://golang.org/) installed (**version 1.16+ is required**), then you can use the below Go command to install.

```sh
$ go get -u github.com/HunnTeRUS/infra-utils-go
```

2. Import it in your code:

```go
import "github.com/HunnTeRUS/infra-utils-go"
```

## Quick start
Implementing the infra-utils package can be with all methods together using server.Start or you can use features separately like:
- prometheus_metrics.PrometheusMetrics
- health.HealthCheck
- gracefully_shutdown.GracefullyShutdownRun

The http.Handler you can use anyone you want and logger interface you can implement another one for yourself and use
to print the logs like you want
###Using infra-utils-go with gin-gonic package
```go
package main

import (
	"net/http"

	"github.com/HunnTeRUS/infra-utils-go/configuration/logger/mocks"
	"github.com/HunnTeRUS/infra-utils-go/prometheus_metrics"
	"github.com/HunnTeRUS/infra-utils-go/server"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	serverHandler := server.NewServerInterface()
	applicationAddress := "8081"
	logsHandler := mocks.NewLogInterfaceMock()
	applicationName := "applicationExampleTest"
	//This health handler can be an ping to an database or another else
	applicationHealthVerifier := func() error {
		return nil
	}

	go serverHandler.Start(
		router,
		applicationAddress,
		logsHandler,
		applicationName,
		applicationHealthVerifier,
	)

	//prometheus_metrics.Handler is the middleware to register prometheus
	//metrics of this endpoint
	router.GET("/test", prometheus_metrics.Handler, handleTestEndpoint)

	if err := router.Run(applicationAddress); err != nil {
		panic(err)
	}
}

func handleTestEndpoint(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
```

### Running this code above, to test you can run this:
To get prometheus metrics:
```
curl --location --request GET 'http://localhost:8080/metrics'
```
To get health application status:
```
curl --location --request GET 'http://localhost:4444/health'
```

And to test gracefully shutdown, just do Ctrl + C on your terminal and see the logs!


