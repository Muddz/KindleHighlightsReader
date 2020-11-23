package convert

import (
	"KindleHighlightsReader/reader"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

//Todo* Return error or bool with the bytes?

//TODO 1) Remove BOM and /r characters
func ToJSON(highlights []reader.Highlight) []byte {
	b, err := json.Marshal(highlights)
	if err != nil {
		log.Println(err)
	}
	return b
}

func ToText(highlights []reader.Highlight) []byte {
	layout := "%s, %s\n\n%s\n_______________________________\n\n"
	sb := strings.Builder{}
	for _, v := range highlights {
		highlight := fmt.Sprintf(layout, v.Author, v.Title, v.Text)
		sb.WriteString(highlight)
	}
	b := []byte(sb.String())
	return b
}

func ToCSV(highlights []reader.Highlight) []byte {
	return nil
}

func ToPDF(highlights []reader.Highlight) []byte {
	return nil
}

func cleanJSON(json string) string {
	return ""
}
