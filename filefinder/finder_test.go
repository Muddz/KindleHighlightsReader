package filefinder

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestGetMyClippingsFile(t *testing.T) {
	testFile, err := getDesktopTestFile()
	if err != nil {
		t.Error("failed to setup test file:", err)
	}

	f := GetMyClippingsFile()
	if f != testFile.Name() {
		t.Errorf("failed to find file: %s, found instead: %s", testFile.Name(), f)
	}

	t.Cleanup(func() {
		if err := os.Remove(testFile.Name()); err != nil {
			log.Print("failed to remove test file: ", testFile.Name())
		}
	})
}

func getDesktopTestFile() (*os.File, error) {
	userDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	fs := string(filepath.Separator)
	fn := fmt.Sprintf("%s%sDesktop%sMy Clippings.txt", userDir, fs, fs)
	f, err := os.OpenFile(fn, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return f, nil
}
