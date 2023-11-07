package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/wojciechpawlinow/find-indexes/internal/config"
)

var logger *zap.SugaredLogger

// Setup initializes the logging infrastructure based on the provided configuration.
// It should be called once during the startup of the application.
func Setup(cfg config.Provider) {
	var level zapcore.Level
	switch cfg.GetString("log_level") {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}
	encoderCfg := zap.NewProductionEncoderConfig()
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.Lock(os.Stdout),
		level,
	)
	l := zap.New(core)
	logger = l.Sugar()
}

// Debug logs a debug message with the given fields.
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

// Info logs an informational message with the given fields.
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Error logs an error message with the given fields.
func Error(args ...interface{}) {
	logger.Error(args...)
}

// Fatalf logs a fatal message with the given format and arguments, then exits the application.
func Fatalf(template string, args ...interface{}) {
	logger.Fatalf(template, args...)
}
