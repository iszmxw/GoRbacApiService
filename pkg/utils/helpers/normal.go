package helpers

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type TimeNormal struct { // 内嵌方式（推荐）
	time.Time
}

func (t TimeNormal) MarshalJSON() ([]byte, error) {
	tune := t.Format(`"2006-01-02 15:04:05"`)
	value, err := t.Value()
	if err != nil || value == nil {
		return json.Marshal(nil)
	}
	return []byte(tune), nil
}

func (t TimeNormal) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	// 前端接收的时间字符串
	str := string(data)
	// 去除接收的str收尾多余的"
	timeStr := strings.Trim(str, "\"")
	_, err = time.Parse("2006-01-02 15:04:05", timeStr)
	return err
}

// Value insert timestamp into mysql need this function.
func (t TimeNormal) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan value time.Time
func (t *TimeNormal) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = TimeNormal{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

// 获取本周几的开始时间

func GetThisWeekStartTime(start string) time.Time {
	// 分割时间
	arr := strings.Split(start, "-")
	if len(arr) < 2 {
		return time.Time{}
	}
	// 星期几
	day := StringToInt(arr[0])
	// 时分
	timeArr := strings.Split(arr[1], ":")
	if len(timeArr) < 2 {
		return time.Time{}
	}
	hour := StringToInt(timeArr[0])
	min := StringToInt(timeArr[1])
	if day > 7 || day < 1 {
		return time.Time{}
	}
	now := time.Now()
	nowWeek := now.Weekday()
	if nowWeek == 0 {
		nowWeek = 7
	}
	offset := day - int(nowWeek)
	weekStartDate := time.Date(now.Year(), now.Month(), now.Day(), hour, min, 0, 0, time.Local).AddDate(0, 0, offset)
	return weekStartDate
}
