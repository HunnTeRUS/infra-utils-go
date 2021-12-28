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
	
	//prometheus_metrics.Handler is the middleware to register prometheus
	//metrics of this endpoint
	router.GET("/test", prometheus_metrics.Handler, handleTestEndpoint)

        //Start method returns a channel to handle and wait for errors
	channelError := serverHandler.Start(
		router,
		applicationAddress,
		logsHandler,
		applicationName,
		applicationHealthVerifier,
	)
	
	if err := channelError; err != nil {
            //DO_SOMETHING
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

## Properties and environment variables
You can change some paths and port numbers for health and prometheus endpoints

**For Prometheus services, you can set:**
- **METRICS_PATH**: Env defined to get the endpoint path for prometheus endpoint. Fallback value: "/metrics"
- **METRICS_ADDRESS**: Env defined to get the address port for prometheus endpoint. Fallback value: "8080"

**For health services, you can set:**
- **HEALTH_CHECKER_PATH**: Env defined to get the endpoint path for health checker endpoint. Fallback value: "/health"
- **HEALTH_CHECKER_ADDRESS**: Env defined to get the address port for health checker endpoint. Fallback value: "4444"

**For gracefully shutdown services, you can set:**
- **GRACEFULLY_SHUTDOWN_TIMER**: Env defined to get timeout value for endpoints when the server receives OS signal to terminate. Fallback value: "5s"

## Tips and usage
1 - The server.Start method is used to call all the services in the application, then, it is using channels to get errors between them. Thus, try to handle this
channel and wait for errors to keep your application safety.
```go
channelError := serverHandler.Start(
    router,
    applicationAddress,
    logsHandler,
    applicationName,
    applicationHealthVerifier,
)

if err := channelError; err != nil {
    //DO_SOMETHING
}
```

2 - The Logger interface is received in all methods because you can instantiate your own logger methods and pass the interface to the methods, then, the logs that
infra-utils-go generate will be handled by yourself.
```go
type logStructExample struct {
}

func NewLogInterfaceExample() logger.Logger {
    return &logStructExample{}
}

func (log *logStructExample) Info(message string) {
    fmt.Println("INFO: ", message)
}

func (log *logStructExample) Error(message string, err error) {
    fmt.Println(fmt.Sprintf("ERROR: %v, MESSAGE: %s", err, message))
}
```

3 - The http.Handler does not need to be an instance of gin-gonic, you can implement an interface that implements the methods *ServeHTTP(ResponseWriter, *Request)** method
and pass it to the server.Start method or to the gracefully_shutdown.GracefullyShutdownRun directly.
```go
type ServerInterfaceExample interface {
    ServeHTTP(http.ResponseWriter, *http.Request)
}

type serverExample struct {
}

func NewServerInterfaceExample() ServerInterfaceExample {
    return &serverExample{}
}

func (s *serverExample) ServeHTTP(http.ResponseWriter, *http.Request) {
    //IMPLEMENTS_ME
}
```