package helpers

// 转换器
import (
	"bytes"
	"encoding/json"
	"gorbac/pkg/logger"
	"strconv"
	"strings"
)

// Int64ToString 将 int64 转换为 string
func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

// IntToString 整型转字符串
func IntToString(num int) string {
	return strconv.Itoa(num)
}

// Uint64ToString 将 uint64 转换为 string
func Uint64ToString(num uint64) string {
	return strconv.FormatUint(num, 10)
}

// StringToInt 将字符串转换为 int
func StringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		logger.Error(err)
		return 0
	}
	return i
}

// StringToInt64 将字符串转换为 int64
func StringToInt64(str string) int64 {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		//logger.Error(err)
		return 0
	}
	return i
}

// Uint2String 将字符串转换为 int
func Uint2String(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		logger.Error(err)
	}
	return i
}

// Struct2json Struct转换json
func Struct2json(value interface{}) string {
	bs, _ := json.Marshal(value)
	var out bytes.Buffer
	err := json.Indent(&out, bs, "", "\t")
	if err != nil {
		logger.Error(err)
	}
	return out.String()
}

// 字符串转换float64，空字符串返回 0

func StrToFloat64(str string) float64 {
	if len(str) <= 0 {
		return 0
	}
	f64, Err := strconv.ParseFloat(str, 64)
	if Err != nil {
		logger.Error(Err)
		f64 = 0
	}
	return f64
}

// 转换

func CloseDataConvert(close string) []float64 {
	var floatList []float64
	close = strings.Replace(close, "c_close", "", -1)
	close = strings.Replace(close, "\n", ",", -1)
	close = strings.Trim(close, ",")
	strList := strings.Split(close, ",")
	for _, item := range strList {
		if item == "" {
			continue
		}
		floatList = append(floatList, StrToFloat64(item))
	}
	return floatList
}
