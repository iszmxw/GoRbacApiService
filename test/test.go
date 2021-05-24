package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

// IsExist 检测文件是否存在
func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}

// Write 写日志方法
func Write(i int) {
	num := fmt.Sprintf("%+v", i)
	content := "第【" + num + "】次   打印日志快乐时光"

	// 打印日志内容到控制台
	fmt.Println(content)

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
	// 打印日志内容到文件
	loggerFile.Println(content)
}

func main() {

	ch := make(chan int)

	fmt.Println(ch)

	go func() {
		var sum = 0
		for i := 0; i < 10; i++ {
			sum += i
		}
		ch <- sum
		ch <- 123
	}()
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}
