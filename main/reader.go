package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const programVersion = "v1.0.0"
const msgWelcome = `KindleHighlightsReader %s is a little program that reads all of your highlights 
from your Amazon Kindle and can save them nicely formatted in either Json, PDF or as an CSV file`

const (
	msgSetFilePath       = `Enter the full path for the "My Clippings.txt" file: `
	msgSetOutputFormats  = `Enter one or more of the following output formats JSON, PDF, CSV separated by spaces: `
	msgOutputDestination = `Enter a destination path for the output files: `
)

var highlights []Highlight
var validOutputFormats = []string{"JSON", "PDF", "CSV"}

var inputFormats []string
var inputFields []string
var inputSrc string
var inputDst string
var scanner = bufio.NewScanner(os.Stdin)

type Highlight struct {
	Title           string
	AuthorFirstName string
	AuthorLastName  string
	Text            string
}

func main() {
	greetMsg()
	readFilePath()
	fmt.Println("")
	readOutputFormats()
	fmt.Println("")
	readFileOutputDestination()
}

func greetMsg() {
	fmt.Printf(msgWelcome, programVersion)
	fmt.Println("\n ")
}

func readFilePath() {
	fmt.Print(msgSetFilePath)
	for {
		if !scanner.Scan() && scanner.Err() != nil {
			fmt.Println("Error:", scanner.Err())
			return
		}
		inputFilepath := scanner.Text()
		if len(inputFilepath) > 0 && fileExists(inputFilepath) {
			fmt.Println(inputFilepath)
			return
		} else {
			fmt.Print("Error! Couldn't find text file. Try again: ")
		}
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
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
		inputFilepath := scanner.Text()
		if len(inputFilepath) > 0 && validateOutputValues(inputFilepath) {
			fmt.Println(inputFilepath)
			return
		} else {
			fmt.Printf("Error! Invalid formats. Try again: ")
		}
	}
}

func validateOutputValues(outputFormats string) bool {
	var result = false
	formats := strings.Fields(outputFormats)[:len(validOutputFormats)]
	for _, format := range formats {
		for _, validFormat := range validOutputFormats {
			result = strings.EqualFold(format, validFormat)
			if result {
				break
			}
		}
	}
	return result
}

func readFileOutputDestination() {
	fmt.Print(msgOutputDestination)
	for {
		if !scanner.Scan() && scanner.Err() != nil {
			fmt.Println("Error:", scanner.Err())
			return
		}
		inputFilepath := scanner.Text()
		if len(inputFilepath) > 0 && destinationExists(inputFilepath) {
			fmt.Println(inputFilepath)
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

//
//func readExcludedFields()bool{
//
//}
//
//func getFileContent(path string) string(){
//
//}

func printHighlight(title string, authorFn string, authorLn string, text string) {
	msg := fmt.Sprintf("Title: %s	|	Author: %s %s\n%s", title, authorFn, authorLn, text)
	fmt.Println(msg)
}
