package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

// LogError 当存在错误时记录日志
func LogError(err error) {
	if err != nil {
		LogWrite(err)
	}
}

func SystemError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

// LogInfo 记录日志
func LogInfo(info interface{}) {
	LogWrite(info)
}

// IsExist 检测文件是否存在
func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}

// LogWrite 写日志方法
func LogWrite(info interface{}) {
	info = fmt.Sprintf("%+v", info)
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
	loggerFile := log.New(file, "", log.LstdFlags|log.Llongfile)
	loggerFile.SetFlags(log.LstdFlags) // 设置每行日志打印格式  仅显示时间
	log.Println("start")
	// 打印日志内容到控制台
	log.Println(info)
	fmt.Println(info)
	// 打印日志内容到文件
	loggerFile.Println(info)
	log.Println("end")
}
