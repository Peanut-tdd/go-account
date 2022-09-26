package utils

import (
	"archive/zip"
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
)

func Unzip(filepath string, dstDir string) error {
	reader, err := zip.OpenReader(filepath)
	if err != nil {
		return err
	}
	defer reader.Close()

	for _, file := range reader.File {
		if err := unzipFile(file, dstDir); err != nil {
			return err
		}
	}
	return nil
}

func unzipFile(file *zip.File, dstDir string) error {
	filePath := path.Join(dstDir, file.Name)
	if file.FileInfo().IsDir() {
		if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
			return err
		}
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}
	// open the file
	rc, err := file.Open()
	if err != nil {
		return err
	}
	defer rc.Close()

	// create the file
	w, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer w.Close()

	// save the decompressed file content
	_, err = io.Copy(w, rc)
	return err
}

func UnzipFile(zipFile string, destDir string, pattern string) error {

	zipReader, err := zip.OpenReader(zipFile)

	if err != nil {

		return err

	}

	defer zipReader.Close()

	var decodeName string

	for _, f := range zipReader.File {

		if f.Flags == 0 {

			//如果标致位是0  则是默认的本地编码   默认为gbk

			i := bytes.NewReader([]byte(f.Name))

			decoder := transform.NewReader(i, simplifiedchinese.GB18030.NewDecoder())

			content, _ := ioutil.ReadAll(decoder)

			decodeName = string(content)

		} else {

			//如果标志为是 1 << 11也就是 2048  则是utf-8编码

			decodeName = f.Name

		}

		if pattern != "" {
			re := regexp.MustCompile(pattern)
			match := re.MatchString(decodeName)
			if match {
				continue
			}
		}

		fpath := filepath.Join(destDir, decodeName)
		if f.FileInfo().IsDir() {

			os.MkdirAll(fpath, os.ModePerm)

		} else {

			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {

				return err

			}

			inFile, err := f.Open()

			if err != nil {

				return err

			}

			defer inFile.Close()

			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())

			if err != nil {

				return err

			}

			defer outFile.Close()

			_, err = io.Copy(outFile, inFile)

			if err != nil {

				return err

			}

			wr(fpath, re(fpath))

		}

	}

	return nil

}

func re(filepath string) string {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	// GBK转UTF8
	data, err = ioutil.ReadAll(transform.NewReader(bytes.NewBuffer(data), simplifiedchinese.GB18030.NewDecoder()))
	if err != nil {
		log.Fatal(err)
	}

	// 打印UTF8
	return string(data)

}

func wr(filepath, content string) {
	err := ioutil.WriteFile(filepath, []byte(content), 0666)
	if err != nil {
		log.Fatal(err)
	}
	return
}
