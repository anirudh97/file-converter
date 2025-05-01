package parser

import (
	"testing"

	"github.com/anirudh97/file-converter/internal/models"
)

// MockParser implements Parser interface for testing
type MockParser struct {
	name string
}

func (m *MockParser) Parse(inputFile string) (models.DataRecord, error) {
	// Simple implementation that just returns an empty record
	return models.DataRecord{}, nil
}

func TestRegisterAndGetParser(t *testing.T) {
	// Clean the parsers map before testing
	for k := range parsers {
		delete(parsers, k)
	}

	// Create mock parsers
	csvParser := &MockParser{name: "csv"}
	jsonParser := &MockParser{name: "json"}

	tests := []struct {
		name      string
		format    string
		parser    Parser
		wantErr   bool
		operation string
	}{
		{
			name:      "Register CSV parser",
			format:    "csv",
			parser:    csvParser,
			wantErr:   false,
			operation: "register",
		},
		{
			name:      "Register JSON parser",
			format:    "json",
			parser:    jsonParser,
			wantErr:   false,
			operation: "register",
		},
		{
			name:      "Register duplicate parser",
			format:    "csv",
			parser:    csvParser,
			wantErr:   true,
			operation: "register",
		},
		{
			name:      "Get CSV parser",
			format:    "csv",
			wantErr:   false,
			operation: "get",
		},
		{
			name:      "Get nonexistent parser",
			format:    "xml",
			wantErr:   true,
			operation: "get",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.operation == "register" {
				err := Register(tt.format, tt.parser)
				if (err != nil) != tt.wantErr {
					t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else if tt.operation == "get" {
				got, err := GetParser(tt.format)
				if (err != nil) != tt.wantErr {
					t.Errorf("GetParser() error = %v, wantErr %v", err, tt.wantErr)
				}
				if err == nil && got == nil {
					t.Errorf("GetParser() returned nil parser for format %s", tt.format)
				}
			}
		})
	}
}

func TestParserExists(t *testing.T) {
	// Clean the parsers map before testing
	for k := range parsers {
		delete(parsers, k)
	}

	// Register a test parser
	mockParser := &MockParser{name: "test"}
	Register("test", mockParser)

	tests := []struct {
		name   string
		format string
		want   bool
	}{
		{
			name:   "Existing parser",
			format: "test",
			want:   true,
		},
		{
			name:   "Nonexistent parser",
			format: "unknown",
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parserExists(tt.format); got != tt.want {
				t.Errorf("parserExists() = %v, want %v", got, tt.want)
			}
		})
	}
}
