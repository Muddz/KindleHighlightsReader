package reader

import (
	"KindleHighlightsReader/highlight"
	"bytes"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func ReadHighlights(path string) ([]highlight.Highlight, error) {
	c, err := getFileContent(path)
	if err != nil {
		return nil, fmt.Errorf("\n%w", err)
	}
	h, err := parseHighlights(c)
	if err != nil {
		return nil, fmt.Errorf("\n%w", err)
	}
	return h, nil
}

func getFileContent(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("\n%w", err)
	}
	return string(b), nil
}

func parseHighlights(content string) ([]highlight.Highlight, error) {
	p := `(?m)(^.*)\((.*)\)\r?\n(^.*)\r?\n\r?\n(^.*)`
	r, err := regexp.Compile(p)
	if err != nil {
		return nil, fmt.Errorf("\n%w", err)
	}
	matches := r.FindAllStringSubmatch(content, -1)
	var highlights []highlight.Highlight
	for _, matchGroups := range matches {
		var title = matchGroups[1]
		var author = matchGroups[2]
		var text = matchGroups[4]

		title = removeBOM(title)
		text = removeBOM(text)
		title = strings.TrimSpace(title)
		text = strings.TrimRight(text, "\r")

		h := highlight.New(title, author, text)
		highlights = append(highlights, h)
	}
	return highlights, nil
}

func removeBOM(text string) string {
	bom := "\xef\xbb\xbf"
	b := []byte(text)
	o := []byte(bom)
	b = bytes.ReplaceAll(b, o, nil)
	return string(b)
}
