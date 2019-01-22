package split

import (
	"fmt"

	pdf "github.com/unidoc/unidoc/pdf/model"
)

// Split splits a PDF document
func Split(rows, cols int, pdfReader *pdf.PdfReader, pdfWriter *pdf.PdfWriter) {
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

		croppedWidth := width / float64(cols)
		croppedHeight := height / float64(rows)

		xs := make([]float64, cols+1) // x coordinates of every corner
		ys := make([]float64, rows+1) // y coordinates of every corner

		for j := 0; j <= cols; j++ {
			xs[j] = croppedWidth * float64(j)
		}

		for j := 0; j <= rows; j++ {
			ys[j] = croppedHeight * float64(j)
		}

		for y := rows - 1; y >= 0; y-- {
			for x := 0; x < cols; x++ {
				// every (x,y) is a lower-left coordinate

				// duplicate the page for every cut-out section because we are passing by reference
				// so if the mediabox moves, every section of the page will be affected
				p := page.Duplicate()

				mmbox, err := p.GetMediaBox()
				if err != nil {
					panic(fmt.Sprintf("error processing page %d: %v", i, err))
				}
				(*mmbox).Llx = xs[x]
				(*mmbox).Lly = ys[y]
				(*mmbox).Urx = xs[x+1]
				(*mmbox).Ury = ys[y+1]

				// crop the page
				p.MediaBox = mmbox

				err = pdfWriter.AddPage(p)
				if err != nil {
					panic(fmt.Sprintf("error cropping page %d: %v", i, err))
				}
			}
		}
	}
}
