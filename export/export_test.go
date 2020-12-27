package export

import (
	"KindleHighlightsReader/reader"
	"log"
	"os"
	"strings"
	"testing"
)

func TestAsJSON(t *testing.T) {
	path, _ := AsJSON(getTestHighlights())

	if !strings.Contains(path, "JSON.txt") {
		t.Error("Failed to find JSON title in filename: ", path)
	}

	if !fileExist(path) {
		t.Error("Failed to find JSON file: ", path)
	}

	t.Cleanup(func() {
		if err := os.Remove(path); err != nil {
			log.Println(err)
		}
	})
}

func TestAsTxt(t *testing.T) {
	path, _ := AsTxt(getTestHighlights())

	if !strings.Contains(path, "TEXT.txt") {
		t.Error("Failed to find TEXT title in filename: ", path)
	}

	if !fileExist(path) {
		t.Error("Failed to find txt file: ", path)
	}
	t.Cleanup(func() {
		if err := os.Remove(path); err != nil {
			log.Println(err)
		}
	})
}

func TestAsCSV(t *testing.T) {
	path, _ := AsCSV(getTestHighlights())

	if !strings.Contains(path, "CSV.txt") {
		t.Error("Failed to find CSV title in filename: ", path)
	}

	if !fileExist(path) {
		t.Error("Failed to find CSV file: ", path)
	}
	t.Cleanup(func() {
		if err := os.Remove(path); err != nil {
			log.Println(err)
		}
	})
}

func TestAsPDF(t *testing.T) {
	path, _ := AsPDF(getTestHighlights())

	if !strings.Contains(path, ".pdf") {
		t.Error("Failed to find pdf extension in filename: ", path)
	}

	if !fileExist(path) {
		t.Error("Failed to find pdf file: ", path)
	}
	t.Cleanup(func() {
		if err := os.Remove(path); err != nil {
			log.Println(err)
		}
	})
}

func TestExport(t *testing.T) {
	filename := "My Clippings Test.pdf"
	f, _ := export(filename, nil)

	if !fileExist(f) {
		t.Error("Failed to find test file:", filename)
	}

	t.Cleanup(func() {
		if err := os.Remove(filename); err != nil {
			log.Print("Failed to remove test file: ", filename)
		}
	})
}

func fileExist(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func getTestHighlights() []reader.Highlight {
	h1 := reader.Highlight{
		Text:   "text",
		Author: "author",
		Title:  "title",
	}

	h2 := reader.Highlight{
		Text:   "text",
		Author: "author",
		Title:  "title",
	}
	return []reader.Highlight{h1, h2}
}
