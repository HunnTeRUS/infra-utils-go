package mocks

import (
	"fmt"

	"github.com/HunnTeRUS/infra-utils-go/configuration/logger"
)

type logStruct struct {
}

func NewLogInterfaceMock() logger.Logger {
	return &logStruct{}
}

func (log *logStruct) Info(message string) {
	fmt.Println("INFO: ", message)
}

func (log *logStruct) Error(message string, err error) {
	fmt.Println(fmt.Sprintf("ERROR: %v, MESSAGE: %s", err, message))
}
