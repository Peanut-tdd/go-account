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

func Arrcmp(src []string, dest []string) ([]string, []string) {
	msrc := make(map[string]byte) //按源数组建索引
	mall := make(map[string]byte) //源+目所有元素建索引

	var set []string //交集

	//1.源数组建立map
	for _, v := range src {
		msrc[v] = 0
		mall[v] = 0
	}
	//2.目数组中，存不进去，即重复元素，所有存不进去的集合就是并集
	for _, v := range dest {
		l := len(mall)
		mall[v] = 1
		if l != len(mall) { //长度变化，即可以存
			l = len(mall)
		} else { //存不了，进并集
			set = append(set, v)
		}
	}
	//3.遍历交集，在并集中找，找到就从并集中删，删完后就是补集（即并-交=所有变化的元素）
	for _, v := range set {
		delete(mall, v)
	}
	//4.此时，mall是补集，所有元素去源中找，找到就是删除的，找不到的必定能在目数组中找到，即新加的
	var added, deleted []string
	for v, _ := range mall {
		_, exist := msrc[v]
		if exist {
			deleted = append(deleted, v)
		} else {
			added = append(added, v)
		}
	}

	return deleted, added
}
