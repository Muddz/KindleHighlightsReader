package filefinder

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
)

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
	return f, nil
}

func TestGetMyClippingsFile(t *testing.T) {
	f, err := makeTestFileForDesktop()
	if err != nil {
		t.Error("Failed to setup test file:", err)
	}

	mcf := GetMyClippingsFile()
	if mcf != f.Name() {
		t.Error("Failed to automatically test file: ", f.Name())
	}

	t.Cleanup(func() {
		tearDownTestFile(f)
	})
}

func tearDownTestFile(f *os.File) {
	if err := f.Close(); err != nil {
		log.Print("Failed to close test file: ", f.Name())
	}
	if err := os.Remove(f.Name()); err != nil {
		log.Print("Failed to remove test file: ", f.Name())
	}
}
