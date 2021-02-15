package export

import (
	"KindleHighlightsReader/convert"
	"KindleHighlightsReader/highlight"
	"fmt"
	"os"
	"runtime"
)

func AsTxt(h []highlight.Highlight) (string, error) {
	b := convert.ToText(h)
	path, err := export("MyClippings.txt", b)
	return path, err
}

func AsJSON(h []highlight.Highlight) (string, error) {
	b := convert.ToJSON(h)
	path, err := export("MyClippings.json", b)
	return path, err
}

func AsCSV(h []highlight.Highlight) (string, error) {
	b := convert.ToCSV(h)
	path, err := export("MyClippings.csv", b)
	return path, err
}

func AsPDF(h []highlight.Highlight) (string, error) {
	b := convert.ToPDF(h)
	path, err := export("MyClippings.pdf", b)
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
