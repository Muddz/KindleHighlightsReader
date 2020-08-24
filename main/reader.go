package main

import (
	"fmt"
	"log"
)

type Highlight struct {
	Title           string
	AuthorFirstName string
	AuthorLastName  string
	Text            string
}

const regexPattern = "/[a-Z]"

var highlights []Highlight

func main() {

}

func printHighlight(title string, authorFn string, authorLn string, text string) {
	msg := fmt.Sprintf("Title: %s	|	Author: %s %s\n%s", title, authorFn, authorLn, text)
	log.Println(msg)
}
