// Package mocks is used to return a instance of logs interfaces and methods
package mocks

import (
	"fmt"

	"github.com/HunnTeRUS/infra-utils-go/configuration/logger"
)

type logStruct struct {
}

// NewLogInterfaceMock is used if you want to simply instantiate a log interface and
// used in the infra-utils-go unit tests
func NewLogInterfaceMock() logger.Logger {
	return &logStruct{}
}

//Info method receive a string message and print it into the logs
func (log *logStruct) Info(message string) {
	fmt.Println("INFO: ", message)
}

//Error method receive a string message and a error to print it into the logs
func (log *logStruct) Error(message string, err error) {
	fmt.Println(fmt.Sprintf("ERROR: %v, MESSAGE: %s", err, message))
}
