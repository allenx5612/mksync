package execpath

import (
	"os"
	"path/filepath"
)

// ExecPath used to return executable dir
func ExecPath() (path string) {
	execFile, err := os.Executable()
	if err != nil {
		panic(err)
	}
	path = filepath.Dir(execFile)
	return path
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
