package logger

import (
	"log/slog"
	"testing"
)

func TestNewLogger(t *testing.T) {
	logger := NewLogger()
	if logger == nil {
		t.Fatal("Expected logger to be non-nil")
	}

	// Check if the logger is of the correct type
	_, ok := logger.Handler().(*slog.TextHandler)
	if !ok {
		t.Fatal("Expected logger to use TextHandler")
	}
}
