package convert

import (
	"KindleHighlightsReader/highlight"
	"testing"
)

func TestToText(t *testing.T) {
	h := getTestHighlights()
	b := ToText(h)
	result := string(b)
	validText :=
		"text\n\nauthor, title\n________________________________\n\n" +
			"text\n\nauthor, title\n________________________________\n\n"

	if result != validText {
		t.Error("Failed to convert highlights to text. Data missing")
	}
}

func TestToJSON(t *testing.T) {
	h := getTestHighlights()
	b := ToJSON(h)
	result := string(b)
	validJson := `[{"Title":"title","Author":"author","Text":"text"},{"Title":"title","Author":"author","Text":"text"}]`
	if result != validJson {
		t.Errorf("Failed to correctly convert highlights to JSON")
	}
}

func TestToCSV(t *testing.T) {
	h := getTestHighlights()
	b := ToCSV(h)
	result := string(b)
	validCSV := "Title,Author,Text\ntitle,author,text\ntitle,author,text\n"
	if result != validCSV {
		t.Error("Failed to convert highlights to CSV. No headers was found")
	}
}

func TestToPDF(t *testing.T) {
	h := getTestHighlights()
	b := ToPDF(h)
	if len(b) < len(h) {
		t.Error("Failed to convert highlights to PDF. Data missing")
	}
}

func getTestHighlights() []highlight.Highlight {
	h1 := highlight.New("title", "author", "text")
	h2 := highlight.New("title", "author", "text")
	return []highlight.Highlight{h1, h2}
}
