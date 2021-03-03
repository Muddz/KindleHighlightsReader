package message

import "fmt"

const Greeting = `KindleHighlightsReader v.1.0.0 is a program that reads your highlights from your Kindle 
with options to clean and format your highlights before exporting them 
in any of following formats: text, json, csv or pdf.`

const EnterSource = `Enter the path for "My Clippings.txt" or drag the file into the window: `

const EnterTrimOptions = `(1/5) Choose trimming options for highlights separated by spaces:
[1] Trim redundant characters before the text ('the. Hello' > 'Hello') 
[2] Trim redundant characters after the text  ('Hello. the' > 'Hello')
[3] Skip`

const EnterQuotationOption = `(2/5) Choose a quotations option for highlights:
[1] Insert quotations
[2] Remove all quotations
[3] Skip`

const EnterPeriodOption = `(3/5) Choose a period option for highlights:
[1] Insert periods
[2] Remove periods
[3] Skip`

const EnterCapitalizationOption = `(4/5) Capitalize highlights:
[1] Yes (hello > Hello)
[2] Skip`

const EnterExportOptions = `(5/5) Choose your export formats: text, json, csv, pdf, separated by spaces: `

func GetGreeting() string {
	return fmt.Sprintf("%s\n", Greeting)
}
