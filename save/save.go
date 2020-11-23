package save

import (
	"KindleHighlightsReader/convert"
	"KindleHighlightsReader/reader"
	"fmt"
	"log"
	"os"
)

func ToJSON(highlights []reader.Highlight, destination string) bool {
	filename := fmt.Sprintf("%s/%s", destination, "My Clippings Json.txt")
	b := convert.ToJSON(highlights)
	if f := createNewFile(filename); f != nil {
		return writeToFile(f, b)
	}
	return false
}

func ToTxt(highlights []reader.Highlight, destination string) bool {
	filename := fmt.Sprintf("%s/%s", destination, "My Clippings Text.txt")
	b := convert.ToText(highlights)
	if f := createNewFile(filename); f != nil {
		return writeToFile(f, b)
	}
	return false
}

func toPDF() {

}

func toCSV(highlights []reader.Highlight, destination string) {

	//ok := save("",nil)
}

//Todo Return error or bool?
func save(filename string, b []byte) bool {
	if f := createNewFile(filename); f != nil {
		return writeToFile(f, b)
	}
	return false
}

//Todo Return error or bool?
func createNewFile(filename string) *os.File {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Println("Failed to create file:", filename)
		return nil
	}
	return file
}

//Todo Return error or bool?
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
