package logger_test

import (
	"github.com/allenx5612/cboost/sources/logger"
	"go.uber.org/zap"
	"testing"
)

func TestInitLogger(t *testing.T) {
	logPath := "../../log/test.log"
	zapLogger := logger.InitLogger("debug")
	zapLogger.Info("log init string")
	zapLogger.Warn("test warn", zap.String("logPath", logPath))
	zapLogger.Debug("test debug", zap.String("logPath", logPath))
	zapLogger.Error("test error", zap.String("logPath", logPath))
}
