package logger

import (
	"errors"
	"testing"
)

func TestLogger(t *testing.T) {
	SetLogLevel(2)
	LogDebug("LogDebug")
	LogError("LogError")
	LogNormal("LogNormal")
	Log("Log")

	LogDebugF(errors.New("LogDebug Error"), "LogDebug")
	LogErrorF(errors.New("LogError Error"), "LogError")
	LogNormalF(errors.New("LogNormal Error"), "LogNormal")
	LogF(errors.New("Log Error"), "Log")
}
