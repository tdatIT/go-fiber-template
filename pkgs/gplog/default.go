package gplog

import "go.uber.org/zap"

var DefaultConfig = &LogConfig{
	Level:       "debug",
	LogFormat:   JsonFormat,
	TimeFormat:  ISO8601TimeEncoder,
	ServiceName: "init-service",
}

// singleton logger
var _logger = NewLogger(DefaultConfig)

func SetLogger(l *GPLogger) {
	_logger = l
}

func GetZapInstance() *zap.Logger {
	return _logger._zaplog
}

func Info(template string, field ...zap.Field) {
	_logger._zaplog.Info(template, field...)
}

func Error(template string, field ...zap.Field) {
	_logger._zaplog.Error(template, field...)
}

func Warn(template string, field ...zap.Field) {
	_logger._zaplog.Warn(template, field...)
}

func Debug(template string, field ...zap.Field) {
	_logger._zaplog.Debug(template, field...)
}

func DPanic(template string, field ...zap.Field) {
	_logger._zaplog.DPanic(template, field...)
}

func Panic(template string, field ...zap.Field) {
	_logger._zaplog.Panic(template, field...)
}

func Fatal(template string, field ...zap.Field) {
	_logger._zaplog.Fatal(template, field...)
}
func Debugf(template string, args ...interface{}) {
	_logger.sugarLogger.Debugf(template, args...)
}

func Infof(template string, args ...interface{}) {
	_logger.sugarLogger.Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	_logger.sugarLogger.Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	_logger.sugarLogger.Errorf(template, args...)
}

func DPanicf(template string, args ...interface{}) {
	_logger.sugarLogger.DPanicf(template, args...)
}

func Panicf(template string, args ...interface{}) {
	_logger.sugarLogger.Panicf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	_logger.sugarLogger.Fatalf(template, args...)
}
