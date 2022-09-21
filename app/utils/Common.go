package utils

import (
	"time"
)

const STORAGE = "./storage"

func CurrentDateTime() string {
	datetime := time.Now().Format("2006-01-02 15:04:05")
	return datetime
}

func CurrentYmd() string {

	Y := time.Now().Format("2006")
	M := time.Now().Format("01")
	D := time.Now().Format("02")
	//hour := time.Now().Format("15")
	//min := time.Now().Format("04")
	//second := time.Now().Format("05")

	return Y + "_" + M + "_" + D

}

//CsvFileDir 生成下载文件地址
func CsvFileDir(sourceType int) string {
	var fileDir string
	switch sourceType {
	//微信
	case 1:
		fileDir = STORAGE + "/download/wx/"
		break
	case 2: //快手
		fileDir = STORAGE + "/download/ks/" + CurrentYmd() + "/"
		break
	}
	return fileDir
}

//CsvFileName 生成下载文件名
func CsvFileName(sourceType int) string {
	dateTime := CurrentYmd()
	var filename string
	switch sourceType {
	//微信
	case 1:
		filename = dateTime + "_bill.csv"
		break
	case 2: //快手
		filename = dateTime + "_bill.zip"
		break
	}

	return filename
}

func CsvFilePath(sourceType int) string {
	return CsvFileDir(sourceType) + CsvFileName(sourceType)
}

//StringToTime 字符串转时间
func StringToTime(tm string) time.Time {
	todayZero, _ := time.ParseInLocation("2006-01-02 15:04:05", tm, time.Local)
	return todayZero
}
