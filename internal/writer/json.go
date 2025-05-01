package writer

import (
	"encoding/json"
	"os"

	"github.com/anirudh97/file-converter/internal/models"
)

type JSONWriter struct{}

func init() {
	Register("json", &JSONWriter{})
}

func (c *JSONWriter) Write(data models.DataRecord, outputFile string) error {

	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")

	return encoder.Encode(data)

}
