package reader

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

type Highlight struct {
	Text   string
	Author string
	Title  string
}

func ReadHighlights(path string) []Highlight {
	fileContent := getFileContent(path)
	highlights := parseHighlights(fileContent)
	return highlights
}

func getFileContent(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println(err)
	}
	content := fmt.Sprintf("=\n%s", string(data))
	return content
}

func parseHighlights(content string) []Highlight {
	pattern := `(?m)(^.*)(\(.*\))\r?\n(^.*)\r?\n\r?\n(^.*)`
	regex := regexp.MustCompile(pattern)
	matches := regex.FindAllStringSubmatch(content, -1)
	var highlights []Highlight
	for _, matchGroups := range matches {

		var title = matchGroups[1]
		var author = matchGroups[2]
		var text = matchGroups[4]

		title = removeBOM(title)
		author = removeParentheses(author)
		text = removeBOM(text)
		text = removeCarriageReturn(text)

		highlights = append(highlights, Highlight{
			Title:  title,
			Author: author,
			Text:   text,
		})
	}
	return highlights
}

func removeParentheses(author string) string {
	if len(author) < 1 {
		return ""
	}
	hasStartParenthesis := strings.Contains(author, "(")
	hasEndParenthesis := strings.Contains(author, ")")
	if hasStartParenthesis && hasEndParenthesis {
		author = strings.ReplaceAll(author, "(", "")
		author = strings.ReplaceAll(author, ")", "")
	}

	return author
}

func removeCarriageReturn(text string) string {
	result := strings.TrimSuffix(text, "\r")
	return result
}

func removeBOM(text string) string {
	bom := "\xef\xbb\xbf"
	b := []byte(text)
	o := []byte(bom)
	b = bytes.ReplaceAll(b, o, nil)
	return string(b)
}
