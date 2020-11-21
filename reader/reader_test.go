package reader

import (
	"testing"
)

func TestReadHighlightFile(t *testing.T) {
	highlights, err := ReadHighlightFile("./My Clippings_test.txt")
	if err != nil && len(highlights) != 2 {
		t.Error("Failed to read highlights from file. ERROR: ", err)
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

func TestRemoveAuthor(t *testing.T) {
	testString := "The story of Microsoft (Bill Gates)"
	testString = removeAuthor(testString)
	if testString != "The story of Microsoft" {
		t.Errorf("Failed to remove author from title %s", testString) //TODO parse another object
	}
}

func TestExtractAuthor(t *testing.T) {
	testString := "The story of Microsoft (Bill Gates)"
	testString = extractAuthor(testString)
	if testString != "Bill Gates" {
		t.Errorf("Failed to extract author's name from %s", testString) //TODO parse another object
	}
}

func TestRemoveAuthorParentheses(t *testing.T) {
	testString := "(Bill Gates)"
	testString = removeAuthorParentheses(testString)
	if testString != "Bill Gates" {
		t.Errorf("Failed to remove parentheses from %s", testString)
	}
}

func TestGetFileContent(t *testing.T) {
	testFile := "My Clippings_test.txt"
	content := getFileContent(testFile)
	if len(content) < 1 {
		t.Errorf("Failed to get any content of from file %s", testFile)
	}
}

func TestIfFileExist(t *testing.T) {
	testFile := "My Clippings_test.txt"
	if !fileExists(testFile) {
		t.Errorf("Failed to detect existence of file %s", testFile)
	}
}
