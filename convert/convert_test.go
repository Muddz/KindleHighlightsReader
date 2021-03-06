package convert

import (
	"KindleHighlightsReader/highlight"
	"testing"
)

func TestToText(t *testing.T) {
	h := getTestHighlights()
	b := ToText(h)
	actual := string(b)
	expected := "text\n\nauthor, title\n________________________________\n\ntext\n\nauthor, title\n________________________________\n\n"
	if actual != expected {
		t.Error("failed to convert highlights to text.")
	}
}

func TestToJSON(t *testing.T) {
	h := getTestHighlights()
	b, err := ToJSON(h)
	if err != nil {
		t.Error(err)
	}
	actual := string(b)
	expected := `[{"Title":"title","Author":"author","Text":"text"},{"Title":"title","Author":"author","Text":"text"}]`
	if actual != expected {
		t.Errorf("failed to correctly convert highlights to JSON")
	}
}

func TestToCSV(t *testing.T) {
	h := getTestHighlights()
	b, err := ToCSV(h)
	if err != nil {
		t.Error(err)
	}
	actual := string(b)
	expected := "Title,Author,Text\ntitle,author,text\ntitle,author,text\n"
	if actual != expected {
		t.Error("failed to convert highlights to CSV")
	}
}

func TestToPDF(t *testing.T) {
	h := getTestHighlights()
	b, err := ToPDF(h)
	if err != nil {
		t.Error(err)
	}
	if len(b) < len(h) {
		t.Error("failed to convert highlights to PDF")
	}
}

func getTestHighlights() []highlight.Highlight {
	h1 := highlight.New("title", "author", "text")
	h2 := highlight.New("title", "author", "text")
	return []highlight.Highlight{h1, h2}
}
