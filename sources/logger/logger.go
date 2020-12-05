package logger

import (
	"fmt"
	"github.com/allenx5612/cboost/sources/execpath"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"time"
)

// InitLogger init logger config, logLevel to control what level in log will
// record, default is info, levels(debug, info, warn, error). Debug level will
// turn on the standard output.
func InitLogger(logLevel string) (logger *zap.Logger) {
	execPath := execpath.ExecPath()
	fileName := fmt.Sprintf("cboost-%s.log", time.Now().Format("2006-01-02"))
	logPath := filepath.Join(execPath, "./log", fileName)
	name := "cboost"
	hook := lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    10,
		MaxAge:     7,
		MaxBackups: 30,
		Compress:   false,
	}

	encoderConfig := zapcore.EncoderConfig{
		MessageKey:    "msg",
		LevelKey:      "level",
		TimeKey:       "time",
		NameKey:       "logger",
		CallerKey:     "file",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder,
		EncodeTime:    zapcore.ISO8601TimeEncoder,
		EncodeCaller:  zapcore.ShortCallerEncoder,
		EncodeName:    zapcore.FullNameEncoder,
	}
	atomicLevel := zap.NewAtomicLevel()
	var writes = []zapcore.WriteSyncer{zapcore.AddSync(&hook)}
	switch logLevel {
	case "debug":
		atomicLevel.SetLevel(zap.DebugLevel)
		writes = append(writes, zapcore.AddSync(os.Stdout))
	case "info":
		atomicLevel.SetLevel(zap.InfoLevel)
	case "warn":
		atomicLevel.SetLevel(zap.WarnLevel)
	case "error":
		atomicLevel.SetLevel(zap.ErrorLevel)
	default:
		atomicLevel.SetLevel(zap.InfoLevel)
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(writes...),
		atomicLevel,
	)
	caller := zap.AddCaller()
	development := zap.Development()
	field := zap.Fields(zap.String("appname", name))

	zapLogger := zap.New(core, caller, development, field)
	return zapLogger
}

func SetLogger() (logger *zap.Logger) {
	logger = InitLogger("warn")
	return logger
}
