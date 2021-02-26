package convert

import (
	"KindleHighlightsReader/highlight"
	"fmt"
	"testing"
)

func TestError(t *testing.T) {

	err1 := fmt.Errorf("failed converting highlights to json\n")
	err2 := fmt.Errorf("failed at marsheling highlights to json string\n > %w", err1)
	err3 := fmt.Errorf("couldn't find file /Muddz/Desktop/my_clippings.txt\n > %w", err2)
	err4 := fmt.Errorf("failed at loading user file\n > %w", err3)

	fmt.Print(err4)

}

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
	b, err := ToJSON(h)
	if err != nil {
		t.Error(err)
	}

	result := string(b)
	validJson := `[{"Title":"title","Author":"author","Text":"text"},{"Title":"title","Author":"author","Text":"text"}]`
	if result != validJson {
		t.Errorf("Failed to correctly convert highlights to JSON")
	}
}

func TestToCSV(t *testing.T) {
	h := getTestHighlights()
	b, err := ToCSV(h)
	if err != nil {
		t.Error(err)
	}
	result := string(b)
	validCSV := "Title,Author,Text\ntitle,author,text\ntitle,author,text\n"
	if result != validCSV {
		t.Error("Failed to convert highlights to CSV. No headers was found")
	}
}

func TestToPDF(t *testing.T) {
	h := getTestHighlights()
	b, err := ToPDF(h)
	if err != nil {
		t.Error(err)
	}
	if len(b) < len(h) {
		t.Error("Failed to convert highlights to PDF. Data missing")
	}
}

func getTestHighlights() []highlight.Highlight {
	h1 := highlight.New("title", "author", "text")
	h2 := highlight.New("title", "author", "text")
	return []highlight.Highlight{h1, h2}
}
