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
	highlights := highlightsParser(fileContent)
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

func highlightsParser(content string) []Highlight {
	pattern := `(?m)(^=+)\r?\n(^.*)\r?\n(^.*)\r?\n\r?\n(^.*)`
	regex := regexp.MustCompile(pattern)
	matches := regex.FindAllStringSubmatch(content, -1)
	var highlights []Highlight
	for _, match := range matches {

		var title = match[2]
		var text = match[4]

		//TODO make this prettier

		title = removeBOM(title)
		text = removeBOM(text)

		highlights = append(highlights, Highlight{
			Title:  removeAuthorFromTitle(title),
			Author: getAuthorFromTitle(title),
			Text:   cleanText(text),
		})
	}
	return highlights
}

//TODO rename in tests
func removeAuthorFromTitle(title string) string {
	result := strings.Split(title, " (")
	return result[0]
}

//TODO rename in tests
func getAuthorFromTitle(title string) string {
	pattern := `(\(.*\))`
	match, err := regexp.Compile(pattern)
	if err != nil {
		log.Println(err)
	}
	author := match.FindString(title)
	author = removeAuthorParentheses(author)
	return author
}

//TODO rename in tests
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

//TODO needs test
func cleanText(text string) string {
	result := strings.TrimSuffix(text, "\r")
	result = strings.TrimSuffix(result, "\\")
	result = strings.TrimSpace(text)
	return result
}

//TODO needs test
func removeBOM(text string) string {
	bom := "\xef\xbb\xbf"
	b := []byte(text)
	o := []byte(bom)
	b = bytes.ReplaceAll(b, o, nil)
	return string(b)
}
