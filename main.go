package main

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
	pdf "github.com/unidoc/unidoc/pdf/model"
	"github.com/zihaoyu/cut-pdf/split"
)

func main() {
	input := flag.String("input", "", "input file path")
	output := flag.String("output", "", "output file path")
	rows := flag.Int("rows", 1, "divide each page into this many rows, 1 means no change")
	cols := flag.Int("cols", 1, "divide each page into this columns, 1 means no change")

	flag.Parse()

	fInput, err := os.Open(*input)
	if err != nil {
		panic(fmt.Sprintf("error opening input file: %v", err))
	}
	defer fInput.Close()

	pdfReader, err := pdf.NewPdfReader(fInput)
	if err != nil {
		panic(fmt.Sprintf("error opening pdf: %v", err))
	}

	pdfWriter := pdf.NewPdfWriter()

	split.Split(*rows, *cols, pdfReader, &pdfWriter)

	fOutput, err := os.Create(*output)
	if err != nil {
		panic(fmt.Sprintf("error creating output file: %v", err))
	}
	defer fOutput.Close()

	err = pdfWriter.Write(fOutput)
	if err != nil {
		panic(fmt.Sprintf("error writing output file: %v", err))
	}

}
