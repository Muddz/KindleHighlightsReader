package main

import (
	"KindleHighlightsReader/export"
	"KindleHighlightsReader/filefinder"
	"KindleHighlightsReader/highlight"
	"KindleHighlightsReader/message"
	"KindleHighlightsReader/option"
	"KindleHighlightsReader/reader"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	optionTrimBefore = iota + 1
	optionTrimAfter
	optionTrimSkip
)

const (
	optionSetPeriod = iota + 1
	optionRemovePeriod
	optionSkipPeriod
)

const (
	optionDoubleQuotations = iota + 1
	optionRemoveQuotations
	optionSkipQuotations
)

const (
	optionCapitalize = iota + 1
	optionSkipCapitalize
)

var validExportOptions = []string{"TEXT", "JSON", "CSV", "PDF"}
var validQuotationsOptions = []int{optionDoubleQuotations, optionRemoveQuotations, optionSkipQuotations}
var validPeriodOptions = []int{optionSetPeriod, optionRemovePeriod, optionSkipPeriod}
var validTrimOptions = []int{optionTrimBefore, optionTrimAfter, optionTrimSkip}
var validCapitalizationOptions = []int{optionCapitalize, optionSkipCapitalize}
var scanner = bufio.NewScanner(os.Stdin)

func main() {
	fmt.Println(message.GetGreeting())

	src := findSource()
	if len(src) == 0 {
		src = readSource()
	}

	highlights := reader.ReadHighlights(src)
	if len(highlights) > 0 {
		printHighlights(highlights)
	}

	//-------------------------------------------------------------------------------

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

	quotationOption := readQuotationOption()
	switch quotationOption {
	case optionDoubleQuotations:
		for i, v := range highlights {
			highlights[i].Text = option.SetQuotations(v.Text)
		}
		break
	case optionRemoveQuotations:
		for i, v := range highlights {
			highlights[i].Text = option.RemoveQuotations(v.Text)
		}
		break
	}

	// -------------------------------------------------------------------------------

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

	//-------------------------------------------------------------------------------

	capitalizeOption := readCapitalizeOption()
	switch capitalizeOption {
	case optionCapitalize:
		for i, v := range highlights {
			highlights[i].Text = option.Capitalize(v.Text)
		}
		break
	}

	//-------------------------------------------------------------------------------

	exportOptions := readExportOptions()
	var exportResults []string
	for k, _ := range exportOptions {
		switch k {
		case "text", "TEXT":
			path, err := export.AsTxt(highlights)
			if err != nil {
				log.Print(err)
				break
			}
			exportResults = append(exportResults, path)
		case "json", "JSON":
			path, err := export.AsJSON(highlights)
			if err != nil {
				log.Print(err)
				break
			}
			exportResults = append(exportResults, path)
		case "csv", "CSV":
			path, err := export.AsCSV(highlights)
			if err != nil {
				log.Print(err)
				break
			}
			exportResults = append(exportResults, path)
		case "pdf", "PDF":
			path, err := export.AsPDF(highlights)
			if err != nil {
				log.Print(err)
				break
			}
			exportResults = append(exportResults, path)
		}
	}

	fmt.Println(" ")
	for _, v := range exportResults {
		fmt.Println("Exported to:", v)
	}

	time.Sleep(time.Second * 15)
}

func findSource() string {
	src := filefinder.GetMyClippingsFile()
	var input string
	if len(src) > 0 {
		fmt.Printf("Found a 'My Clippings.txt' file at %s\n", src)
		fmt.Printf("Press ENTER to continue with that file or specify another path:")

		for {
			input = scanInput()
			input = trimSrc(input)
			if len(input) == 0 {
				return src
			}
			break
		}

		for {
			if fileExist(input) {
				return input
			} else {
				fmt.Printf("Couldn't find file: %s. Try again: ", input)
				input = scanInput()
				input = trimSrc(input)
			}
		}
	}
	return ""
}

func readSource() string {
	fmt.Println(message.EnterSource)
	for {
		input := scanInput()
		input = trimSrc(input)
		if fileExist(input) {
			return input
		} else {
			fmt.Printf("Couldn't find file: %s. Try again: ", input)
		}
	}
}

func trimSrc(src string) string {
	if strings.ContainsAny(src, "\"") {
		src = strings.ReplaceAll(src, "\"", "")
	}
	return strings.TrimSpace(src)
}

func fileExist(path string) bool {
	info, err := os.Stat(path)
	if info == nil || os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func readTrimOptions() []int {
	fmt.Println("")
	fmt.Println(message.EnterTrimOptions)
OUTER:
	for {
		input := scanInput()
		inputs := strings.Fields(input)

		if len(inputs) < 1 || len(inputs) > len(validTrimOptions) {
			fmt.Printf("Error! Choose 1-%d options. Try again: ", len(validTrimOptions))
			continue
		}

		for _, v := range inputs {
			n, _ := strconv.Atoi(v)
			if n < 1 || n > 3 {
				fmt.Printf("Error! Option not valid. Try again: ")
				continue OUTER
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

func readPeriodOption() int {
	fmt.Println("")
	fmt.Println(message.EnterPeriodOption)
	for {
		input := scanInput()
		n, _ := strconv.Atoi(input)
		if n < 1 || n > len(validPeriodOptions) {
			fmt.Print("Error! Choose one option between 1-3. Try again: ")
			continue
		}
		return n
	}
}

func readQuotationOption() int {
	fmt.Println("")
	fmt.Println(message.EnterQuotationOption)
	for {
		input := scanInput()
		i, _ := strconv.Atoi(input)
		if i < 1 || i > len(validQuotationsOptions) {
			fmt.Printf("Error! Choose one option between 1-4. Try again: ")
			continue
		}
		return i
	}
}

func readCapitalizeOption() int {
	fmt.Println("")
	fmt.Println(message.EnterCapitalizationOption)
	for {
		input := scanInput()
		i, _ := strconv.Atoi(input)
		if i < 1 || i > len(validCapitalizationOptions) {
			fmt.Printf("Error! Choose one option between the valid options. Try again: ")
			continue
		}
		return i
	}
}

func readExportOptions() map[string]string {
	fmt.Println("")
	fmt.Println(message.EnterExportOptions)
	for {
		input := scanInput()
		inputs := strings.Fields(input)
		formats := make(map[string]string)
		for _, v := range inputs {
			formats[v] = v
		}
		if len(formats) < 1 || len(formats) > len(validExportOptions) {
			fmt.Printf("Error! Choose between 1 to 4 options. Try again: ")
			continue
		}
		if !isExportFormatsValid(formats) {
			fmt.Printf("Error! Invalid export inputs. Try again: ")
			continue
		}
		return formats
	}
}

func isExportFormatsValid(formats map[string]string) bool {
	var result = false
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

func printHighlights(highlights []highlight.Highlight) {
	highlightsCount := 0
	booksCount := make(map[string]int)
	for _, v := range highlights {
		msg := fmt.Sprintf("\nTitle: %s\nAuthor: %s\nText: %s\n", v.Title, v.Author, v.Text)
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

func scanInput() string {
	if !scanner.Scan() && scanner.Err() != nil {
		fmt.Println("ERROR! Failed to read input:", scanner.Err())
		return ""
	}
	input := scanner.Text()
	return strings.TrimSpace(input)
}
