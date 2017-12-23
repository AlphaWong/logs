package logs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLalamoveLoggerPassDebug(t *testing.T) {
	Logger().Debug("I am not a Debug")
	Logger().Info("I am not an Info")
	Logger().Warn("I am not a Warn")
	Logger().Error("I am not an Error")
	// It should not be called as the it will return exit code 3
	// Logger().Fatal("I am not a Fatal")
	// By default, loggers are unbuffered. However, since zap's low-level APIs allow buffering,
	// calling Sync before letting your process exit is a good habit.
	defer Logger().Sync()
	assert.True(t, true)
}
