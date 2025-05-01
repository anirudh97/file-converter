package parser

import (
	"fmt"

	"github.com/anirudh97/file-converter/internal/models"
)

type Parser interface {
	Parse(inputFile string) (models.DataRecord, error)
}

var parsers = make(map[string]Parser)

func parserExists(format string) bool {
	_, exists := parsers[format]

	return exists
}

func Register(format string, parser Parser) error {

	if parserExists(format) {
		return fmt.Errorf("parser already exists for format: %s", format)
	}

	parsers[format] = parser

	return nil
}

func GetParser(format string) (Parser, error) {
	if !parserExists(format) {
		return nil, fmt.Errorf("no parser registered for format: %s", format)
	}

	return parsers[format], nil
}

func ClearRegistry() {
	parsers = make(map[string]Parser)
}
