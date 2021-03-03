package filefinder

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestGetMyClippingsFile(t *testing.T) {
	f, err := makeTestFileForDesktop()
	if err != nil {
		t.Error("failed to setup test file:", err)
	}

	cf := GetMyClippingsFile()
	if cf != f.Name() {
		t.Error("failed to automatically find test file:", f.Name())
	}

	t.Cleanup(func() {
		if err := os.Remove(f.Name()); err != nil {
			log.Print("failed to remove test file: ", f.Name())
		}
	})
}

func makeTestFileForDesktop() (*os.File, error) {
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
