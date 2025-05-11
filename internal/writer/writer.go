package writer

import (
	"fmt"

	"github.com/anirudh97/file-converter/internal/models"
)

type Writer interface {
	Write(data models.DataRecord, outputFile string) error
}

var writers = make(map[string]Writer)

func writerExists(format string) bool {
	_, exists := writers[format]

	return exists
}

func Register(format string, writer Writer) error {

	if writerExists(format) {
		return fmt.Errorf("writer already exists for format: %s", format)
	}

	writers[format] = writer

	return nil
}

func GetWriter(format string) (Writer, error) {
	if !writerExists(format) {
		return nil, fmt.Errorf("no writer registered for format: %s", format)
	}

	return writers[format], nil
}
func ClearRegistry() {
	writers = make(map[string]Writer)
}
