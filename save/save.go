package save

import (
	"KindleHighlightsReader/reader"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

//TODO 1) Remove BOM and /r characters
//func ToJSON(highlights []reader.Highlight, destination string) bool {
//	b, err := json.Marshal(highlights)
//	if err != nil {
//		log.Println(err)
//	}
//	filename := fmt.Sprintf("%s/%s", destination, "My Clippings Json.txt")
//	if f := createNewFile(filename); f != nil {
//		return writeToFile(f, b)
//	}
//	return false
//}

func ToJSON(highlights []reader.Highlight, destination string) bool {
	filename := fmt.Sprintf("%s/%s", destination, "My Clippings Json.txt")
	b := convertToJSON(highlights)
	if f := createNewFile(filename); f != nil {
		return writeToFile(f, b)
	}
	return false
}

func convertToJSON(highlights []reader.Highlight) []byte {
	b, err := json.Marshal(highlights)
	if err != nil {
		log.Println(err)
	}
	return b
}

//func ToTxt(highlights []reader.Highlight, destination string) bool {
//	layout := "%s, %s\n\n%s\n_______________________________\n\n"
//	sb := strings.Builder{}
//	for _, v := range highlights {
//		highlight := fmt.Sprintf(layout, v.Author, v.Title, v.Text)
//		sb.WriteString(highlight)
//	}
//	filename := fmt.Sprintf("%s/%s", destination, "My Clippings text.txt")
//	if f := createNewFile(filename); f != nil {
//		return writeToFile(f, []byte(sb.String()))
//	}
//	return false
//}

func ToTxt(highlights []reader.Highlight, destination string) bool {
	txt := convertToText(highlights)
	filename := fmt.Sprintf("%s/%s", destination, "My Clippings text.txt")
	if f := createNewFile(filename); f != nil {
		return writeToFile(f, []byte(txt))
	}
	return false
}

func convertToText(highlights []reader.Highlight) string {
	layout := "%s, %s\n\n%s\n_______________________________\n\n"
	sb := strings.Builder{}
	for _, v := range highlights {
		highlight := fmt.Sprintf(layout, v.Author, v.Title, v.Text)
		sb.WriteString(highlight)
	}
	return sb.String()
}

func toPDF() {

}

func toCSV(highlights []reader.Highlight, destination string) {

	if f := createNewFile(filename); f != nil {
		return writeToFile(f, []byte(txt))
	}
}

func convertToCSV(highlights []reader.Highlight) string {
	return ""
}

func writeToFile(f *os.File, b []byte) bool {
	i, err := f.Write(b)
	if err != nil {
		log.Println("Failed to write to file:", f.Name())
		return false
	}
	if err := f.Close(); err != nil {
		log.Println("Failed to close file:", f.Name())
	}
	return i == len(b)
}

func createNewFile(filename string) *os.File {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Println("Failed to create file:", filename)
		return nil
	}
	return file
}
