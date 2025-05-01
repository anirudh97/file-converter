package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFileExists(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "fileexists-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	existingFile := tmpfile.Name()
	defer os.Remove(existingFile)
	tmpfile.Close()

	nonExistingFile := filepath.Join(os.TempDir(), "non-existing-file.txt")
	os.Remove(nonExistingFile)

	tests := []struct {
		name     string
		filePath string
		want     bool
	}{
		{
			name:     "Existing file",
			filePath: existingFile,
			want:     true,
		},
		{
			name:     "Non-existing file",
			filePath: nonExistingFile,
			want:     false,
		},
		{
			name:     "Empty path",
			filePath: "",
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FileExists(tt.filePath); got != tt.want {
				t.Errorf("FileExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateFile(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "validatefile-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	existingFile := tmpfile.Name()
	defer os.Remove(existingFile)
	tmpfile.Close()

	nonExistingFile := filepath.Join(os.TempDir(), "non-existing-validate.txt")
	os.Remove(nonExistingFile)

	tests := []struct {
		name     string
		filePath string
		wantErr  bool
	}{
		{
			name:     "Valid file",
			filePath: existingFile,
			wantErr:  false,
		},
		{
			name:     "Non-existing file",
			filePath: nonExistingFile,
			wantErr:  true,
		},
		{
			name:     "Empty path",
			filePath: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateFile(tt.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateSupportedExtension(t *testing.T) {
	tests := []struct {
		name    string
		ext     string
		wantErr bool
	}{
		{
			name:    "CSV extension",
			ext:     "csv",
			wantErr: false,
		},
		{
			name:    "JSON extension",
			ext:     "json",
			wantErr: false,
		},
		{
			name:    "Unsupported extension",
			ext:     "txt",
			wantErr: true,
		},
		{
			name:    "Empty extension",
			ext:     "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateSupportedExtension(tt.ext)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateSupportedExtension() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
