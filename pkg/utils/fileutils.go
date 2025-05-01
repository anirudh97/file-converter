package utils

import (
	"fmt"
	"os"
	"strings"
)

func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)

	return err == nil

}

func ValidateFile(filePath string) error {

	if !FileExists(filePath) {
		return fmt.Errorf("input file does not exist: %s", filePath)
	}

	file, err := os.Open(filePath)

	if err != nil {
		return fmt.Errorf("cannot open file: %v", err)
	}

	defer file.Close()

	return nil
}

func ValidateSupportedExtension(ext string) error {

	supportedFormats := map[string]bool{
		"csv":  true,
		"json": true,
	}

	if !supportedFormats[ext] {
		var formats []string

		for f := range supportedFormats {
			formats = append(formats, f)
		}
		return fmt.Errorf("unsupported file format: %s. Supported formats: %s", ext, strings.Join(formats, ", "))
	}

	return nil
}
