package save

import (
	"KindleHighlightsReader/convert"
	"KindleHighlightsReader/reader"
	"fmt"
	"log"
	"os"
)

func ToJSON(highlights []reader.Highlight) bool {
	filename := fmt.Sprintf("%s\\%s", getUserDesktopPath(), "My Clippings JSON.txt")
	b := convert.ToJSON(highlights)
	if f := createNewFile(filename); f != nil {
		return writeToFile(f, b)
	}
	return false
}

func ToTxt(highlights []reader.Highlight) bool {
	filename := fmt.Sprintf("%s\\%s", getUserDesktopPath(), "My Clippings Text.txt")
	b := convert.ToText(highlights)
	if f := createNewFile(filename); f != nil {
		return writeToFile(f, b)
	}
	return false
}

//Todo Return error or bool
func ToPDF(highlights []reader.Highlight) bool {
	filename := fmt.Sprintf("%s\\%s", getUserDesktopPath(), "My Clippings.pdf")
	b := convert.ToPDF(highlights)
	if f := createNewFile(filename); f != nil {
		return writeToFile(f, b)
	}
	return false
}

//Todo Return error or bool?
func ToCSV(highlights []reader.Highlight) bool {
	filename := fmt.Sprintf("%s\\%s", getUserDesktopPath(), "My Clippings.csv")
	b := convert.ToCSV(highlights)
	if f := createNewFile(filename); f != nil {
		return writeToFile(f, b)
	}
	return false
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
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
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
	defer func() {
		if err := f.Close(); err != nil {
			log.Println("Failed to close file:", f.Name())
		}
	}()
	return i == len(b)
}

//Todo Return error too?
func getUserDesktopPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Println(err)
	}
	return fmt.Sprintf("%s\\Desktop", homeDir)
}
