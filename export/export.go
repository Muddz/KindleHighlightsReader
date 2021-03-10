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
	return export(b, "MyClippings.txt")
}

func AsJSON(h []highlight.Highlight) (string, error) {
	b, err := convert.ToJSON(h)
	if err != nil {
		return "", fmt.Errorf("\n%w", err)
	}
	return export(b, "MyClippings.json")
}

func AsCSV(h []highlight.Highlight) (string, error) {
	b, err := convert.ToCSV(h)
	if err != nil {
		return "", fmt.Errorf("\n%w", err)
	}
	return export(b, "MyClippings.csv")
}

func AsPDF(h []highlight.Highlight) (string, error) {
	b, err := convert.ToPDF(h)
	if err != nil {
		return "", fmt.Errorf("\n%w", err)
	}
	return export(b, "MyClippings.pdf")
}

func export(b []byte, filename string) (string, error) {
	path, err := makePath(filename)
	if err != nil {
		return "", fmt.Errorf("\n%w", err)
	}

	f, err := createFile(path)
	if err != nil {
		return "", fmt.Errorf("\n%w", err)
	}
	defer f.Close()

	if err = writeToFile(f, b); err != nil {
		return "", fmt.Errorf("\n%w", err)
	}
	return f.Name(), nil
}

func createFile(path string) (*os.File, error) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return nil, fmt.Errorf("\n%w", err)
	}
	return f, nil
}

func makePath(filename string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("\n%w", err)
	}
	fs := string(filepath.Separator)
	return fmt.Sprintf("%s%sDesktop%s%s", homeDir, fs, fs, filename), nil
}

func writeToFile(f *os.File, b []byte) error {
	_, err := f.Write(b)
	if err != nil {
		return fmt.Errorf("\n%w", err)
	}
	return nil
}
