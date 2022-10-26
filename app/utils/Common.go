package utils

import (
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
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
func CsvFileDir(sourceType int, payChannel int) string {
	var fileDir string
	fileDir = STORAGE + "/download"
	switch sourceType {
	//微信
	case 1:
		fileDir = fileDir + "/wx/"
		break
	case 2: //快手
		fileDir = fileDir + "/ks/" + CurrentYmd() + "/"
		break
	case 4: //快应用
		switch payChannel {
		case 2: //支付宝
			fileDir = fileDir + "/fapp/alipay/" + CurrentYmd() + "/"
			break
		}
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
	case 4:
		filename = dateTime + "_bill.zip"
	}

	return filename
}

func CsvFilePath(sourceType int, payChannel int) string {
	return CsvFileDir(sourceType, payChannel) + CsvFileName(sourceType)
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

func CheckInsertData(db_trade_no []string, insertSliceMapBills []map[string]interface{}, mapBills map[string]map[string]interface{}) []map[string]interface{} {
	if len(db_trade_no) == 0 {
		for _, value := range mapBills {
			insertSliceMapBills = append(insertSliceMapBills, value)
		}
	} else {
		for _, trade_no := range db_trade_no {
			if value, ok := mapBills[trade_no]; ok {
				//updateSliceMapBills = append(updateSliceMapBills, value)
			} else {
				insertSliceMapBills = append(insertSliceMapBills, value)
			}
		}
	}
	return insertSliceMapBills
}

func UrlDownLoad(url string, path string) {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// Create output file
	err = os.MkdirAll(filepath.Dir(path), os.ModeDir)
	if err != nil {
		panic(err)
	}

	out, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer out.Close()
	// copy stream
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}
}

func JsonEncode(data interface{}) {
	json, err := jsoniter.Marshal(data)

	if err != nil {
		fmt.Println("json marshal failed")

	}

	fmt.Printf("json：%s", json)
	return

}


func Strval(value interface{}) string {
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}
