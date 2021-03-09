package option

import (
	"fmt"
	"regexp"
	"strings"
)

var quotationsMarks = "'\"“”‘’"

func SetQuotations(text string) string {
	chars := []rune(text)
	firstChar := string(chars[0])
	lastChar := string(chars[len(chars)-1])
	quotations := "'\"‘’"

	if strings.ContainsAny(firstChar, quotations) {
		chars[0] = '“'
	} else if firstChar != "“" {
		chars = append(chars, 0)
		copy(chars[1:], chars[0:])
		chars[0] = '“'
	}

	if strings.ContainsAny(lastChar, quotations) {
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
	if strings.ContainsAny(firstChar, quotationsMarks) {
		chars = append(chars[1:])
	}
	if strings.ContainsAny(lastChar, quotationsMarks) {
		chars = append(chars[:len(chars)-1])
	}
	return string(chars)
}

func SetPeriod(text string) string {
	runes := []rune(text)
	lastPos := len(runes)
	lastChar := string(runes[lastPos-1])
	if strings.ContainsAny(lastChar, quotationsMarks) {
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
	lastRunePos := len(runes)
	lastChar := string(runes[lastRunePos-1])
	if strings.ContainsAny(lastChar, quotationsMarks) {
		isPeriod := string(runes[lastRunePos-2])
		if isPeriod == "." {
			content := string(runes[:lastRunePos-2])
			quotation := string(runes[lastRunePos-1])
			return fmt.Sprintf("%v%s", content, quotation)
		}
	} else {
		return strings.TrimRight(text, ".")
	}
	return text
}

func TrimBefore(text string) string {
	pattern := `^\.\s|^\w+\.\s`
	match, _ := regexp.Compile(pattern)
	return match.ReplaceAllString(text, "")
}

func TrimAfter(text string) string {
	pattern := `\s[A-Z]\w+$`
	match, _ := regexp.Compile(pattern)
	return match.ReplaceAllString(text, "")
}

func Capitalize(text string) string {
	pattern := `(^\W*)([a-z])`
	regex, _ := regexp.Compile(pattern)
	firstLetter := regex.FindString(text)
	uppercase := strings.ToUpper(firstLetter)
	result := regex.ReplaceAllString(text, uppercase)
	return result
}
