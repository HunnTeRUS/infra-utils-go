package main

import (
	"net/http"

	"github.com/HunnTeRUS/infra-utils-go/server"
	"github.com/gin-gonic/gin"
)

type HandleTest interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

func NewHandleTest() HandleTest {
	return &HandleTestStruct{}
}

type HandleTestStruct struct {
}

var (
	router = gin.Default()
)

func main() {
	test := NewHandleTest()

	server.Start(test, ":8081", testFunc)
}

func testFunc() error {
	return nil
}

func (t *HandleTestStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("test"))
	return
}

func GetAccessToken(c *gin.Context) {
