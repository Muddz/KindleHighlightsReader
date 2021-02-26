package convert

import (
	"KindleHighlightsReader/highlight"
	"bufio"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/phpdave11/gofpdf"
	"reflect"
	"strings"
)

func ToText(h []highlight.Highlight) []byte {
	layout := "%s\n\n%s, %s\n________________________________\n\n"
	sb := strings.Builder{}
	for _, v := range h {
		row := fmt.Sprintf(layout, v.Text, v.Author, v.Title)
		sb.WriteString(row)
	}
	b := []byte(sb.String())
	return b
}

func ToJSON(h []highlight.Highlight) ([]byte, error) {
	b, err := json.Marshal(h)
	if err != nil {
		return nil, fmt.Errorf("%w\n", err)
	}
	return b, nil
}

func ToCSV(highlights []highlight.Highlight) ([]byte, error) {
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
		return nil, fmt.Errorf("%w\n", err)
	}

	for _, v := range highlights {
		row := []string{v.Title, v.Author, v.Text}
		if err := writer.Write(row); err != nil {
			return nil, fmt.Errorf("%w\n", err)
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return nil, fmt.Errorf("%w\n", err)
	}
	return b.Bytes(), nil
}

func ToPDF(h []highlight.Highlight) ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	for i, v := range h {
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
		return nil, fmt.Errorf("%w\n", err)
	}
	if err := w.Flush(); err != nil {
		return nil, fmt.Errorf("%w\n", err)
	}
	return b.Bytes(), nil
}

func utfToCP1252(text string) string {
	//https://github.com/djimenez/iconv-go
	//https://github.com/signintech/gopdf
	//https://godoc.org/github.com/jung-kurt/gofpdf#example-Fpdf-CellFormat-Codepage
	return ""
}
