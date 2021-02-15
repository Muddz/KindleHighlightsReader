package filefinder

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func GetMyClippingsFile() string {
	if isWindowsOS() {
		return getMyClippingsWindows()
	} else {
		return getMyClippingsUnix()
	}
}

func isWindowsOS() bool {
	if runtime.GOOS == "windows" {
		return true
	}
	return false
}

func getMyClippingsWindows() string {
	path := fmt.Sprintf("%s\\My Clippings.txt", getDesktopPath())
	if fileExist(path) {
		return path
	}
	drivers := []string{"D", "E", "F", "G", "H", "I"}
	for _, v := range drivers {
		path := fmt.Sprintf("%s:\\documents\\My Clippings.txt", v)
		if fileExist(path) {
			return path
		}
	}
	return ""
}

func getMyClippingsUnix() string {
	path := fmt.Sprintf("%s/My Clippings.txt", getDesktopPath())
	if fileExist(path) {
		return path
	}
	path = "/Volumes/Kindle/documents/My Clippings.txt"
	if fileExist(path) {
		return path
	}
	return ""
}

func getDesktopPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Println(err)
	}
	fs := string(filepath.Separator)
	return fmt.Sprintf("%s%sDesktop", homeDir, fs)
}

func fileExist(path string) bool {
	info, err := os.Stat(path)
	if info == nil || os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
