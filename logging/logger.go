package logging

import (
	"log"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.SugaredLogger
var loggerConfig zap.Config

func init() {
	loggerConfig = zap.NewProductionConfig()
	loggerConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	loggerConfig.Level = zap.NewAtomicLevel() //Info Level by default
	loggerConfig.DisableCaller = true
	buildLogger()
	logger.Debug("Logger initialized.")
}

func buildLogger() {
	_logger, err := loggerConfig.Build()
	if err != nil {
		log.Fatal(err)
	}

	defer _logger.Sync()
	logger = _logger.Sugar()
}

func SetVerbose(verbose bool) {
	if verbose {
		loggerConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		buildLogger()
	} else {
		loggerConfig.Level = zap.NewAtomicLevel()
		buildLogger()
	}
}

func Info(message string) {
	logger.Info(message)
}

func Debug(message string, fields ...zap.Field) {
	logger.Debug(message)
}

func Error(message string) {
	logger.Error(message)
}

func Fatal(message string) {
	logger.Fatal(message)
}
