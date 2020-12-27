package message

import "fmt"

const Greeting = `KindleHighlightsReader 1.0.0 is a program that reads all of your highlights from the My Clippings.txt file 
found in Kindle devices with options to clean and style all highlights 
and exports them in the following file formats: Text, Json, CSV or PDF.`

const EnterSource = `Enter the path for "My Clippings.txt" which can be found in your Kindle device: `
const EnterExportOptions = `Enter one or more of the following export formats: text, json, csv, pdf, separated by spaces: `

const EnterTrimOptions = `Choose trimming options for highlight texts:
(1) Trim 1-5 characters before the texts: "the. Hello"  "th. Hello"  "e. Hello" 
(2) Trim 1-5 characters after the texts: "Hello. the"  "Hello. th"  "Hello." 
(3) Do nothing`

const EnterPeriodOption = `Choose a period option for highlight texts:
(1) Insert periods
(2) Remove periods
(3) Do nothing`

const EnterQuotationOption = `Choose a quotation option for highlight texts:
(1) Double quotation "Hi"
(2) Single quotation 'Hi'
(3) Remove all quotation
(4) Do nothing`

func GetGreeting() string {
	return fmt.Sprintf("%s\n", Greeting)
}
