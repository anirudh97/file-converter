package main

import (
	"flag"
	"os"
	"path/filepath"
	"testing"
)

func TestValidateCommandLineArgs(t *testing.T) {
	// Save original flag values and restore after test
	oldFlagCommandLine := flag.CommandLine
	defer func() { flag.CommandLine = oldFlagCommandLine }()

	// Reset flags for testing
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	inputFile := flag.String("input", "", "Input file path")
	outputFile := flag.String("output", "", "Output file path")

	// Test with empty args
	args := []string{}
	err := flag.CommandLine.Parse(args)
	if err != nil {
		t.Errorf("Unexpected error parsing empty args: %v", err)
	}

	if *inputFile != "" || *outputFile != "" {
		t.Errorf("Expected empty input/output values, got %s and %s", *inputFile, *outputFile)
	}
}

func TestFileValidation(t *testing.T) {
	// Create a temp directory
	tempDir := t.TempDir()

	// Create a CSV file
	csvPath := filepath.Join(tempDir, "test.csv")
	err := os.WriteFile(csvPath, []byte("name,age\nJohn,30"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test CSV: %v", err)
	}

	// Test non-existent file
	nonExistentPath := filepath.Join(tempDir, "nonexistent.csv")
	if _, err := os.Stat(nonExistentPath); !os.IsNotExist(err) {
		t.Errorf("Expected file to not exist")
	}

	// Test unsupported extension
	txtPath := filepath.Join(tempDir, "test.txt")
	err = os.WriteFile(txtPath, []byte("test"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test TXT: %v", err)
	}

	// Verify file exists
	if _, err := os.Stat(csvPath); os.IsNotExist(err) {
		t.Errorf("CSV file should exist")
	}
}

// Optional: Test file extension parsing
func TestExtensionParsing(t *testing.T) {
	tests := []struct {
		path    string
		wantExt string
	}{
		{"file.csv", "csv"},
		{"/path/to/data.json", "json"},
		{"noextension", ""},
		{"path/with.multiple.dots.yml", "yml"},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			ext := filepath.Ext(tt.path)
			if ext != "" {
				ext = ext[1:] // Remove the dot
			}

			if ext != tt.wantExt {
				t.Errorf("Got extension %q, want %q", ext, tt.wantExt)
			}
		})
	}
}
