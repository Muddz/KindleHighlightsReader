package clean

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func WhiteSpaces(text string) string {
	fmt.Println(strings.TrimSpace(text))
	return strings.TrimSpace(text)
}

func Prefixes(text string) string {
	pattern := `^\.\s|^\w+\.\s`
	match, err := regexp.Compile(pattern)
	if err != nil {
		log.Println(err)
	}
	return match.ReplaceAllString(text, "")
}

func PostFixes(text string) string {
	pattern := `\.\s\w+$`
	match, err := regexp.Compile(pattern)
	if err != nil {
		log.Println(err)
	}
	return match.ReplaceAllString(text, "")
}
