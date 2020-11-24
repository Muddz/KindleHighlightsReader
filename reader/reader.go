package reader

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

type Highlight struct {
	Text   string
	Author string
	Title  string
}

func ReadHighlightFile(path string) ([]Highlight, error) {
	if !fileExists(path) {
		return nil, errors.New("failed to find file")
	}
	fileContent := getFileContent(path)
	highlights := highlightsParser(fileContent)
	return highlights, nil
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) || info == nil {
		return false
	}
	return !info.IsDir()
}

func getFileContent(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println(err)
	}
	content := fmt.Sprintf("=\n%s", string(data))
	return content
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
	author := match.FindString(title)
	author = removeAuthorParentheses(author)
	return author
}

func removeAuthorParentheses(author string) string {
	if len(author) < 1 {
		return ""
	}
	chars := []rune(author)
	firstChar := string(author[0])
	lastChar := string(author[len(author)-1])

	if firstChar == string('(') || lastChar == string(')') {
		chars = append(chars[1 : len(chars)-1])
	}
	return string(chars)
}
