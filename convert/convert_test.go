package convert

import (
	"KindleHighlightsReader/reader"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestToJSON(t *testing.T) {
	b := ToJSON(getTestData())
	ok := json.Valid(b)
	if !ok {
		t.Errorf("Failed to correctly convert highlights to JSON")
	}
}

func TestToText(t *testing.T) {
	highlights := getTestData()
	b := ToText(highlights)
	content := string(b)
	if len(content) < 1 {
		t.Error("Failed to convert highlights to text. No text was found")
	}
	h := highlights[0]
	if !strings.Contains(content, h.Text) {
		t.Error("Failed to convert highlights to text. No matching values was found")
	}
}

func TestToCSV(t *testing.T) {
	highlights := getTestData()
	b := ToCSV(highlights)
	content := string(b)

	h := highlights[0]
	v := reflect.ValueOf(&h).Elem()
	typeOf := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fieldName := typeOf.Field(i).Name
		fieldValue := v.Field(i).Interface()
		fieldValueTxt := fmt.Sprintf("%v", fieldValue)
		if !strings.Contains(content, fieldName) {
			t.Error("Failed to convert highlights to CSV. No headers was found")
		}
		if !strings.Contains(content, fieldValueTxt) {
			t.Error("Failed to convert highlights to CSV. No values was found")
		}
	}

}

func getTestData() []reader.Highlight {
	h1 := reader.Highlight{
		Text:   "Hello World!",
		Author: "Muddz",
		Title:  "Developer",
	}

	h2 := reader.Highlight{
		Text:   "Hello World!",
		Author: "Muddz",
		Title:  "Developer",
	}
	return []reader.Highlight{h1, h2}
}
