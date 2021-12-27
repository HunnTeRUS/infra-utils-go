package env

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGet_return_fallback(t *testing.T) {

	defer os.Clearenv()
	fallbackValue := "TEST_FALLBACK"

	returnValue := Get("UNKNOW", fallbackValue)

	assert.EqualValues(t, returnValue, fallbackValue)

}

func TestGet_return_searched_value(t *testing.T) {

	defer os.Clearenv()
	value := "TEST_VALUE"

	os.Setenv("KEY_TEST", "TEST_VALUE")
	fallbackValue := "TEST_FALLBACK"

	returnValue := Get("KEY_TEST", fallbackValue)

	assert.EqualValues(t, returnValue, value)
}

func TestGetDuration_return_fallback(t *testing.T) {

	defer os.Clearenv()
	fallbackValue := 5 * time.Millisecond

	returnValue := GetDuration("UNKNOW_TIME_VALUE", fallbackValue)

	assert.EqualValues(t, returnValue, fallbackValue)

}

func TestGetDuration_return_searched_value(t *testing.T) {

	defer os.Clearenv()
	value := "300ms"

	os.Setenv("TIME_VALUE_DURATION", value)
	fallbackValue := 10 * time.Millisecond

	returnValue := GetDuration("TIME_VALUE_DURATION", fallbackValue)

	assert.EqualValues(t, returnValue, 300*time.Millisecond)
}
