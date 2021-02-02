# KindleHighlightReader <img width="35px" src="https://www.flaticon.com/svg/static/icons/svg/845/845938.svg">


KindleHighlightsReader is a program that reads your highlights from your Kindle with options to clean and style the highlights and exports them in the following formats: text, json, csv or pdf.


## Features in 1.0.0
- Trim highlights for redundant characters before and after like: `ed. Hello` > `Hello`
- Insert periods on all highlights.
- Insert double quotations on all highlights.
- Remove any quotation from all highlights.
- Capitalize first letter on all highlights.
- Export as .txt .json .csv or .pdf files.
- Works for Windows and MacOS

#### JSON

The JSON format is useful for developers who wants to unmarshal the JSON string to objects in any langauge to be used for an app, webpage or database. The JSON output is saved in a `.json` file and has the following format:

```
[
  {
    "Title": "",
    "Author": "",
    "Text": ""
  },
]
```

## Usage

Download and run the program which will open in a command prompt window and automatically find your `My Clippings.txt` file if its already on the Desktop or if your Kindle is plugged into your computer. 

- [Download for Windows](https://github.com/Muddz/KindleHighlightReader/raw/master/KindleHighlightsReade.exe)
- [Download for MacOS](https://github.com/Muddz/KindleHighlightReader/raw/master/KindleHighlightsReaderMacOS)
- Or clone the project and run it from your IDE.


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
