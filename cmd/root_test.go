package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/gkwa/ourport/internal/logger"
	"github.com/go-logr/logr"
)

func TestCustomLogger(t *testing.T) {
	var buf bytes.Buffer
	testLogger := logger.NewConsoleLoggerWithWriter(&buf, true, false)

	// Test the logger
	testLogger.Info("test message", "key", "value")

	output := buf.String()
	if len(output) == 0 {
		t.Errorf("Expected log output, but got none")
	}

	t.Logf("Log output: %s", output)

	if !strings.Contains(output, "test message") {
		t.Errorf("Expected log output to contain 'test message', but got: %s", output)
	}
}

func TestJSONLogger(t *testing.T) {
	var buf bytes.Buffer
	testLogger := logger.NewConsoleLoggerWithWriter(&buf, true, true)

	// Test the logger
	testLogger.Info("test message", "key", "value")

	output := strings.TrimSpace(buf.String())
	if len(output) == 0 {
		t.Errorf("Expected log output, but got none")
	}

	// Check if it's valid JSON
	var jsonData map[string]interface{}
	if err := json.Unmarshal([]byte(output), &jsonData); err != nil {
		t.Errorf("Expected valid JSON, but got error: %v", err)
	}

	t.Logf("Log output: %s", output)

	if !strings.Contains(output, "test message") {
		t.Errorf("Expected log output to contain 'test message', but got: %s", output)
	}
}

func TestLoggerFromContext(t *testing.T) {
	var buf bytes.Buffer
	testLogger := logger.NewConsoleLoggerWithWriter(&buf, true, false)

	ctx := context.WithValue(context.Background(), loggerKey, testLogger)

	retrievedLogger := LoggerFrom(ctx)
	if retrievedLogger == (logr.Logger{}) {
		t.Error("Expected to retrieve logger from context, but got empty logger")
	}

	// Test that the retrieved logger works
	retrievedLogger.Info("context test")

	output := buf.String()
	if !strings.Contains(output, "context test") {
		t.Errorf("Expected log output to contain 'context test', but got: %s", output)
	}
}

func TestLoggerFromContextEmpty(t *testing.T) {
	ctx := context.Background()

	retrievedLogger := LoggerFrom(ctx)

	// Should return a discard logger, not panic
	retrievedLogger.Info("this should not appear anywhere")

	// If we get here without panicking, the test passes
}
