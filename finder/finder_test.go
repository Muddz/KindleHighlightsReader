package finder

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func setupTestFileForDesktop() (*os.File, error) {
	userDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	var fn string
	if isWindowsOS() {
		fn = fmt.Sprintf("%s\\Desktop\\My Clippings.txt", userDir)
	} else {
		fn = fmt.Sprintf("%s/Desktop/My Clippings.txt", userDir)
	}

	f, err := os.OpenFile(fn, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func TestGetMyClippingsFile(t *testing.T) {
	f, err := setupTestFileForDesktop()
	if err != nil {
		t.Error("Failed to setup test file")
	}

	cf := GetMyClippingsFile()

	if cf != f.Name() {
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
