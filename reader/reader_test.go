package reader

import (
	"testing"
)

//Todo refactor this test
func TestReadHighlightFile(t *testing.T) {
	highlights := ReadHighlights("./My Clippings_test.txt")
	if len(highlights) != 2 {
		t.Error("Failed to read all highlights from file")
	}
}

//Todo refactor this test
func TestGetFileContent(t *testing.T) {
	testFile := "My Clippings_test.txt"
	content := getFileContent(testFile)
	if len(content) < 1 {
		t.Errorf("Failed to get any content of from file %s", testFile)
	}
}

//Todo refactor this test/Should we just read the testfile?
func TestHighlightParser(t *testing.T) {
	testString :=
		`==========
Book title (Author)
- Highlight information

Text
==========
Book title (Author)
- Highlight information

Text`
	highlights := parseHighlights(testString)
	for _, v := range highlights {
		if v.Title != "Book title" && v.Author != "Author" && v.Text != "Text" {
			t.Error("Failed to parse highlights to struct")
		}
	}
}

func TestRemoveControlChars(t *testing.T) {
	testString := "Hello\r"
	result := removeCarriageReturn(testString)
	if result != "Hello" {
		t.Errorf("Failed to clean text for control characters from '%s'", testString)
	}
}

func TestRemoveBOM(t *testing.T) {
	testString := "\xef\xbb\xbfHello"
	result := removeBOM(testString)
	if result != "Hello" {
		t.Errorf("Failed to clean text for BOM '%s'", testString)
	}
}
