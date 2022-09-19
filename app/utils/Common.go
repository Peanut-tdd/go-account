package utils

import (
	"time"
)
const STORAGE =  "./storage"

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

func CsvFilePath(sourceType int) string {

	dateTime := CurrentYmd()
	var CsvSource string
	switch sourceType {
	case 1:
		CsvSource = "wx"
		break
	case 2:
		CsvSource = "ks"
		break
	}

	return STORAGE + "/download/" + CsvSource + "/" + dateTime + "_bill.csv"

}
