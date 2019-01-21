package main

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
	pdf "github.com/unidoc/unidoc/pdf/model"
)

func main() {
	input := flag.String("input", "", "input file path")
	output := flag.String("output", "", "output file path")
	rows := flag.Int("rows", 1, "divide each page into this many rows, 1 means no change")
	cols := flag.Int("cols", 1, "divide each page into this columns, 1 means no change")
	row := flag.Int("row", 1, "row number of the crop")
	col := flag.Int("col", 1, "column number of the crop")

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

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		panic(fmt.Sprintf("error getting number of pages: %v", err))
	}

	for i := 1; i <= numPages; i++ {
		page, err := pdfReader.GetPage(i)
		if err != nil {
			panic(fmt.Sprintf("error getting page %d: %v", i, err))
		}

		mbox, err := page.GetMediaBox()
		if err != nil {
			panic(fmt.Sprintf("error reading page %d: %v", i, err))
		}

		width := (*mbox).Urx - (*mbox).Llx
		height := (*mbox).Ury - (*mbox).Lly

		croppedWidth := width / float64(*cols)
		croppedHeight := height / float64(*rows)

		xs := make([]float64, *cols+1) // x coordinates of every corner
		ys := make([]float64, *rows+1) // y coordinates of every corner

		for j := 0; j <= *cols; j++ {
			xs[j] = croppedWidth * float64(j)
		}

		for j := 0; j <= *rows; j++ {
			ys[j] = croppedHeight * float64(j)
		}

		x := *col - 1
		y := *rows - *row

		// every (x,y) is a lower-left coordinate
		(*mbox).Llx = xs[x]
		(*mbox).Lly = ys[y]
		(*mbox).Urx = xs[x+1]
		(*mbox).Ury = ys[y+1]

		// crop the page
		page.MediaBox = mbox

		err = pdfWriter.AddPage(page)
		if err != nil {
			panic(fmt.Sprintf("error cropping page %d: %v", i, err))
		}

	}

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
