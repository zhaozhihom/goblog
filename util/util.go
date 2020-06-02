package util

import (
	"blog/config"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

var logger = config.GetLogger()

// CovertPageToOffset 把页码和页数转换为开始和偏移量
// 2, 5 => 5, 5
func CovertPageToOffset(page string, size string) (int, int) {
	isize, _ := strconv.Atoi(size)
	ipage, _ := strconv.Atoi(page)
	return isize * (ipage - 1), isize
}


// CovertHTMLToPdf 把html字符串转换为pdf
func CovertHTMLToPdf(html *string) (string, error) {
	// Create new PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		logger.Error(err)
	}

	// Set global options
	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	// pdfg.Cover.UserStyleSheet.Set("true")
	// pdfg.Grayscale.Set(true)

	htmlplus := `<html><head><meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
									<style>
										html { font-family: "Open Sans","Clear Sans","Helvetica Neue",Helvetica,Arial,sans-serif;												 
													 color: rgb(51, 51, 51);
													 line-height: 1.6;
													}
										table {
														width:100%;
												}
									</style>
								</head>` + *html	+ "</html>"
	pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(htmlplus)))

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		logger.Error(err)
	}

	now := time.Now()
	nowStr := now.Format(time.RFC3339)
	filePath := "./resumes/"+ nowStr + ".pdf"
	// Write buffer contents to file on disk
	err = pdfg.WriteFile(filePath)
	if err != nil {
		logger.Error(err)
	}

	log.Println("Done")
	return filePath, nil
} 