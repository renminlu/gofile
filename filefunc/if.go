package fileFunc

import (
	"os"
)

//判断文件或目录是否存在
func FileExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

//判断：是文件=true，是目录或其他=false
func IsFile(f string) bool {
	fi, e := os.Stat(f)
	if e != nil {
		return false
	}
	return !fi.IsDir()
}

func IsDir(f string) bool {

	fi, e := os.Stat(f)
	if e != nil {
		return false
	}
	return fi.IsDir()
}
