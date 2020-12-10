package export

import (
	"log"
	"os"
	"testing"
)

func TestAsJSON(t *testing.T) {
	path, _ := AsJSON(nil)
	filename := getUserDesktopPath() + "\\My Clippings JSON.txt"
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

func fileExist(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func TestCreateNewFile(t *testing.T) {
	filename := "./testfile.txt"
	f := createFile(filename)
	if f == nil {
		t.Error("Failed to create file:", filename)
	}

	t.Cleanup(func() {
		tearDownTestFile(f)
	})
}

func TestWriteToFile(t *testing.T) {
	filename := "./testfile.txt"
	content := "Hello World!"
	data := []byte(content)
	f := createFile(filename)
	ok := writeToFile(f, data)

	if !ok {
		t.Error("Failed to write ''Hello World' to file:", filename)
	}

	t.Cleanup(func() {
		tearDownTestFile(f)
	})
}

func tearDownTestFile(f *os.File) {
	if err := f.Close(); err != nil {
		log.Print("Failed to close test file: ", f.Name())
	}
	if err := os.Remove(f.Name()); err != nil {
		log.Print("Failed to remove test file: ", f.Name())
	}
}
