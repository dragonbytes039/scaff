package fst

import "os"

func DirExist(dir string) bool {

	dirInfo, err := os.Stat(dir)

	if os.IsNotExist(err) {
		return false
	}

	if err != nil {
		return false
	}

	return dirInfo.IsDir()

}
