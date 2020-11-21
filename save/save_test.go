package save

import (
	"os"
	"testing"
)

func TestToJSON(t *testing.T) {

}

func TestToTxt(t *testing.T) {

}

func TestToPDF(t *testing.T) {

}

func TestToCSV(t *testing.T) {

}

func TestWriteToFile(t *testing.T) {
	filename := "./testfile.txt"
	content := "Hello World!"
	f := createNewFile(filename)
	data := []byte(content)
	ok := writeToFile(f, data)

	if !ok {
		t.Error("Failed to write \"Hello World\" to file:", filename)
	}
	_ = f.Close()
	_ = os.Remove(filename)
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
