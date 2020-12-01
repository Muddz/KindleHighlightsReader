package main

import (
	"KindleHighlightsReader/clean"
	"KindleHighlightsReader/export"
	"KindleHighlightsReader/filefinder"
	"KindleHighlightsReader/message"
	"KindleHighlightsReader/punctuations"
	"KindleHighlightsReader/reader"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	optionCleanPrefix = iota + 1
	optionCleanPostFix
)

const (
	optionInsertFullStop = iota + 1
	optionRemoveFullStop
)

const (
	optionSingleQuotation = iota + 1
	optionDoubleQuotation
	optionNoQuotation
)

var ExportFormats = []string{"TEXT", "JSON", "CSV", "PDF"}
var punctuationsOptions = []int{optionSingleQuotation, optionDoubleQuotation, optionNoQuotation}
var fullStopOptions = []int{optionInsertFullStop, optionRemoveFullStop}
var cleaningOptions = []int{optionCleanPrefix, optionCleanPostFix}

var input string
var formats []string
var scanner = bufio.NewScanner(os.Stdin)

func main() {
	fmt.Println(message.GetGreeting())

	src := readSourcePath()
	src = strings.TrimSpace(src)
	highlights := reader.ReadHighlights(src)
	printHighlights(highlights)

	//-------------------------------------------------------------------------------

	cleaningOptions := readCleaning()
	for _, v := range cleaningOptions {
		switch v {
		case optionCleanPrefix:
			for i, v := range highlights {
				highlights[i].Text = clean.Prefixes(v.Text)
			}
			break
		case optionCleanPostFix:
			for i, v := range highlights {
				highlights[i].Text = clean.PostFixes(v.Text)
			}
			break
		}
	}

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

	// -------------------------------------------------------------------------------

	quotationOption := readQuotationMarks()
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

	exportFormats := readExportFormats()
	fmt.Println("Exporting with following formats: ", exportFormats)
	for _, v := range exportFormats {
		switch v {
		case "json":
			export.AsJSON(highlights)
			break
		case "text":
			export.AsTxt(highlights)
			break
		case "pdf":
			export.AsPDF(highlights)
			break
		case "csv":
			export.AsCSV(highlights)
			break
		}
	}

	//-------------------------------------------------------------------------------

	fmt.Println("File saved to your desktop")
	//Press X to exit or R to try again
}

func readSourcePath() string {
	//####
	src := filefinder.GetMyClippingsFile()
	if len(src) > 1 {
		fmt.Printf("Found a 'My Clippings.txt' file at %s\n", src)
	}
	for {
		fmt.Printf("Press C to contiune with that file or X to specify another path")
		if !scanner.Scan() && scanner.Err() != nil {
			fmt.Println("Error:", scanner.Err())
		}

		input = scanner.Text()
		input = strings.TrimSpace(input)
		if strings.EqualFold(input, "c") {
			return src
		} else if strings.EqualFold(input, "x") {
			break
		} else {
			continue
		}
	}
	//####
	fmt.Print(message.EnterSource)
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

func readExportFormats() []string {
	fmt.Print(message.EnterExportFormats)
	for {
		if !scanner.Scan() && scanner.Err() != nil {
			fmt.Println("Error:", scanner.Err())
			return nil
		}
		input := scanner.Text()
		if validateExportFormats(input) {
			formats := strings.Fields(input)
			if len(formats) > 3 {
				formats = formats[:len(ExportFormats)]
				return formats
			} else {
				fmt.Println(input + "\n")
				return formats
			}
		} else {
			fmt.Printf("Error! Invalid export format(s). Try again: ")
		}
	}
}

func validateExportFormats(input string) bool {
	var result = false
	formats = strings.Fields(input)

	if len(formats) < 1 {
		return false
	} else if len(formats) > 3 {
		formats = formats[0:len(ExportFormats)]
	}
	for _, format := range formats {
		for _, validFormat := range ExportFormats {
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
	stats := fmt.Sprintf("### Found %d highlights from %d different books ###\n", highlightsCount, len(booksCount))
	fmt.Println(stats)
}

func readQuotationMarks() int {
	fmt.Println(message.EnterQuotationOption)
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
	fmt.Println(message.EnterFullStopOption)
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

func readCleaning() []int {
	fmt.Println(message.EnterCleanOptions)
	for {
		if !scanner.Scan() && scanner.Err() != nil {
			fmt.Println("Error:", scanner.Err())
		}
		input := scanner.Text()
		input = strings.TrimSpace(input)
		inputs := strings.Fields(input)

		if len(inputs) < 1 || len(inputs) > 3 {
			fmt.Printf("Error! Minimum 1 and Maxiumum 3 options. Try again: ")
			continue
		}

		for _, v := range inputs {
			i, _ := strconv.Atoi(v)
			if i < 1 || i > 3 {
				fmt.Printf("Error!!!")
				continue
			}
		}

		var result []int
		for _, v := range inputs {
			i, _ := strconv.Atoi(v)
			result = append(result, i)
		}

		return result
	}
}
