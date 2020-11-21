package message

import "fmt"

const ProgramVersion = "v1.0.0"

const Greeting = `Welcome! KindleHighlightsReader %s is a program that can read all of your Kindle highlights 
and saves them nicely formatted in either Json, PDF, Text or as an CSV file, with many formatting and styling options.`

const SetSrcPath = `Enter the path for the "My Clippings.txt" file from your Kindle device: `
const SetDstPath = `Enter a destination path for the output file(s) or leave empty for desktop path: `
const SetOutputFormats = `Enter one or more of the following output formats JSON, PDF, CSV, Text separated by spaces: `

const OptionQuotationMarks = `Choose your quotation marks option for every highlight text:
1 - YES (Double quotation "Hi")
2 - YES (Single quotation 'Hi')
3 - REMOVE
4 - SKIP`

const OptionFullstops = `Choose your full-stop/period option for every highlight text:
1 - YES
2 - REMOVE
3 - SKIP`

const OptionCleanTexts = `Clean every highlight with the following: 
1 - Remove white spaces before and after (" Hello"	"Hello"	" Hello ")
2 - Remove 1-3 characters before and after every highlight (".Hello"  "d.Hello"	 "d. Hello")
3 - SKIP`

func GetGreeting() string {
	return fmt.Sprintf(Greeting+"\n", ProgramVersion)
}
