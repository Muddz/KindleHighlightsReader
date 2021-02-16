package export

import (
	"KindleHighlightsReader/convert"
	"KindleHighlightsReader/highlight"
	"fmt"
	"os"
	"path/filepath"
)

func AsTxt(h []highlight.Highlight) (string, error) {
	b := convert.ToText(h)
	return export("MyClippings.txt", b)
}

func AsJSON(h []highlight.Highlight) (string, error) {
	b, err := convert.ToJSON(h)
	if err != nil {
		return "", err
	}
	return export("MyClippings.json", b)
}

func AsCSV(h []highlight.Highlight) (string, error) {
	b, err := convert.ToCSV(h)
	if err != nil {
		return "", err
	}
	return export("MyClippings.csv", b)
}

func AsPDF(h []highlight.Highlight) (string, error) {
	b, err := convert.ToPDF(h)
	if err != nil {
		return "", err
	}
	return export("MyClippings.pdf", b)
}

func export(filename string, b []byte) (string, error) {
	path := makePath(filename)
	f, err := createFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to create file at %s error: %v", path, err)
	}
	if err = writeToFile(f, b); err != nil {
		return "", fmt.Errorf("failed writing to file %s error: %v", path, err)
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
	fs := string(filepath.Separator)
	return fmt.Sprintf("%s%sDesktop%s%s", homeDir, fs, fs, filename)
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
