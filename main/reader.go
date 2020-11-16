package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var mockValues []Highlight

const programVersion = "v1.0.0"
const msgWelcome = `Welcome! KindleHighlightsReader %s is a program that reads all of your Kindle highlights 
and can save them nicely formatted in either Json, PDF or as an CSV file.`

const (
	msgSetSrcPath        = `Enter the path for the "My Clippings.txt" file: `
	msgSetOutputFormats  = `Enter one or more of the following output formats JSON, PDF, CSV separated by spaces: `
	msgSetDstPath        = `Enter a destination path for the output file(s) or leave empty for desktop path: `
	msgSetQuotationMarks = `Wrap every highlight in quotation marks?
1 - YES (Double quotation "Hi")
2 - YES (Single quotation 'Hi')
3 - REMOVE
4 - SKIP`
)

const (
	optionSingleQuotation = iota
	optionDoubleQuotation
	optionNoQuotation
	optionSkipQuotationS
)

var highlights []Highlight
var validOutputFormats = []string{"JSON", "PDF", "CSV"}
var options Options

var srcPath string
var dstPath string
var outputFormats []string

var scanner = bufio.NewScanner(os.Stdin)

type Options struct {
	QuotationMarks int
}

type Highlight struct {
	Title           string
	AuthorFirstName string
	AuthorLastName  string
	Text            string
}

func main() {
	//TODO should methods return a value and append it to the global member variables?
	//greetMsg()
	//readSrcPath()
	//readDstPath()
	//readOutputFormats()
	//
	//fmt.Println("STARTING...")
	//fmt.Println("\n ")
	//time.Sleep(time.Second * 2)

	//readOptionQuotationMarks()

	mockValues = append(mockValues,

		Highlight{
			Title:           "The Book of Pook",
			AuthorFirstName: "Pook",
			AuthorLastName:  "",
			Text:            "'The greatest risk you can take in life is not to risk it all!'"},

		Highlight{
			Title:           "The Book of Pook",
			AuthorFirstName: "Pook",
			AuthorLastName:  "",
			Text:            "'You can be the smartest person in the world, the most talented, the most persistent, but you will never win in the world or with women unless you embrace the glory of RISK."},

		Highlight{
			Title:           "'The Book of Pook",
			AuthorFirstName: "Pook",
			AuthorLastName:  "",
			Text:            "If you start treating a woman like precious gold, she will believe she is gold. And once she believes it, she will DUMP YOU because YOU have given her the sense that she is BETTER then you.'"},

		Highlight{
			Title:           "The Manual: What Women Want and How to Give It to Them",
			AuthorFirstName: "W.",
			AuthorLastName:  "Anton",
			Text:            "\"To spend a lot of money on women that you have not had sex with is also a bad idea for a range of other reasons. First, you risk making a woman feel uncomfortable, either by making her feel like she owes you something or by making her feel like a whore because you expect sex in return.\""},

		Highlight{
			Title:           "The Manual: What Women Want and How to Give It to Them",
			AuthorFirstName: "W.",
			AuthorLastName:  "Anton",
			Text:            "\"if a woman tells you that she wants to meet you again, you must not act surprised to hear it as if you were expecting her to turn you down. That will only make her suspicious and doubtful about whether she made the right choice"},

		Highlight{
			Title:           "The Manual: What Women Want and How to Give It to Them",
			AuthorFirstName: "W.",
			AuthorLastName:  "Anton",
			Text:            "If a woman does behave badly and deserves to be put in her place, you have to do so and not treat her differently only because she is beautiful.\""},

		Highlight{
			Title:           "Models: Attract Women Through Honesty",
			AuthorFirstName: "Mark",
			AuthorLastName:  "Manson",
			Text:            "A man who feels like he needs to buy or steal a woman’s attention or affection through entertainment, money or superficiality is a man who is not confident in his identity and who is not genuinely attractive."},

		Highlight{
			Title:           "Models: Attract Women Through Honesty",
			AuthorFirstName: "Mark",
			AuthorLastName:  "Manson",
			Text:            "Non-neediness is when a man places a higher priority on his own perception of himself than the perceptions of others."},

		Highlight{
			Title:           "Models: Attract Women Through Honesty",
			AuthorFirstName: "Mark",
			AuthorLastName:  "Manson",
			Text:            "As with any type of failure, it’s not until you’ve been rejected a certain amount that you realize how insignificant it actually is, how you spent so much time worrying about nothing, and how you’re free to act however you choose"},
	)

	//for _, v := range mockValues{
	//	v.Text = punctuations.RemoveQuotations(v.Text)
	//	v.Text = punctuations.SetFullStop(v.Text)
	//	punctuations.RemoveFullStop(v.Text)
	//}

	//printHighlight(mockValues)
	//setSingleQuotations(mockValues)
}

func greetMsg() {
	fmt.Printf(msgWelcome, programVersion)
	fmt.Println("\n ")
}

func readSrcPath() {
	fmt.Print(msgSetSrcPath)
	for {
		if !scanner.Scan() && scanner.Err() != nil {
			fmt.Println("Error:", scanner.Err())
			return
		}
		srcPath = scanner.Text()
		if len(srcPath) > 0 && fileExists(srcPath) {
			fmt.Println(srcPath + "\n")
			return
		} else {
			fmt.Print("Error! Couldn't find text file. Try again: ")
		}
	}
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) || info == nil {
		return false
	}
	return !info.IsDir()
}

func readOutputFormats() {
	fmt.Print(msgSetOutputFormats)
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
	fmt.Print(msgSetDstPath)
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

func printHighlight(highlights []Highlight) {
	highlightsCount := 0
	booksCount := make(map[string]int)
	for _, v := range highlights {
		msg := fmt.Sprintf("Title: %s\nAuthor: %s %s\nText: %s\n", v.Title, v.AuthorFirstName, v.AuthorLastName, v.Text)
		fmt.Println(msg)
		if len(v.Text) > 0 {
			highlightsCount++
		}
		if len(v.Title) > 0 {
			booksCount[v.Title] = 0
		}
	}
	msg := fmt.Sprintf("### Found %d highlights from %d different books", highlightsCount, len(booksCount))
	fmt.Println(msg)
}

func readOptionQuotationMarks() {
	fmt.Println(msgSetQuotationMarks)
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
