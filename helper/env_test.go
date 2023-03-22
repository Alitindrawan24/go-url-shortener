package helper

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	os.Setenv("APP_ENV", "testing")
}

func TestGetEnvValueCorrect(t *testing.T) {
	key := "APP_ENV"

	value := GetEnvVariableKey(key, "")
	assert.Equal(t, os.Getenv(key), value)
}

func TestGetEnvAlternativeValueCorrect(t *testing.T) {
	key := "APP_STATUS"

	value := GetEnvVariableKey(key, "good")
	assert.Equal(t, "good", value)
}
