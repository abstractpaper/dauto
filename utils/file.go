package utils

import "os"

// Check whether a file exists
func FileExists(path string) (exists bool, err error) {
	info, err := os.Stat(path)
	exists = !os.IsNotExist(err) && !info.IsDir()
	if !exists {
		err = nil
	}

	return
}
