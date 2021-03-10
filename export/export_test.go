package export

import (
	"KindleHighlightsReader/highlight"
	"log"
	"os"
	"strings"
	"testing"
)

func TestAsTxt(t *testing.T) {
	path, err := AsTxt(getTestHighlights())
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(path, ".txt") {
		t.Error("failed to find .txt extension in filename:", path)
	}
	if !fileExist(path) {
		t.Error("failed to find exported file:", path)
	}
	t.Cleanup(func() {
		if err := os.Remove(path); err != nil {
			log.Println(err)
		}
	})
}

func TestAsJSON(t *testing.T) {
	path, err := AsJSON(getTestHighlights())
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(path, ".json") {
		t.Error("failed to find .json extension in filename:", path)
	}
	if !fileExist(path) {
		t.Error("failed to find exported file:", path)
	}
	t.Cleanup(func() {
		if err := os.Remove(path); err != nil {
			log.Println(err)
		}
	})
}

func TestAsCSV(t *testing.T) {
	path, err := AsCSV(getTestHighlights())
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(path, ".csv") {
		t.Error("failed to find .csv extension in filename:", path)
	}
	if !fileExist(path) {
		t.Error("failed to find exported file:", path)
	}
	t.Cleanup(func() {
		if err := os.Remove(path); err != nil {
			log.Println(err)
		}
	})
}

func TestAsPDF(t *testing.T) {
	path, err := AsPDF(getTestHighlights())
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(path, ".pdf") {
		t.Error("failed to find .pdf extension in filename:", path)
	}
	if !fileExist(path) {
		t.Error("failed to find exported file:", path)
	}
	t.Cleanup(func() {
		if err := os.Remove(path); err != nil {
			log.Println(err)
		}
	})
}

func TestExport(t *testing.T) {
	filename := "MyClippingsTest.pdf"
	f, err := export(nil, filename)
	if err != nil {
		t.Error(err)
	}
	if !fileExist(f) {
		t.Error("failed to find test file:", filename)
	}

	t.Cleanup(func() {
		if err := os.Remove(f); err != nil {
			log.Print("failed to remove test file:", filename)
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

func getTestHighlights() []highlight.Highlight {
	h1 := highlight.New("title", "author", "title")
	h2 := highlight.New("title", "author", "title")
	return []highlight.Highlight{h1, h2}
}
