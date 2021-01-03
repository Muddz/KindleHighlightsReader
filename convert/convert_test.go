package convert

import (
	"KindleHighlightsReader/reader"
	"testing"
)

func TestToText(t *testing.T) {
	h := getTestHighlights()
	b := ToText(h)
	c := string(b)
	validText :=
		"text\n\nauthor, title\n________________________________\n\n" +
			"text\n\nauthor, title\n________________________________\n\n"

	if c != validText {
		t.Error("Failed to convert highlights to text. Data missing")
	}
}

func TestToJSON(t *testing.T) {
	h := getTestHighlights()
	b := ToJSON(h)
	c := string(b)
	validJson := `[{"Text":"text","Author":"author","Title":"title"},{"Text":"text","Author":"author","Title":"title"}]`
	if c != validJson {
		t.Errorf("Failed to correctly convert highlights to JSON")
	}
}

func TestToCSV(t *testing.T) {
	h := getTestHighlights()
	b := ToCSV(h)
	c := string(b)
	validCSV := "Text,Author,Title\ntext,author,title\ntext,author,title\n"
	if c != validCSV {
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
