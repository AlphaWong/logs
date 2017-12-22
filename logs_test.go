package logs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLalamoveLoggerPassInfo(t *testing.T) {
	Logger().Info("I am not an error")
	// By default, loggers are unbuffered. However, since zap's low-level APIs allow buffering,
	// calling Sync before letting your process exit is a good habit.
	defer Logger().Sync()
	assert.True(t, true)
}

func TestGetLalamoveLoggerPassError(t *testing.T) {
	Logger().Error("I am an error")
	// By default, loggers are unbuffered. However, since zap's low-level APIs allow buffering,
	// calling Sync before letting your process exit is a good habit.
	defer Logger().Sync()
	assert.True(t, true)
}
