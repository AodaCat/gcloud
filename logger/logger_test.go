package logger

import (
	"testing"
	"time"
)

func TestLogger(t *testing.T) {
	initLog("logs", "test", "111", true, 1*time.Hour)
	logger.Info("test.........")
}
