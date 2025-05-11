package parser

import (
	"os"
	"testing"
)

func TestCSVParser(t *testing.T) {

	tmpfile, err := os.CreateTemp("", "test-*.csv")
	if err != nil {
		t.Fatalf("Failed to create temp : %v", err)
	}

	defer os.Remove(tmpfile.Name())

	content := "name,age,city\nJohn,30,New York\nAlice,25,Seattle"
	if _, err := tmpfile.Write([]byte(content)); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}

	parser := &CSVParser{}
	result, err := parser.Parse(tmpfile.Name())

	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	if len(result) != 2 {
		t.Errorf("Parse() returned %d records, expected 2", len(result))
	}

	if result[0]["name"] != "John" {
		t.Errorf("First record name = %v, expected 'John'", result[0]["name"])
	}
	if result[1]["city"] != "Seattle" {
		t.Errorf("Second record city = %v, expected 'Seattle'", result[1]["city"])
	}

}
