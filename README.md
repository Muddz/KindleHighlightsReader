# KindleHighlightReader <img width="35px" src="https://www.flaticon.com/svg/static/icons/svg/845/845938.svg">

KindleHighlightReader is a program that reads all of your highlights from the *My Clippings.txt* file found in your Kindle device with options to clean and style all highlights and exports them in the following file formats: `Text`, `Json`, `CSV` or `PDF`. 


The JSON format is useful for developers who want's to unmarshal the JSON to objects in any langauge to be used in for an app, webpage or database. The JSON has the following format:

```
[
  {
    "Title": "",
    "Author": "",
    "Text": ""
  }
]
```


## Features in v1.0.0
- Trim all highlights for words or characters before and after like: `"ed. Hello"`   =>   `"Hello"`
- Insert periods on all highlights.
- Insert double or single quotations on all highlights.
- Remove all quotation from all highlights.
- Export as *.txt*  *.json*  *.csv* or *.pdf* files.
- Works on Windows and MacOS

## Usage

**Windows:** 

Run the `KindleHighlightReader.exe`. The program will open in a CMD window and can automatically find the `My Clippings.txt` file if its on the *Desktop* or if your Kindle device is plugged into your Windows computer. [Download for Windows](https://github.com/Muddz/KindleHighlightReader/raw/master/KindleHighlightsReade.exe)

**MacOs:**
On Mac you just need to run the binary by. [Download for MacOS](https://github.com/Muddz/KindleHighlightReader/raw/master/KindleHighlightsReaderMacOS)


## License

    Copyright 2020 Muddi Walid

    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License
    You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
