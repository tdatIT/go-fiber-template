package gplog

import (
	"bytes"
	"encoding/json"
	"testing"
)

// Helper function to capture and verify log output
func verifyLogOutput(t *testing.T, buf *bytes.Buffer, expectedLevel string, expectedMsg string) {
	t.Helper() // Marks the function as a test helper

	// Print the buffer content to see the log output
	t.Log("Log output:")
	t.Log(buf.String())

	// Check if the buffer contains JSON log output
	output := buf.String()
	if output == "" {
		t.Fatal("Log output is empty")
	}

	// Decode the JSON log to verify its structure
	var logEntries []map[string]interface{}
	decoder := json.NewDecoder(buf)
	for decoder.More() {
		var entry map[string]interface{}
		if err := decoder.Decode(&entry); err != nil {
			t.Fatalf("Failed to decode log entry: %v", err)
		}
		logEntries = append(logEntries, entry)
	}

	if len(logEntries) == 0 {
		t.Fatal("No log entries found in output")
	}

	// Verify that the log entries contain expected fields
	for _, entry := range logEntries {
		if level, ok := entry["level"]; !ok || level != expectedLevel {
			t.Errorf("Log entry has incorrect or missing 'level' field: %v", entry)
		}
		if msg, ok := entry["msg"]; !ok || msg != expectedMsg {
			t.Errorf("Log entry has incorrect or missing 'msg' field: %v", entry)
		}
	}

	t.Log("Log format is correct")
}
