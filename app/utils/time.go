package utils

import (
	"database/sql/driver"
	"time"
)

const TimeFormat = "2006-01-02 15:04:05"

type SelfTime time.Time

// 写入 mysql 时调用
func (t SelfTime) Value() (driver.Value, error) {
	// 0001-01-01 00:00:00 属于空值，遇到空值解析成 null 即可
	if t.String() == "0001-01-01 00:00:00" {
		return nil, nil
	}
	return []byte(time.Time(t).Format(TimeFormat)), nil
}

// 检出 mysql 时调用
func (t *SelfTime) Scan(v interface{}) error {
	// mysql 内部日期的格式可能是 2006-01-02 15:04:05 +0800 CST 格式，所以检出的时候还需要进行一次格式化
	tTime, _ := time.Parse("2006-01-02 15:04:05 +0800 CST", v.(time.Time).String())
	*t = SelfTime(tTime)
	return nil
}

// 用于 fmt.Println 和后续验证场景
func (t SelfTime) String() string {
	return time.Time(t).Format(TimeFormat)
}


func (t SelfTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TimeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, TimeFormat)
	b = append(b,'"')
	return b, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// The time is expected to be a quoted string in RFC 3339 format.
func (t *SelfTime) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	var err error

	str := string(data)

	t1, err := time.ParseInLocation(TimeFormat,str,time.Local)
	*t = SelfTime(t1)
	return err
}
