package main

import (
	"KindleHighlightsReader/clean"
	"KindleHighlightsReader/export"
	"KindleHighlightsReader/finder"
	"KindleHighlightsReader/message"
	"KindleHighlightsReader/punctuations"
	"KindleHighlightsReader/reader"
	"bufio"
	"fmt"
	"log"
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

var exportFormats = []string{"TEXT", "JSON", "CSV", "PDF"}
var punctuationsOptions = []int{optionSingleQuotation, optionDoubleQuotation, optionNoQuotation}
var fullStopOptions = []int{optionInsertFullStop, optionRemoveFullStop}
var cleaningOptions = []int{optionCleanPrefix, optionCleanPostFix}
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
	var exportPaths []string
	for _, v := range exportFormats {
		switch v {
		case "json":
			path, err := export.AsJSON(highlights)
			if err != nil {
				log.Print(err)
				break
			}
			exportPaths = append(exportPaths, path)
		case "text":
			path, err := export.AsTxt(highlights)
			if err != nil {
				log.Print(err)
				break
			}
			exportPaths = append(exportPaths, path)
		case "pdf":
			path, err := export.AsPDF(highlights)
			if err != nil {
				log.Print(err)
				break
			}
			exportPaths = append(exportPaths, path)
		case "csv":
			path, err := export.AsCSV(highlights)
			if err != nil {
				log.Print(err)
				break
			}
			exportPaths = append(exportPaths, path)
		}
	}

	for _, v := range exportPaths {
		fmt.Println("Exported to:", v)
	}

}

func readSourcePath() string {

	//Auto scan
	src := finder.GetMyClippingsFile()
	if len(src) > 1 {
		fmt.Printf("Found a 'My Clippings.txt' file at %s\n", src)
		for {
			fmt.Printf("Enter C to continue with that file or X to specify another path: ")
			input := scanInput()
			if strings.EqualFold(input, "c") {
				return src
			}
			if strings.EqualFold(input, "x") {
				break
			}
		}
	}

	//Manuel Scan
	fmt.Print(message.EnterSource)
	for {
		input := scanInput()
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
	fmt.Println(message.EnterExportFormats)
	for {
		input := scanInput()
		if validateExportFormats(input) {
			formats := strings.Fields(input)
			if len(formats) > 3 {
				formats = formats[:len(exportFormats)]
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
	formats := strings.Fields(input)
	if len(formats) < 1 {
		return false
	} else if len(formats) > 3 {
		formats = formats[0:len(exportFormats)]
	}
	for _, format := range formats {
		for _, validFormat := range exportFormats {
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
		input := scanInput()
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
		input := scanInput()
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
		input := scanInput()
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

func scanInput() string {
	if !scanner.Scan() && scanner.Err() != nil {
		fmt.Println("ERROR! Failed to read input:", scanner.Err())
		return ""
	}
	input := scanner.Text()
	return strings.TrimSpace(input)
}
