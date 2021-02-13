package option

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func SetDoubleQuotations(text string) string {
	chars := []rune(text)
	firstChar := string(chars[0])
	lastChar := string(chars[len(chars)-1])

	if firstChar == "'" || firstChar == "\"" {
		chars[0] = '“'
	} else if firstChar != "“" {
		chars = append(chars, 0)
		copy(chars[1:], chars[0:])
		chars[0] = '“'
	}

	if lastChar == "'" || lastChar == "\"" {
		chars[len(chars)-1] = '”'
	} else if lastChar != "”" {
		chars = append(chars, '”')
	}
	return string(chars)
}

func RemoveQuotations(text string) string {
	chars := []rune(text)
	firstChar := string(chars[0])
	lastChar := string(chars[len(chars)-1])

	if firstChar == "'" || firstChar == "\"" || firstChar == "“" { //Make this as an array
		chars = append(chars[1:])
	}
	if lastChar == "'" || lastChar == "\"" || lastChar == "”" { //Make this as an array
		chars = append(chars[:len(chars)-1])
	}
	return string(chars)
}

func SetPeriod(text string) string {
	lastChar := string(text[len(text)-1])
	if lastChar == "'" || lastChar == "\"" || lastChar == "”" { //Make this as an array
		content := text[:len(text)-1]
		quotation := text[len(text)-1:]
		return fmt.Sprintf("%s.%s", content, quotation)
	} else {
		return text + "."
	}
}

func RemovePeriod(text string) string {
	chars := []rune(text)
	lastChar := string(chars[len(chars)-1])
	if lastChar == "'" || lastChar == "\"" || lastChar == "”" { //Make this as an array
		content := text[:len(text)-2]
		quotation := text[len(text)-1:]
		return fmt.Sprintf("%s%s", content, quotation)
	} else {
		return text[:len(text)-1]
	}
}

func TrimBefore(text string) string {
	pattern := `^\.\s|^\w+\.\s`
	match, err := regexp.Compile(pattern)
	if err != nil {
		log.Println(err)
	}
	return match.ReplaceAllString(text, "")
}

func TrimAfter(text string) string {
	pattern := `\s\w+$`
	match, err := regexp.Compile(pattern)
	if err != nil {
		log.Println(err)
	}
	return match.ReplaceAllString(text, "")
}

func Capitalize(text string) string {
	pattern := `(^\W*)([a-z])`
	regex := regexp.MustCompile(pattern)
	firstLetter := regex.FindString(text)
	uppercase := strings.ToUpper(firstLetter)
	result := regex.ReplaceAllString(text, uppercase)
	return result
}
