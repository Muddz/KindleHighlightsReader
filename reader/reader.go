package reader

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

type Highlight struct {
	Title  string
	Author string
	Text   string
}

func ReadHighlightFile(path string) ([]Highlight, error) {
	if !fileExists(path) {
		return nil, errors.New("failed to find file")
	}
	fileContent := getFileContent(path)
	highlights := highlightsParser(fileContent)
	return highlights, nil
}

func highlightsParser(content string) []Highlight {
	pattern := `(?m)(^=+)\r?\n(^.*)\r?\n(^.*)\r?\n\r?\n(^.*)`
	regex := regexp.MustCompile(pattern)
	matches := regex.FindAllStringSubmatch(content, -1)
	var highlights []Highlight
	for _, match := range matches {
		var title = match[2]
		var author = extractAuthor(title)
		var text = match[4]
		highlights = append(highlights, Highlight{
			Title:  title,
			Author: author,
			Text:   text,
		})
	}
	return highlights
}

func extractAuthor(text string) string {
	pattern := `(\(.*\))`
	match, err := regexp.Compile(pattern)
	if err != nil {
		log.Println(err)
	}
	return match.FindString(text)
}

func removeAuthorParentheses(text string) string {
	chars := []rune(text)
	firstChar := string(text[0])
	lastChar := string(text[len(text)-1])

	if firstChar == string('(') || lastChar == string(')') {
		chars = append(chars[1 : len(chars)-1])
	}
	return string(chars)
}

func getFileContent(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println(err)
	}
	return string(data)
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) || info == nil {
		return false
	}
	return !info.IsDir()
}
