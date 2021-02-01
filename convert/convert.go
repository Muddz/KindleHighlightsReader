package convert

import (
	"KindleHighlightsReader/reader"
	"bufio"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/phpdave11/gofpdf"
	"log"
	"reflect"
	"strings"
)

//TODO should these methods return the error too?
func ToText(highlights []reader.Highlight) []byte {
	layout := "%s\n\n%s, %s\n________________________________\n\n"
	sb := strings.Builder{}
	for _, v := range highlights {
		highlight := fmt.Sprintf(layout, v.Text, v.Author, v.Title)
		sb.WriteString(highlight)
	}
	b := []byte(sb.String())
	return b
}

func ToJSON(highlights []reader.Highlight) []byte {
	b, err := json.Marshal(highlights)
	if err != nil {
		log.Println(err)
	}
	return b
}

func ToCSV(highlights []reader.Highlight) []byte {
	var headers []string
	h := highlights[0]
	v := reflect.ValueOf(&h).Elem()
	typeOf := v.Type()
	for i := 0; i < v.NumField(); i++ {
		header := typeOf.Field(i).Name
		headers = append(headers, header)
	}

	var b bytes.Buffer
	writer := csv.NewWriter(&b)
	if err := writer.Write(headers); err != nil {
		log.Println("Failed to write csv headers", err)
	}

	for _, v := range highlights {
		row := []string{v.Title, v.Author, v.Text}
		if err := writer.Write(row); err != nil {
			log.Println("Failed to write csv values:", err)
		}
	}
	writer.Flush()

	if err := writer.Error(); err != nil {
		log.Println("Failed to flush csv writer,", err)
		return nil
	}
	return b.Bytes()
}

func ToPDF(highlights []reader.Highlight) []byte {
	pdf := gofpdf.New("P", "mm", "A4", "")
	for i, v := range highlights {
		if i%5 == 0 {
			pdf.AddPage()
		}
		text := fmt.Sprintf("\n%s", v.Text)
		pdf.SetFont("Arial", "", 14)
		pdf.MultiCell(0, 10, text, "0", "0", false)

		author := fmt.Sprintf("%s, %s\n\n", v.Author, v.Title)
		author = strings.TrimRight(author, "\r\n")
		pdf.SetFont("Arial", "i", 10)
		pdf.MultiCell(0, 10, author, "0", "0", false)
	}
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	if err := pdf.Output(w); err != nil {
		log.Println("Failed to output PDF")
	}
	if err := w.Flush(); err != nil {
		log.Println("Failed to flush PDF writer")
	}
	return b.Bytes()
}

func utfToCP1252(text string) string {
	//https://github.com/djimenez/iconv-go
	//https://github.com/signintech/gopdf
	//https://godoc.org/github.com/jung-kurt/gofpdf#example-Fpdf-CellFormat-Codepage
	return ""
}
