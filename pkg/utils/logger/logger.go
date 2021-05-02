package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

var RequestId string

// IsExist 检测文件是否存在
func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}

// Format 格式化日志
func Format(content interface{}) string {
	Time := time.Now().Format("2006-01-02 15:04:05")
	if len(RequestId) <= 0 {
		log.Fatalln("当前请求ID不存在" + RequestId)
	}
	formatInfo := "[" + Time + "]" + "【当前请求ID】" + RequestId + "【" + fmt.Sprintf("%+v", content) + "】"
	return formatInfo
}

// Write 写日志方法
func Write(content interface{}) {
	info := Format(content)
	// 获取当天时间
	today := time.Now().Format("2006-01-02")
	logName := "logs/" + string(today) + ".log"

	// 不存在就创建日志文件
	if !IsExist(logName) {
		_, err := os.Create(logName) // 创建日志文件
		if err != nil {
			// 打印日志 并退出程序
			log.Fatalln("fail to create " + logName + " file!")
		}
	}

	file, err := os.OpenFile(logName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		// 打印日志 并退出程序
		log.Fatalln("fail to open " + logName + " file!")
	}

	// 创建logger对象　这种方式会显示触发日志文件行数
	log.SetFlags(0) // 去除时间前缀
	loggerFile := log.New(file, "", log.LstdFlags|log.Llongfile)
	// 打印日志内容到控制台
	//log.Println(info)
	fmt.Println(info)
	// 打印日志内容到文件
	loggerFile.Println(info)
}

// LogError 当存在错误时记录日志
func LogError(err error) {
	if err != nil {
		Write(err)
	}
}

// LogInfo 记录日志
func LogInfo(info interface{}) {
	Write(info)
}
