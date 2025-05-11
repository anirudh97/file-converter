package converter

import (
	"errors"
	"testing"

	"github.com/anirudh97/file-converter/internal/models"
	"github.com/anirudh97/file-converter/internal/parser"
	"github.com/anirudh97/file-converter/internal/writer"
)

// MockParser ignores the actual file and returns predefined data
type MockParser struct{}

func (m *MockParser) Parse(inputFile string) (models.DataRecord, error) {
	// Instead of trying to open the file, just check the filename
	if inputFile == "error.csv" {
		return nil, errors.New("mock parsing error")
	}
	// Return mock data
	return models.DataRecord{
		{"name": "John", "age": "30"},
		{"name": "Alice", "age": "25"},
	}, nil
}

// MockWriter doesn't actually write to a file
type MockWriter struct{}

func (m *MockWriter) Write(data models.DataRecord, outputFile string) error {
	if outputFile == "error.json" {
		return errors.New("mock writing error")
	}
	return nil
}

func TestConverter_Convert(t *testing.T) {
	// Clear existing registrations to avoid interference
	parser.ClearRegistry()
	writer.ClearRegistry()

	// Register our mock parser and writer
	parser.Register("csv", &MockParser{})
	parser.Register("test", &MockParser{})
	writer.Register("json", &MockWriter{})
	writer.Register("test", &MockWriter{})

	tests := []struct {
		name       string
		converter  *Converter
		wantErr    bool
		errMessage string
	}{
		{
			name: "Successful conversion",
			converter: &Converter{
				InputFile:  "test.csv",
				InputExt:   "csv",
				OutputFile: "test.json",
				OutputExt:  "json",
			},
			wantErr: false,
		},
		{
			name: "Parser not found",
			converter: &Converter{
				InputFile:  "test.unknown",
				InputExt:   "unknown",
				OutputFile: "test.json",
				OutputExt:  "json",
			},
			wantErr:    true,
			errMessage: "error in getting parser",
		},
		{
			name: "Writer not found",
			converter: &Converter{
				InputFile:  "test.csv",
				InputExt:   "csv",
				OutputFile: "test.unknown",
				OutputExt:  "unknown",
			},
			wantErr:    true,
			errMessage: "error in getting writer",
		},
		{
			name: "Parsing error",
			converter: &Converter{
				InputFile:  "error.csv",
				InputExt:   "csv",
				OutputFile: "test.json",
				OutputExt:  "json",
			},
			wantErr:    true,
			errMessage: "error in parsing file",
		},
		{
			name: "Writing error",
			converter: &Converter{
				InputFile:  "test.csv",
				InputExt:   "csv",
				OutputFile: "error.json",
				OutputExt:  "json",
			},
			wantErr:    true,
			errMessage: "error in writing output file",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.converter.Convert()

			if (err != nil) != tt.wantErr {
				t.Errorf("Convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr && err != nil {
				if tt.errMessage != "" && !containsString(err.Error(), tt.errMessage) {
					t.Errorf("Convert() error message = %v, want to contain %v", err.Error(), tt.errMessage)
				}
			}
		})
	}
}

func containsString(s, substr string) bool {
	return len(s) >= len(substr) && s[0:len(substr)] == substr
}
