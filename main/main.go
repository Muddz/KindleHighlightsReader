package main

import (
	"KindleHighlightsReader/message"
	"KindleHighlightsReader/punctuations"
	"KindleHighlightsReader/reader"
	"KindleHighlightsReader/save"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	optionSingleQuotation = iota + 1
	optionDoubleQuotation
	optionNoQuotation
)

const (
	optionInsertFullStop = iota + 1
	optionRemoveFullStop
)

var validOutputFormats = []string{"TEXT", "JSON", "CSV", "PDF"}
var input string
var formats []string
var scanner = bufio.NewScanner(os.Stdin)

func main() {

	fmt.Println(message.GetGreeting())
	src := readSrcPath()
	src = strings.TrimSpace(src)
	highlights := reader.ReadHighlightFile(src)
	printHighlights(highlights)

	//-------------------------------------------------------------------------------

	fullStopOption := readFullStop()
	switch fullStopOption {
	case optionInsertFullStop:
		for i, v := range highlights {
			highlights[i].Text = punctuations.SetFullStop(v.Text)
		}
		break
	case optionRemoveFullStop:
		for i, v := range highlights {
			highlights[i].Text = punctuations.RemoveFullStop(v.Text)
		}
		break
	}

	//-------------------------------------------------------------------------------

	quotationOption := readOptionQuotationMarks()
	switch quotationOption {
	case optionSingleQuotation:
		for _, v := range highlights {
			v.Text = punctuations.SetSingleQuotations(v.Text)
		}
		break
	case optionDoubleQuotation:
		for _, v := range highlights {
			v.Text = punctuations.SetDoubleQuotations(v.Text)
		}
		break
	case optionNoQuotation:
		for _, v := range highlights {
			v.Text = punctuations.RemoveQuotations(v.Text)
		}
		break
	}

	//-------------------------------------------------------------------------------

	outputFormats := readOutputFormats()
	fmt.Println("Saving to the following formats: ", outputFormats)
	for _, v := range outputFormats {
		switch v {
		case "json":
			save.ToJSON(highlights)
			break
		case "text":
			save.ToTxt(highlights)
			break
		case "pdf":
			save.ToPDF(highlights)
		case "csv":
			save.ToCSV(highlights)
			break
		}
	}

	//-------------------------------------------------------------------------------

	fmt.Println("File saved to your desktop")
}

func readSrcPath() string {
	fmt.Print(message.SetSrcPath)
	for {
		if !scanner.Scan() && scanner.Err() != nil {
			fmt.Println("Error:", scanner.Err())
			return ""
		}
		input = scanner.Text()
		input = strings.TrimSpace(input)
		if fileExist(input) {
			fmt.Println(input + "\n")
			return input
		} else {
			fmt.Print("Error! Couldn't find text file. Try again: ")
		}
	}
}

func fileExist(path string) bool {
	info, err := os.Stat(path)
	if info == nil || os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func readOutputFormats() []string {
	fmt.Print(message.SetOutputFormats)
	for {
		if !scanner.Scan() && scanner.Err() != nil {
			fmt.Println("Error:", scanner.Err())
			return nil
		}
		input := scanner.Text()
		if validateOutputFormats(input) {
			formats := strings.Fields(input)
			if len(formats) > 3 {
				formats = formats[:len(validOutputFormats)]
				return formats
			} else {
				fmt.Println(input + "\n")
				return formats
			}
		} else {
			fmt.Printf("Error! Invalid output format(s). Try again: ")
		}
	}
}

func validateOutputFormats(input string) bool {
	var result = false
	formats = strings.Fields(input)

	if len(formats) < 1 {
		return false
	} else if len(formats) > 3 {
		formats = formats[0:len(validOutputFormats)]
	}
	for _, format := range formats {
		for _, validFormat := range validOutputFormats {
			result = strings.EqualFold(format, validFormat)
			if result {
				break
			}
		}
	}
	return result
}

//TODO is the argument passing correct?
func printHighlights(highlights []reader.Highlight) {
	highlightsCount := 0
	booksCount := make(map[string]int)
	for _, v := range highlights {
		msg := fmt.Sprintf("Title: %s\nAuthor: %s\nText: %s\n", v.Title, v.Author, v.Text)
		fmt.Println(msg)
		if len(v.Text) > 0 {
			highlightsCount++
		}
		if len(v.Title) > 0 {
			booksCount[v.Title] = 0
		}
	}
	msg := fmt.Sprintf("### Found %d highlights from %d different books ###", highlightsCount, len(booksCount))
	fmt.Println(msg)
}

func readOptionQuotationMarks() int {
	fmt.Println(message.OptionQuotationMarks)
	for {
		if !scanner.Scan() && scanner.Err() != nil {
			fmt.Println("Error:", scanner.Err())
			return -1
		}
		input := scanner.Text()
		i, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
		}
		if i > 4 && i < 1 {
			fmt.Printf("Error! Choose only one input between 1-4. Try again: ")
		} else {
			return i
		}
	}
}

func readFullStop() int {
	fmt.Println(message.OptionFullstops)
	for {
		if !scanner.Scan() && scanner.Err() != nil {
			fmt.Println("Error:", scanner.Err())
			return -1
		}
		input := scanner.Text()
		i, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
		}
		if i > 3 && i < 1 {
			fmt.Printf("Error! Choose only one input between 1-3. Try again: ")
		} else {
			return i
		}
	}
}
