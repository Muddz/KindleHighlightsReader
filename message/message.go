package message

import "fmt"

const Greeting = `KindleHighlightsReader 1.0.0 is a program that reads your highlights 
from your Kindle with options to clean and style the them and 
exports them in any of the following formats: Text, Json, CSV or PDF.`

const EnterSource = `Enter the path for "My Clippings.txt" or drag the file into the window: `

const EnterTrimOptions = `Choose trimming options for highlights separated by spaces:
[1] Trim redundant characters before the text ('the. Hello' > 'Hello') 
[2] Trim redundant characters after the text  ('Hello. the' > 'Hello')
[3] Skip`

const EnterPeriodOption = `Choose a period option for highlights:
[1] Insert periods
[2] Remove periods
[3] Skip`

const EnterQuotationOption = `Choose a quotations option for highlights:
[1] Insert double quotations
[2] Remove all quotations
[3] Skip`

const EnterCapitalizationOption = `Capitalize highlights:
[1] Yes (hello > Hello)
[2] Skip`

const EnterExportOptions = `Choose your export formats: text, json, csv, pdf, separated by spaces: `

func GetGreeting() string {
	return fmt.Sprintf("%s\n", Greeting)
}
