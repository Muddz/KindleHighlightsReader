package message

import "fmt"

const Greeting = `Welcome! KindleHighlightsReader 1.0.0 is a program that can read all of your Kindle highlights 
and saves them nicely formatted in the following file formats: Json, PDF, Text or CSV with formatting and styling options.`

const SetSrcPath = `Enter the path for "My Clippings.txt" which can be found in your Kindle device: `
const SetOutputFormats = `Enter one or more of the following output formats: text, json, csv, pdf, separated by spaces: `

const OptionQuotationMarks = `Choose a quotation marks option for highlight texts:
1 - Double quotation "Hi"
2 - Single quotation 'Hi'
3 - Remove all quotation
4 - Do nothing`

const OptionFullstops = `Choose a full stop option for highlight texts:
1 - Insert full stop
2 - Remove full stop
3 - Do nothing`

const OptionCleanTexts = `Choose text cleaning options for highlight texts: 
1 - Remove white spaces before and after (" Hello"	" Hello ")
2 - Remove 1-3 characters before or after every highlight (".Hello"  "d.Hello"  "d. Hello" or "Hello. asd")
3 - Do nothing`

func GetGreeting() string {
	return fmt.Sprintf("%s\n", Greeting)
}
