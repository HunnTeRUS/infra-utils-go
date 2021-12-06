package env

import (
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	LogLevel  = "LOG_LEVEL"
	LogFormat = "LOG_FORMAT"
)

func Get(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func GetInt(key string, fallback int) int {
	value, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}
	i, _ := strconv.Atoi(value)
	return i
}

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

func GetStringSlice(key string) []string {
	if v := Get(key, ""); v != "" {
		s := strings.ReplaceAll(v, " ", "")
		return strings.Split(s, ",")
	}

	return []string{}
}

func GetDuration(key string, fallback time.Duration) time.Duration {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	i, _ := time.ParseDuration(value)
	return i
}
