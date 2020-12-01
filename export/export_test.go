package export

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestAsJSON(t *testing.T) {
	filename := getUserDesktopPath() + "\\My Clippings JSON.txt"
	ok := AsJSON(nil)
	if !ok {
		t.Error("Failed to export to JSON. Result was not ok")
	}
	if !fileExist(filename) {
		t.Error("Failed to find JSON file: ", filename)
	}
	_ = os.Remove(filename)
}

func TestAsTxt(t *testing.T) {
	filename := getUserDesktopPath() + "\\My Clippings Text.txt"
	ok := AsTxt(nil)
	if !ok {
		t.Error("Failed to export to Text. Result was not ok")
	}
	if !fileExist(filename) {
		t.Error("Failed to find Text file: ", filename)
	}
	_ = os.Remove(filename)
}

func TestAsPDF(t *testing.T) {
	filename := getUserDesktopPath() + "\\My Clippings.pdf"
	ok := AsPDF(nil)
	if !ok {
		t.Error("Failed to export to PDF. Result was not ok")
	}
	if !fileExist(filename) {
		t.Error("Failed to find PDF file: ", filename)
	}
	_ = os.Remove(filename)
}

func TestAsCSV(t *testing.T) {
	filename := getUserDesktopPath() + "\\My Clippings.csv"
	ok := AsCSV(nil) //TODO this will throw an error, but still correctly passes the test
	if !ok {
		t.Error("Failed to export to csv. Result was not ok")
	}
	if !fileExist(filename) {
		t.Error("Failed to find csv file: ", filename)
	}
	_ = os.Remove(filename)
}

func TestSave(t *testing.T) {
	filename := "./testfile.txt"
	content := []byte("Hello World")
	ok := save(filename, content)

	if !fileExist(filename) && !ok {
		t.Errorf("Failed to export file %s, with content '%s'", filename, string(content))
	}

	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("Failed to read file:", filename)
	}

	if string(b) != string(content) {
		t.Error("Written content didn't match read content in file", filename)
	}
	_ = os.Remove(filename)
}

func fileExist(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func TestCreateNewFile(t *testing.T) {
	filename := "./testfile.txt"
	f := createNewFile(filename)
	if f == nil {
		t.Error("Failed to create file:", filename)
	}
	_ = f.Close()
	_ = os.Remove(filename)
}

func TestWriteToFile(t *testing.T) {
	filename := "./testfile.txt"
	content := "Hello World!"
	data := []byte(content)
	f := createNewFile(filename)
	ok := writeToFile(f, data)

	if !ok {
		t.Error("Failed to write ''Hello World' to file:", filename)
	}
	_ = f.Close()
	_ = os.Remove(filename)
}
