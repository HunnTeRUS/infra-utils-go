//Package env provides different methods to get variables from os environment.
package env

import (
	"os"
	"time"
)

const (
	//LogLevel refers to log level env var, that means to get level of the application logs
	LogLevel = "LOG_LEVEL"
	//LogFormat refers to log format env var, that means to get format of the application logs
	LogFormat = "LOG_FORMAT"
)

//Get is used to get env string variable and if its doesn`t exists, will return a fallback passed as parameter
func Get(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

//GetDuration is used to get env time duration variable and if its doesn`t exists, will return a fallback passed as parameter
func GetDuration(key string, fallback time.Duration) time.Duration {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	i, _ := time.ParseDuration(value)
	return i
}
