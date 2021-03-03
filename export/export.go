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
	f, err := export("MyClippings.txt", b)
	if err != nil {
		return "", fmt.Errorf("\n%w", err)
	}
	return f, nil
}

func AsJSON(h []highlight.Highlight) (string, error) {
	b, err := convert.ToJSON(h)
	if err != nil {
		return "", fmt.Errorf("\n%w", err)
	}
	f, err := export("MyClippings.json", b)
	if err != nil {
		return "", fmt.Errorf("\n%w", err)
	}
	return f, nil
}

func AsCSV(h []highlight.Highlight) (string, error) {
	b, err := convert.ToCSV(h)
	if err != nil {
		return "", fmt.Errorf("\n%w", err)
	}
	f, err := export("MyClippings.csv", b)
	if err != nil {
		return "", fmt.Errorf("\n%w", err)
	}
	return f, nil
}

func AsPDF(h []highlight.Highlight) (string, error) {
	b, err := convert.ToPDF(h)
	if err != nil {
		return "", fmt.Errorf("\n%w", err)
	}
	f, err := export("MyClippings.pdf", b)
	if err != nil {
		return "", fmt.Errorf("\n%w", err)
	}
	return f, nil
}

func export(filename string, b []byte) (string, error) {
	path, err := makePath(filename)
	if err != nil {
		return "", fmt.Errorf("\n%w", err)
	}

	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return "", fmt.Errorf("\n%w", err)
	}
	defer f.Close()

	if err = writeToFile(f, b); err != nil {
		return "", fmt.Errorf("\n%w", err)
	}
	return f.Name(), nil
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
	defer f.Close()
	return nil
}
