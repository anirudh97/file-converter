package writer

import (
	"testing"

	"github.com/anirudh97/file-converter/internal/models"
)

type MockWriter struct{}

func (m *MockWriter) Write(data models.DataRecord, outputFile string) error {
	return nil
}

func TestRegisterAndGetWriter(t *testing.T) {
	for k := range writers {
		delete(writers, k)
	}

	jsonWriter := &MockWriter{}
	csvWriter := &MockWriter{}

	tests := []struct {
		name      string
		format    string
		writer    Writer
		wantErr   bool
		operation string
	}{
		{
			name:      "Register JSON writer",
			format:    "json",
			writer:    jsonWriter,
			wantErr:   false,
			operation: "register",
		},
		{
			name:      "Register CSV writer",
			format:    "csv",
			writer:    csvWriter,
			wantErr:   false,
			operation: "register",
		},
		{
			name:      "Register duplicate writer",
			format:    "json",
			writer:    jsonWriter,
			wantErr:   true,
			operation: "register",
		},
		{
			name:      "Get JSON writer",
			format:    "json",
			wantErr:   false,
			operation: "get",
		},
		{
			name:      "Get nonexistent writer",
			format:    "yaml",
			wantErr:   true,
			operation: "get",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.operation == "register" {
				err := Register(tt.format, tt.writer)
				if (err != nil) != tt.wantErr {
					t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else if tt.operation == "get" {
				got, err := GetWriter(tt.format)
				if (err != nil) != tt.wantErr {
					t.Errorf("GetWriter() error = %v, wantErr %v", err, tt.wantErr)
				}
				if err == nil && got == nil {
					t.Errorf("GetWriter() returned nil writer for format %s", tt.format)
				}
			}
		})
	}
}

func TestWriterExists(t *testing.T) {
	for k := range writers {
		delete(writers, k)
	}

	mockWriter := &MockWriter{}
	Register("test", mockWriter)

	tests := []struct {
		name   string
		format string
		want   bool
	}{
		{
			name:   "Existing writer",
			format: "test",
			want:   true,
		},
		{
			name:   "Nonexistent writer",
			format: "unknown",
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := writerExists(tt.format); got != tt.want {
				t.Errorf("writerExists() = %v, want %v", got, tt.want)
			}
		})
	}
}
