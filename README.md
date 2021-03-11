# KindleHighlightReader <img width="35px" src="https://www.flaticon.com/svg/static/icons/svg/845/845938.svg">


KindleHighlightsReader is a program that reads your highlights from your Kindle with options to clean and format your highlights before exporting them in any of following formats: text, json, csv or pdf.


## Features
- Trim highlights for redundant characters before and after: `ed. Hello` > `Hello`
- Insert or remove periods on all highlights.
- Insert or remove quotation marks on all highlights.
- Capitalize first letters on all highlights.
- Export as .txt .json .csv .pdf files.
- Compatible with Windows and MacOS

#### JSON

The JSON format is made for developers and is exported as an .json file and has the following output format:

```
[
  {
    "Title": "",
    "Author": "",
    "Text": ""
  },
]
```

## Download and Usage

The program automatically finds your highlights by looking for the `My Clippings.txt` file in your Kindle if its plugged in to your computer or if the file is on your Desktop. You can also specify the path to the file yourself.

The file can be found in your Kindle device at the following path `/Kindle/documents/My clippings.txt`

**Windows**  
Download and run the .exe file which will open the program in a command prompt window.

[Download for Windows](https://github.com/Muddz/KindleHighlightsReader/releases/download/1.0.0/KindleHighlightsReader_v1.0.0.exe)

**MacOS**  
Download the binary file, right click on it and *open with terminal*.
This is only necessary the first time, any other time just double click on the file itself.

[Download for MacOS](https://github.com/Muddz/KindleHighlightsReader/releases/download/1.0.0/KindleHighlightsReader_MacOS_v1.0.0)


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
