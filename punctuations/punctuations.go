package punctuations

import "fmt"

func SetSingleQuotations(input string) string {
	chars := []rune(input)
	firstChar := string(input[0])
	lastChar := string(input[len(input)-1])

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

func SetDoubleQuotations(input string) string {
	chars := []rune(input)
	firstChar := string(input[0])
	lastChar := string(input[len(input)-1])

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

func RemoveQuotations(input string) string {
	chars := []rune(input)
	firstChar := string(input[0])
	lastChar := string(input[len(input)-1])

	if firstChar == string('\'') || firstChar == string('"') {
		chars = append(chars[1:])
	}
	if lastChar == string('\'') || lastChar == string('"') {
		chars = append(chars[:len(chars)-1])
	}
	return string(chars)
}

func SetFullStop(input string) string {
	chars := []rune(input)
	lastChar := string(input[len(input)-1])
	if lastChar != string('.') {
		chars = append(chars, '.')
	}
	return string(chars)
}

func RemoveFullStop(input string) string {
	chars := []rune(input)
	lastChar := string(input[len(input)-1])
	if lastChar == string('.') {
		chars = chars[:len(chars)-1]
	}
	fmt.Println(string(chars))
	return string(chars)
}
