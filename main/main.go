package main

import (
	"KindleHighlightsReader/export"
	"KindleHighlightsReader/filefinder"
	"KindleHighlightsReader/message"
	"KindleHighlightsReader/option"
	"KindleHighlightsReader/reader"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	optionTrimBefore = iota + 1
	optionTrimAfter
)

const (
	optionSetPeriod = iota + 1
	optionRemovePeriod
)

const (
	optionSingleQuotation = iota + 1
	optionDoubleQuotation
	optionNoQuotation
)

var validExportOptions = []string{"TEXT", "JSON", "CSV", "PDF"}
var validQuotations = []int{optionSingleQuotation, optionDoubleQuotation, optionNoQuotation}
var validPeriod = []int{optionSetPeriod, optionRemovePeriod}
var validTrim = []int{optionTrimBefore, optionTrimAfter}
var scanner = bufio.NewScanner(os.Stdin)

func main() {
	fmt.Println(message.GetGreeting())
	src := readSourcePath()
	src = strings.TrimSpace(src)
	highlights := reader.ReadHighlights(src)
	printHighlights(highlights)

	//-------------------------------------------------------------------------------

	//TODO fix the error message
	trimOptions := readTrimOptions()
	for _, v := range trimOptions {
		if v == optionTrimBefore {
			for i, v := range highlights {
				highlights[i].Text = option.TrimBefore(v.Text)
			}
		}
		if v == optionTrimAfter {
			for i, v := range highlights {
				highlights[i].Text = option.TrimAfter(v.Text)
			}
		}
	}

	//-------------------------------------------------------------------------------

	periodOption := readPeriodOption()
	switch periodOption {
	case optionSetPeriod:
		for i, v := range highlights {
			highlights[i].Text = option.SetPeriod(v.Text)
		}
		break
	case optionRemovePeriod:
		for i, v := range highlights {
			highlights[i].Text = option.RemovePeriod(v.Text)
		}
		break
	}

	// -------------------------------------------------------------------------------

	quotationOption := readQuotationOption()
	switch quotationOption {
	case optionSingleQuotation:
		for _, v := range highlights {
			v.Text = option.SetSingleQuotations(v.Text)
		}
		break
	case optionDoubleQuotation:
		for _, v := range highlights {
			v.Text = option.SetDoubleQuotations(v.Text)
		}
		break
	case optionNoQuotation:
		for _, v := range highlights {
			v.Text = option.RemoveQuotations(v.Text)
		}
		break
	}

	//-------------------------------------------------------------------------------

	exportOptions := readExportOptions()
	var exportPaths []string
	for _, v := range exportOptions {
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

	//TODO) Auto scan
	src := filefinder.GetMyClippingsFile()
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

	//TODO) Manuel Scan
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

func readExportOptions() []string {
	fmt.Println(message.EnterExportOptions)
	for {
		input := scanInput()
		if validateExportFormats(input) {
			formats := strings.Fields(input)
			if len(formats) > len(validExportOptions) {
				formats = formats[:len(validExportOptions)]
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
		formats = formats[0:len(validExportOptions)]
	}
	for _, format := range formats {
		for _, validFormat := range validExportOptions {
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

func readQuotationOption() int {
	fmt.Println(message.EnterQuotationOption)
	for {
		input := scanInput()
		i, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
		}
		if i > len(validQuotations) && i < 1 {
			fmt.Printf("Error! Choose only one input between 1-4. Try again: ")
		} else {
			return i
		}
	}
}

func readPeriodOption() int {
	fmt.Println(message.EnterPeriodOption)
	for {
		input := scanInput()
		i, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
		}
		if i > len(validPeriod) && i < 1 {
			fmt.Printf("Error! Choose only one input between 1-3. Try again: ")
		} else {
			return i
		}
	}
}

func readTrimOptions() []int {
	fmt.Println(message.EnterTrimOptions)
	for {
		input := scanInput()
		inputs := strings.Fields(input)
		if len(inputs) < 1 || len(inputs) > len(validTrim) {
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
