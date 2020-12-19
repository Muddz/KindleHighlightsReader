package option

import (
	"log"
	"regexp"
)

func SetSingleQuotations(text string) string {
	chars := []rune(text)
	firstChar := string(text[0])
	lastChar := string(text[len(text)-1])

	if firstChar == string('"') {
		chars[0] = '\''
	} else if firstChar != string('\'') {
		chars = append(chars, 0)
		copy(chars[1:], chars[0:])
		chars[0] = '\''
	}

	if lastChar == string('"') {
		chars[len(chars)-1] = '\''
	} else if lastChar != string('\'') {
		chars = append(chars, '\'')
	}
	return string(chars)
}

func SetDoubleQuotations(text string) string {
	chars := []rune(text)
	firstChar := string(text[0])
	lastChar := string(text[len(text)-1])

	if firstChar == string('\'') {
		chars[0] = '"'
	} else if firstChar != string('"') {
		chars = append(chars, 0)
		copy(chars[1:], chars[0:])
		chars[0] = '"'
	}

	if lastChar == string('\'') {
		chars[len(chars)-1] = '"'
	} else if lastChar != string('"') {
		chars = append(chars, '"')
	}
	return string(chars)
}

func RemoveQuotations(text string) string {
	chars := []rune(text)
	firstChar := string(text[0])
	lastChar := string(text[len(text)-1])

	if firstChar == string('\'') || firstChar == string('"') {
		chars = append(chars[1:])
	}
	if lastChar == string('\'') || lastChar == string('"') {
		chars = append(chars[:len(chars)-1])
	}
	return string(chars)
}

func SetPeriod(text string) string {
	lastRune := text[len(text)-1]
	lastChar := string(lastRune)
	if lastChar != string('.') {
		return text + "."
	}
	return text
}

func RemovePeriod(text string) string {
	chars := []rune(text)
	lastChar := string(text[len(text)-1])
	if lastChar == string('.') {
		chars = chars[:len(chars)-1]
	}
	return string(chars)
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
	pattern := `\.\s\w+$`
	match, err := regexp.Compile(pattern)
	if err != nil {
		log.Println(err)
	}
	return match.ReplaceAllString(text, "")
}
