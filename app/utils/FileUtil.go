package utils

import (
	"os"
	"path/filepath"
)

//TPFuncReadDirFiles 获得文件夹下所有文件
func TPFuncReadDirFiles(dir string) ([]string, error) {
	var files []string
	//方法一
	var walkFunc = func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		//fmt.Printf("%s\n", path)
		return nil
	}
	err := filepath.Walk(dir, walkFunc)

	return files, err
}
