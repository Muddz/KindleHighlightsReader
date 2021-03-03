package filefinder

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func GetMyClippingsFile() string {
	if p := searchInDesktop(); p != "" {
		return p
	}
	return searchInDevice()
}

func searchInDesktop() string {
	fs := string(filepath.Separator)
	p := fmt.Sprintf("%s%sMy Clippings.txt", getDesktopPath(), fs)
	if fileExist(p) {
		return p
	}
	return ""
}

func searchInDevice() string {
	if runtime.GOOS == "windows" {
		return searchInDeviceWindows()
	}
	return searchInDeviceUnix()
}

func searchInDeviceWindows() string {
	drivers := []string{"D", "E", "F", "G", "H", "I"}
	for _, v := range drivers {
		path := fmt.Sprintf("%s:\\documents\\My Clippings.txt", v)
		if fileExist(path) {
			return path
		}
	}
	return ""
}

func searchInDeviceUnix() string {
	p := "/Volumes/Kindle/documents/My Clippings.txt"
	if fileExist(p) {
		return p
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
