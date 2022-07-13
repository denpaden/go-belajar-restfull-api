package log

import (
	"github.com/denpaden/go-belajar-restfull-api/app/helper"
	"testing"
)

func TestLogWriteToFile(t *testing.T) {
	logger := helper.NewLoggerFile()
	logger.Info("Hello loging")
	logger.Warn("Hello loging")
	logger.Error("Hello loging")
}
