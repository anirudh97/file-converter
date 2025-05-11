package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/anirudh97/file-converter/internal/converter"
	"github.com/anirudh97/file-converter/pkg/utils"
)

func main() {

	// command line flags
	inputFile := flag.String("input", "", "Input file path")
	outputFile := flag.String("output", "", "Output file path")

	flag.Parse()

	fmt.Println(*inputFile)
	// basic validation

	if *inputFile == "" || *outputFile == "" {
		fmt.Println("Error: Both input and output files should be specified")
		fmt.Println("Usage: fileconverter -input <input file path> -output <output file path>")
		flag.PrintDefaults()
		os.Exit(1)
	}

	inputExt := path.Ext(*inputFile)[1:]
	outputExt := path.Ext(*outputFile)[1:]

	if err := utils.ValidateFile(*inputFile); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	if err := utils.ValidateSupportedExtension(inputExt); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	if err := utils.ValidateSupportedExtension(outputExt); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	conv := &converter.Converter{
		InputFile:  *inputFile,
		InputExt:   inputExt,
		OutputFile: *outputFile,
		OutputExt:  outputExt,
	}

	if err := conv.Convert(); err != nil {
		fmt.Printf("Error during conversion: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully converted %s to %s \n", *inputFile, *outputFile)

}
