package converter

import (
	"fmt"

	"github.com/anirudh97/file-converter/internal/parser"

	"github.com/anirudh97/file-converter/internal/writer"
)

type Converter struct {
	InputFile  string
	InputExt   string
	OutputFile string
	OutputExt  string
}

func (c *Converter) Convert() error {

	p, err := parser.GetParser(c.InputExt)
	if err != nil {
		return fmt.Errorf("error in getting parser: %v", err)
	}

	data, err := p.Parse(c.InputFile)
	if err != nil {
		return fmt.Errorf("error in parsing file: %v", err)
	}

	w, err := writer.GetWriter(c.OutputExt)
	if err != nil {
		return fmt.Errorf("error in getting writer: %v", err)
	}

	err = w.Write(data, c.OutputFile)
	if err != nil {
		return fmt.Errorf("error in writing output file: %v", err)
	}

	return nil

}
