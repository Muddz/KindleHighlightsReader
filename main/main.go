package main

import (
	"KindleHighlightsReader/message"
	"KindleHighlightsReader/reader"
	"KindleHighlightsReader/save"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	optionSingleQuotation = iota
	optionDoubleQuotation
	optionNoQuotation
	optionSkipQuotationS
)

var validOutputFormats = []string{"JSON", "PDF", "CSV", "TEXT"}
var options Options
var input string
var formats []string
var scanner = bufio.NewScanner(os.Stdin)

type Options struct {
	QuotationMarks int
}

func main() {
	//fmt.Println(message.GetGreeting())
	//readSrcPath()
	highlights, _ := reader.ReadHighlightFile("C:\\Users\\Muddz\\Desktop\\My Clippings.txt")
	//printHighlight(highlights)
	//if err != nil {
	//	log.Println(err)
	//}

	//save.ToJSON(highlights)
	//save.ToCSV(highlights)
	//save.ToTxt(highlights)
	//save.ToPDF(highlights)

	//save.ToJSON(highlights, getUserDesktopPath())
	//
	//save.ToTxt(highlights, getUserDesktopPath())

	//if len(highlights) > 0 {
	//	printHighlight(highlights)
	//}

	//readDstPath()
	//readOutputFormats()
	//
	//fmt.Println("STARTING...")
	//fmt.Println("\n ")
	//time.Sleep(time.Second * 2)

	//readOptionQuotationMarks()
	//setSingleQuotations(mockValues)

}

func readSrcPath() {
	fmt.Print(message.SetSrcPath)
	for {
		if !scanner.Scan() && scanner.Err() != nil {
			fmt.Println("Error:", scanner.Err())
			return
		}
		input = scanner.Text()
		if fileExist(input) {
			fmt.Println(input + "\n")
			return
		} else {
			fmt.Print("Error! Couldn't find text file. Try again: ")
		}
	}
}

func fileExist(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func readOutputFormats() {
	fmt.Print(message.SetOutputFormats)
	for {
		if !scanner.Scan() && scanner.Err() != nil {
			fmt.Println("Error:", scanner.Err())
			return
		}
		input := scanner.Text()
		if validateOutputFormats(input) {
			formats := strings.Fields(input)
			if len(formats) > 3 {
				formats = formats[:len(validOutputFormats)]
				fmt.Println(formats)
			} else {
				fmt.Println(input + "\n")
			}
			return
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
func printHighlight(highlights []reader.Highlight) {
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

func readOptionQuotationMarks() {
	fmt.Println(message.OptionQuotationMarks)
	for {
		if !scanner.Scan() && scanner.Err() != nil {
			fmt.Println("Error:", scanner.Err())
			return
		}
		input := scanner.Text()
		i, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
		}
		if i > 4 && i < 1 {
			fmt.Printf("Error! Choose only one input between 1-4. Try again: ")
		} else {
			options.QuotationMarks = i
			fmt.Println(i)
			return
		}
	}
}

func setQuotationMarksOption(option int) {
	switch option {

	case 1:

		break

	case 2:

		break

	case 3:

		break

	case 4:
		return

	}
}
