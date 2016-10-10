package main

import (
	"flag"
	"github.com/jung-kurt/gofpdf"
)

func main() {
	var inImg string
	var outPdf string

	flag.StringVar(&inImg, "i", "", "The image to be converted")
	flag.StringVar(&outPdf, "o", "", "The name of the pdf that will be generated,\n\t if no name is specified the name will match that of the input image")
	flag.Parse()

	if inImg == "" {
		flag.PrintDefaults()
		return
	}

	if outPdf == "" {
		outPdf = inImg[:(len(inImg)-4)] + ".pdf"
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(10, 10, 10)
	pdf.AddPage()
	_, pgWidth := pdf.GetPageSize()
	img := pdf.RegisterImage(inImg, "")
	offset := (pgWidth - img.Width()) / 3.5
	pdf.ImageOptions(inImg, offset, 10, 0, 0, true, gofpdf.ImageOptions{ReadDpi: true}, 0, "")
	pdf.OutputFileAndClose(outPdf)
}
