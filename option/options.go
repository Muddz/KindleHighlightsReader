package option

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func SetQuotations(text string) string {
	chars := []rune(text)
	firstChar := string(chars[0])
	lastChar := string(chars[len(chars)-1])

	if firstChar == "'" || firstChar == "\"" || firstChar == "‘" { //Make this as an array
		chars[0] = '“'
	} else if firstChar != "“" {
		chars = append(chars, 0)
		copy(chars[1:], chars[0:])
		chars[0] = '“'
	}

	if lastChar == "'" || lastChar == "\"" || lastChar == "’" { //Make this as an array
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

	if firstChar == "'" || firstChar == "\"" || firstChar == "“" || firstChar == "‘" { //Make this as an array
		chars = append(chars[1:])
	}
	if lastChar == "'" || lastChar == "\"" || lastChar == "”" || lastChar == "’" { //Make this as an array
		chars = append(chars[:len(chars)-1])
	}
	return string(chars)
}

func SetPeriod(text string) string {
	runes := []rune(text)
	lastPos := len(runes)
	lastChar := string(runes[lastPos-1])

	if lastChar == "'" || lastChar == "\"" || lastChar == "”" || lastChar == "’" { //Make this as an array
		isPeriod := string(runes[lastPos-2])
		if isPeriod != "." {
			text := string(runes[:lastPos-1])
			quotation := string(runes[lastPos-1])
			return fmt.Sprintf("%s.%s", text, quotation)
		}
	} else if lastChar != "." {
		return text + "."
	}
	return text
}

func RemovePeriod(text string) string {
	runes := []rune(text)
	lastPos := len(runes)
	lastChar := string(runes[lastPos-1])
	if lastChar == "'" || lastChar == "\"" || lastChar == "”" || lastChar == "’" { //Make this as an array
		content := string(runes[:lastPos-2])
		quotation := string(runes[lastPos-1])
		return fmt.Sprintf("%v%s", content, quotation)
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
