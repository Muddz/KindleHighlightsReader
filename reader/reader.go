package reader

import (
	"bytes"
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
	return string(data)
}

func parseHighlights(content string) []Highlight {
	pattern := `(?m)(^.*)\((.*)\)\r?\n(^.*)\r?\n\r?\n(^.*)`
	regex := regexp.MustCompile(pattern)
	matches := regex.FindAllStringSubmatch(content, -1)
	var highlights []Highlight
	for _, matchGroups := range matches {
		var title = matchGroups[1]
		var author = matchGroups[2]
		var text = matchGroups[4]

		title = removeBOM(title)
		text = removeBOM(text)
		text = removeCarriageReturn(text)

		h := Highlight{Text: title, Author: author, Title: text}
		highlights = append(highlights, h)
	}
	return highlights
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
