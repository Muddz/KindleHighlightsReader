package reader

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
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
		var text = match[4]
		highlights = append(highlights, Highlight{
			Title:  removeAuthor(title),
			Author: extractAuthor(title),
			Text:   text,
		})
	}
	return highlights
}

func removeAuthor(title string) string {
	result := strings.Split(title, " (")
	return result[0]
}

func extractAuthor(title string) string {
	pattern := `(\(.*\))`
	match, err := regexp.Compile(pattern)
	if err != nil {
		log.Println(err)
	}
	title = match.FindString(title)
	return removeAuthorParentheses(title)
}

func removeAuthorParentheses(author string) string {
	chars := []rune(author)
	firstChar := string(author[0])
	lastChar := string(author[len(author)-1])

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
