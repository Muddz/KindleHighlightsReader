package message

import "fmt"

const Greeting = `KindleHighlightsReader 1.0.0 is a program that reads all of your highlights from the "My Clippings.txt" file 
found in your Kindle device with options to clean and style highlights 
and exports them in the following file formats: Text, Json, CSV or PDF.`

const EnterSource = `Enter the path for "My Clippings.txt" which can be found in your Kindle device or drag the file into the window: `

const EnterTrimOptions = `Choose trimming options for highlights:
[1] Trim 1-5 redundant characters before the text ('the. Hello' > 'Hello') 
[2] Trim 1-5 redundant characters after the text  ('Hello. the' > 'Hello')
[3] Skip`

const EnterPeriodOption = `Choose a period option for highlights:
[1] Insert periods
[2] Remove periods
[3] Skip`

const EnterQuotationOption = `Choose a quotation marks option for highlights:
[1] Insert double quotations
[2] Remove all quotations
[3] Skip`

const EnterCapitalizationOption = `Capitalize highlights:
[1] Yes (hello > Hello)
[2] Skip`

const EnterExportOptions = `Enter one or more of the following export formats: text, json, csv, pdf, separated by spaces: `

func GetGreeting() string {
	return fmt.Sprintf("%s\n", Greeting)
}
