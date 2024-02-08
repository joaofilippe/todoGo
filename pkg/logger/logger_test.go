package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewLogger(t *testing.T){
	assert := assert.New(t)

	logger := NewLogger()
	assert.NotNil(logger, "logger coudn't be initialized")
}