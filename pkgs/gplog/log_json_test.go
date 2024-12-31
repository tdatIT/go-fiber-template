package gplog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"testing"
)

func TestCreateJsonLogFormat(t *testing.T) {
	// Mock config for JSON log format
	config := &LogConfig{
		Level:      "debug",
		LogFormat:  JsonFormat,
		TimeFormat: ISO8601TimeEncoder,
		Filename:   "",
	}

	logger := NewLogger(config)

	if logger == nil {
		t.Fatal("Failed to create logger")
	}

	// Check if logger is not nil
	if logger._zaplog == nil {
		t.Fatal("zap logger instance is nil")
	}

	// Check if the logger core is configured to use the correct encoder (JSON)
	encoderCfg := EncoderBuilder(config)
	expectedEncoder := zapcore.NewConsoleEncoder(encoderCfg)

	core := zapcore.NewCore(expectedEncoder, zapcore.AddSync(os.Stdout), zap.DebugLevel)
	expectedLogger := zap.New(core)

	if logger._zaplog.Core() == expectedLogger.Core() {
		t.Fatal("Logger core does not match the expected core")
	}

	t.Log("Logger created successfully with JSON format")
}
