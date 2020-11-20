package main

import (
	"KindleHighlightsReader/message"
	"KindleHighlightsReader/reader"
	"bufio"
	"fmt"
	"log"
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

var validOutputFormats = []string{"JSON", "PDF", "CSV"}
var options Options

var srcPath string
var dstPath string
var outputFormats []string

var scanner = bufio.NewScanner(os.Stdin)

type Options struct {
	QuotationMarks int
}

func main() {
	//fmt.Println(message.GetGreeting())
	//readSrcPath()

	reader.ReadHighlightFile("reader/My Clippings_test.txt")

	//readDstPath()
	//readOutputFormats()
	//
	//fmt.Println("STARTING...")
	//fmt.Println("\n ")
	//time.Sleep(time.Second * 2)

	//readOptionQuotationMarks()

	//for _, v := range mockValues{
	//	v.Text = punctuations.RemoveQuotations(v.Text)
	//	v.Text = punctuations.SetFullStop(v.Text)
	//	punctuations.RemoveFullStop(v.Text)
	//}

	//printHighlight(mockValues)
	//setSingleQuotations(mockValues)
}

func readSrcPath() {
	fmt.Print(message.SetSrcPath)
	for {
		if !scanner.Scan() && scanner.Err() != nil {
			fmt.Println("Error:", scanner.Err())
			return
		}
		srcPath = scanner.Text()
		if len(srcPath) > 0 {
			fmt.Println(srcPath + "\n")
			return
		} else {
			fmt.Print("Error! Couldn't find text file. Try again: ")
		}
	}
}

func readOutputFormats() {
	fmt.Print(message.SetOutputFormats)
	for {
		if !scanner.Scan() && scanner.Err() != nil {
			fmt.Println("Error:", scanner.Err())
			return
		}
		formats := scanner.Text()
		if len(formats) > 0 && validateOutputFormats(formats) {

			//TODO make this prettier
			if len(strings.Fields(formats)) > 3 {
				fmt.Println(strings.Fields(formats)[:len(validOutputFormats)])
			} else {
				fmt.Println(formats + "\n")
			}

			return
		} else {
			fmt.Printf("Error! Invalid formats. Try again: ")
		}
	}
}

func validateOutputFormats(formats string) bool {
	var result = false
	outputFormats = strings.Fields(formats)
	if len(outputFormats) > 3 {
		outputFormats = outputFormats[0:len(validOutputFormats)]
	}
	for _, format := range outputFormats {
		for _, validFormat := range validOutputFormats {
			result = strings.EqualFold(format, validFormat)
			if result {
				break
			}
		}
	}
	return result
}

func readDstPath() {
	fmt.Print(message.SetDstPath)
	for {
		if !scanner.Scan() && scanner.Err() != nil {
			fmt.Println("Error:", scanner.Err())
			return
		}
		dstPath = scanner.Text()
		if len(dstPath) == 0 {
			dstPath = getUserDesktopDst()
		}
		if destinationExists(dstPath) {
			fmt.Println(dstPath + "\n")
			return
		} else {
			fmt.Print("Error! Couldn't find destination path. Try again: ")
		}
	}
}

func destinationExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func getUserDesktopDst() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Println(err)
	}
	return fmt.Sprintf("%s\\Desktop", homeDir)
}

//func printHighlight(highlights []Highlight) {
//	highlightsCount := 0
//	booksCount := make(map[string]int)
//	for _, v := range highlights {
//		msg := fmt.Sprintf("Title: %s\nAuthor: %s %s\nText: %s\n", v.Title, v.Author, v.AuthorLastName, v.Text)
//		fmt.Println(msg)
//		if len(v.Text) > 0 {
//			highlightsCount++
//		}
//		if len(v.Title) > 0 {
//			booksCount[v.Title] = 0
//		}
//	}
//	msg := fmt.Sprintf("### Found %d highlights from %d different books", highlightsCount, len(booksCount))
//	fmt.Println(msg)
//}

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
