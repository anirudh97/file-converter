package writer

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"

	"github.com/anirudh97/file-converter/internal/models"
)

func TestJSONWriter_Write(t *testing.T) {
	testData := models.DataRecord{
		{"name": "John", "age": "30", "city": "New York"},
		{"name": "Alice", "age": "25", "city": "Seattle"},
	}

	tmpfile, err := os.CreateTemp("", "test-*.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	tmpFileName := tmpfile.Name()
	tmpfile.Close()
	defer os.Remove(tmpFileName)

	writer := &JSONWriter{}

	err = writer.Write(testData, tmpFileName)
	if err != nil {
		t.Fatalf("Write() error = %v", err)
	}

	file, err := os.Open(tmpFileName)
	if err != nil {
		t.Fatalf("Failed to open output file: %v", err)
	}
	defer file.Close()

	var result models.DataRecord
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&result); err != nil {
		t.Fatalf("Failed to decode output JSON: %v", err)
	}

	if !reflect.DeepEqual(testData, result) {
		t.Errorf("Write() produced incorrect JSON.\nExpected: %+v\nGot: %+v", testData, result)
	}
}

func TestJSONWriter_WriteError(t *testing.T) {

	testData := models.DataRecord{
		{"name": "John", "age": "30", "city": "New York"},
	}

	writer := &JSONWriter{}
	err := writer.Write(testData, "/nonexistent/directory/file.json")

	if err == nil {
		t.Error("Write() to invalid path did not produce an error")
	}
}

func TestJSONWriter_WriteEmptyData(t *testing.T) {

	testData := models.DataRecord{}

	tmpfile, err := os.CreateTemp("", "test-empty-*.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	tmpFileName := tmpfile.Name()
	tmpfile.Close()
	defer os.Remove(tmpFileName)

	writer := &JSONWriter{}
	err = writer.Write(testData, tmpFileName)
	if err != nil {
		t.Fatalf("Write() with empty data error = %v", err)
	}

	file, err := os.Open(tmpFileName)
	if err != nil {
		t.Fatalf("Failed to open output file: %v", err)
	}
	defer file.Close()

	var result models.DataRecord
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&result); err != nil {
		t.Fatalf("Failed to decode output JSON: %v", err)
	}

	if len(result) != 0 {
		t.Errorf("Expected empty array, got %+v", result)
	}
}
