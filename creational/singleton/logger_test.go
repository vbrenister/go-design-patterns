package singleton

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestShouldInitializeLogger(t *testing.T) {

	logger := getLoggerInstance()

	assert.NotNil(t, logger)
}