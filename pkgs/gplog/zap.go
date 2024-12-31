package gplog

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type GPLogger struct {
	_zaplog     *zap.Logger
	sugarLogger *zap.SugaredLogger
}

func (l GPLogger) GetZapInstance() *zap.Logger {
	return l._zaplog
}

// NewLogger creates a new instance of Logger
func NewLogger(cfg *LogConfig) *GPLogger {
	var logInstance GPLogger

	encoderCfg := EncoderBuilder(cfg)
	var encoder zapcore.Encoder
	if cfg.LogFormat == JsonFormat {
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	}

	logLevel := zap.NewAtomicLevel()

	switch cfg.Level {
	case "debug":
		logLevel.SetLevel(zap.DebugLevel)
	case "info":
		logLevel.SetLevel(zap.InfoLevel)
	case "warn":
		logLevel.SetLevel(zap.WarnLevel)
	case "error":
		logLevel.SetLevel(zap.ErrorLevel)
	case "dpanic":
		logLevel.SetLevel(zap.DPanicLevel)
	case "panic":
		logLevel.SetLevel(zap.PanicLevel)
	case "fatal":
		logLevel.SetLevel(zap.FatalLevel)
	default:
		logLevel.SetLevel(zap.InfoLevel)
	}

	var logWriter zapcore.WriteSyncer
	if cfg.Filename != "" {
		file, err := os.Create(cfg.Filename)
		if err != nil {
			fmt.Errorf("failed to create log file: %v", err)
		}
		logWriter = zapcore.AddSync(file)
	} else {
		logWriter = zapcore.AddSync(os.Stdout)
	}

	core := zapcore.NewCore(encoder, logWriter, logLevel)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	logInstance._zaplog = logger.With(zap.String("service_name", cfg.ServiceName))
	logInstance.sugarLogger = logger.Sugar().With(zap.String("service_name", cfg.ServiceName))

	return &logInstance
}

func EncoderBuilder(cfg *LogConfig) zapcore.EncoderConfig {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    "func",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	switch cfg.TimeFormat {
	case RFC3339TimeEncoder:
		encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	case RFC3339NanoTimeEncoder:
		encoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder
	default:
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	}

	return encoderConfig
}
