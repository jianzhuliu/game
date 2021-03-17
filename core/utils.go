package core

import "os"

func PathIsExist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}

	return true
}
