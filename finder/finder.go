package finder

import (
	"fmt"
	"log"
	"os"
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
	path := fmt.Sprintf("%s\\My Clippings.txt", getUserDesktopPath())
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
	path := fmt.Sprintf("/Kindle/documents/My Clippings.txt")
	return path
}

func getUserDesktopPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Println(err)
	}
	return fmt.Sprintf("%s\\Desktop", homeDir)
}

func fileExist(path string) bool {
	info, err := os.Stat(path)
	if info == nil || os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
