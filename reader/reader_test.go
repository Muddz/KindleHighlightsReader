package reader

import (
	"fmt"
	"testing"
)

func TestReadHighlightFile(t *testing.T) {
	highlights := ReadHighlights("./My Clippings_test.txt")
	if len(highlights) != 2 {
		t.Error("Failed to read all highlights from file")
	}
}

func TestGetFileContent(t *testing.T) {
	testFile := "My Clippings_test.txt"
	content := getFileContent(testFile)
	if len(content) < 1 {
		t.Errorf("Failed to get any content of from file %s", testFile)
	}
}

func TestHighlightParser(t *testing.T) {
	testString :=
		`==========
Book title (Author)
- Highlight information

Quote
==========
Book title (Author)
- Highlight information

Quote`

	highlights := highlightsParser(testString)
	if len(highlights) != 2 {
		t.Errorf("Failed to parse the correct amount of highlight objects")
	}
	for _, highlight := range highlights {
		if len(highlight.Title) == 0 || len(highlight.Author) == 0 || len(highlight.Text) == 0 {
			t.Errorf("Failed to parse fields from highlight objects")
		}
	}
}

func TestRemoveAuthorFromTitle(t *testing.T) {
	testString := "The story of Microsoft (Bill Gates)"
	result := removeAuthorFromTitle(testString)
	if result != "The story of Microsoft" {
		t.Errorf("Failed to remove author from title %s", testString)
	}
}

func TestGetAuthorFromTitle(t *testing.T) {
	testString := "The story of Microsoft (Bill Gates)"
	result := getAuthorFromTitle(testString)
	if result != "Bill Gates" {
		t.Errorf("Failed to extract author's name from %s", testString)
	}
}

func TestRemoveAuthorParentheses(t *testing.T) {
	testString := "(Bill Gates)"
	result := removeAuthorParentheses(testString)
	if result != "Bill Gates" {
		t.Errorf("Failed to remove parentheses from %s", testString)
	}
}

func TestRemoveControlChars(t *testing.T) {
	testString := "Hello\r"
	fmt.Println(testString)
	result := removeControlChars(testString)
	if result != "Hello" {
		t.Errorf("Failed to clean text for control characters from %s", testString)
	}
}

func TestRemoveBOM(t *testing.T) {
	testString := " Hello"
	result := removeBOM(testString)
	if result != "Hello" {
		t.Errorf("Failed to clean text for BOM %s", testString)
	}
}
