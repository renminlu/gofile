package tree

import (
	"os"
	"path/filepath"
)

//递进子目录：仅获取目录……
func GetDirList(dirpath string) (*[]string, error) {
	var ls []string
	dir_err := filepath.Walk(
		dirpath,
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err //retrun err会退出循环遍历
			}
			if f.IsDir() {
				ls = append(ls, path)
				return nil
			}

			return nil
		})
	return &ls, dir_err
}

//递进子目录：仅获取文件……
func GetFileList(dirpath string) (*[]string, error) {
	var ls []string
	dir_err := filepath.Walk(
		dirpath,
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err //retrun err会退出循环遍历
			}
			if !f.IsDir() {
				ls = append(ls, path)
				return nil
			}

			return nil
		})
	return &ls, dir_err
}

//递进子目录：获取目录和文件……
func GetList(dirpath string) (*[]string, error) {
	var ls []string
	dir_err := filepath.Walk(
		dirpath,
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err //retrun err会退出循环遍历
			}

			ls = append(ls, path)
			return nil
		})
	return &ls, dir_err
}
