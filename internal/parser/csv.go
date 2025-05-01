package parser

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/anirudh97/file-converter/internal/models"
	"github.com/anirudh97/file-converter/pkg/utils"
)

type CSVParser struct{}

func init() {
	Register("csv", &CSVParser{})
}

func (p *CSVParser) Parse(inputFile string) (models.DataRecord, error) {

	file, err := os.Open(inputFile)
	if err != nil {
		return nil, fmt.Errorf("error in opening the file %s: %v", inputFile, err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	headers, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("error in reading the headers: %v", err)
	}

	if len(headers) == 0 {
		return models.DataRecord{}, nil
	}

	var data models.DataRecord

	cols := len(headers)

	for {

		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("Failed to read row/record: %v", err)
		}

		row := make(map[string]interface{})
		for i, value := range record {
			if i < cols {
				row[headers[i]] = utils.DetectType(value)
			}
		}

		data = append(data, row)
	}

	return data, nil

}
