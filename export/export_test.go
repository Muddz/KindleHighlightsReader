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
	check(t, path, ".txt", err)
}

func TestAsJSON(t *testing.T) {
	path, err := AsJSON(getTestHighlights())
	check(t, path, ".json", err)
}

func TestAsCSV(t *testing.T) {
	path, err := AsCSV(getTestHighlights())
	check(t, path, ".csv", err)
}

func TestAsPDF(t *testing.T) {
	path, err := AsPDF(getTestHighlights())
	check(t, path, ".pdf", err)
}

func check(t *testing.T, path string, extension string, err error) {
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(path, extension) {
		t.Errorf("failed to find %s extension in filename: %s", extension, path)
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
