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

func TestSave(t *testing.T) {
	filename := "./testfile.txt"
	content := []byte("Hello World")
	ok := save(filename, content)
	if !destinationExists(filename) && !ok {
		t.Errorf("Failed to save file %s, with content '%s'", filename, string(content))
	}
	_ = os.Remove(filename)
}

func destinationExists(path string) bool {
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
		t.Error("Failed to write \"Hello World\" to file:", filename)
	}
	_ = f.Close()
	_ = os.Remove(filename)
}
