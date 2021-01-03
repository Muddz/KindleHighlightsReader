package export

import (
	"KindleHighlightsReader/convert"
	"KindleHighlightsReader/reader"
	"fmt"
	"os"
	"runtime"
)

func AsTxt(highlights []reader.Highlight) (string, error) {
	b := convert.ToText(highlights)
	path, err := export("My Clippings new.txt", b)
	return path, err
}

func AsJSON(highlights []reader.Highlight) (string, error) {
	b := convert.ToJSON(highlights)
	path, err := export("My Clippings.json", b)
	return path, err
}

func AsCSV(highlights []reader.Highlight) (string, error) {
	b := convert.ToCSV(highlights)
	path, err := export("My Clippings.csv", b)
	return path, err
}

func AsPDF(highlights []reader.Highlight) (string, error) {
	b := convert.ToPDF(highlights)
	path, err := export("My Clippings.pdf", b)
	return path, err
}

func export(filename string, b []byte) (string, error) {
	path := makePath(filename)
	f, err := createFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to create file at %s", path)
	}
	if err = writeToFile(f, b); err != nil {
		return "", fmt.Errorf("failed writing to file at %s", path)
	}
	return f.Name(), nil
}

func createFile(path string) (*os.File, error) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func makePath(filename string) string {
	homeDir, _ := os.UserHomeDir()
	if runtime.GOOS == "windows" {
		return fmt.Sprintf("%s\\Desktop\\%s", homeDir, filename)
	} else {
		return fmt.Sprintf("%s/Desktop/%s", homeDir, filename)
	}
}

func writeToFile(f *os.File, b []byte) error {
	_, err := f.Write(b)
	if err != nil {
		return err
	}
	defer func() {
		err = f.Close()
	}()
	return err
}
