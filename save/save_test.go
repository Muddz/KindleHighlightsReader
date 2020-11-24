package save

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestToJSON(t *testing.T) {

	to
}

func TestToTxt(t *testing.T) {

}

func TestToPDF(t *testing.T) {

}

func TestToCSV(t *testing.T) {

}

func TestSave(t *testing.T) {
	filename := "./testfile.txt"
	content := []byte("Hello World")
	ok := save(filename, content)

	if !fileExist(filename) && !ok {
		t.Errorf("Failed to save file %s, with content '%s'", filename, string(content))
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
	return info.IsDir()
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
