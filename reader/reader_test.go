package reader

import (
	"testing"
)

func TestReadHighlights(t *testing.T) {
	highlights, err := ReadHighlights("MyClippings.txt")
	if err != nil {
		t.Fatal(err)
	}

	if len(highlights) != 2 {
		t.Fatalf("failed to parse test highlights, expected: 2 highlights actual was: %d", len(highlights))
	}

	for _, v := range highlights {
		if v.Title != "Title" || v.Author != "Author" || v.Text != "Text" {
			t.Errorf("failed to find correct highlight proporties, actual was: \n%s", v)
		}
	}
}

func TestRemoveBOM(t *testing.T) {
	input := "\xef\xbb\xbfHello"
	actual := removeBOM(input)
	expected := "Hello"
	if actual != expected {
		t.Errorf("failed to remove BOM from: '%s' expected: %s", actual, expected)
	}
}
