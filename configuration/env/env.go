//Package env provides different methods to get variables from os environment.
package env

import (
	"os"
	"strconv"
	"strings"
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

//GetInt is used to get env int variable and if its doesn`t exists, will return a fallback passed as parameter
func GetInt(key string, fallback int) int {
	value, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}
	i, _ := strconv.Atoi(value)
	return i
}

//GetBool is used to get env bool variable
func GetBool(key string) bool {
	value, ok := os.LookupEnv(key)

	if !ok {
		return false
	}
	b, err := strconv.ParseBool(value)
	if err != nil {
		return false
	}

	return b
}

//GetStringSlice is used to get env string slice variable
func GetStringSlice(key string) []string {
	if v := Get(key, ""); v != "" {
		s := strings.ReplaceAll(v, " ", "")
		return strings.Split(s, ",")
	}

	return []string{}
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
