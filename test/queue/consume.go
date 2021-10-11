package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

// consume
var (
	Site              string // 站点ID
	Url               string // 站点Url
	RoutineCountTotal string // 限制线程数
)

func help() {
	fmt.Println("======================help=========================")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Tips：该脚本为异步消费工具...")
	fmt.Println("Please：请输入参数：")
	fmt.Println("Eg：./consume 2 10 ")
	fmt.Println("")
	fmt.Println("以上表示启动站点 uniacid=2 的消费进程，开启十条")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("======================help=========================")
}

func Args() error {
	args := os.Args
	if len(args) < 2 || args == nil || args[1] == "" || args[2] == "" {
		help()
		return errors.New("请检查参数")
	}
	Site = args[1]
	RoutineCountTotal = args[2]
	Url = "https://wxapp.yidaogz.cn/app/index.php?c=entry&a=wxapp&i=" + Site + "&m=lingchi_b11&do=LuckSponsor&action=lottery_redis"
	fmt.Printf("启动的站点id为：%v\n", Site)
	fmt.Printf("消费的Url地址为：%v\n", Url)
	fmt.Printf("开启线程数为：%v\n", RoutineCountTotal)
	return nil
}

func main() {
	if err := Args(); err != nil {
		fmt.Println(err.Error())
		return
	}
	for true {
		num, err := strconv.Atoi(RoutineCountTotal)
		if err != nil {
			fmt.Println("开启线程数异常")
		}
		syncConsume(num)
		fmt.Printf("异步消费%v条结束\n", num)
	}
}

func httpGet(now string) {
	resp, err := http.Get(Url)
	fmt.Printf("当前请求时间：%v\n", now)
	if err != nil {
		// handle error
		fmt.Printf("request err ：%v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("Body.Close() err：%v", err)
		}
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Printf("ioutil.ReadAll err：%v", err)
	}
	fmt.Println(string(body))
}

func syncConsume(num int) {
	// 创建一个阻塞通道
	ch := make(chan bool)
	for i := 1; i <= num; i++ {
		go func(i int) {
			time.Sleep(time.Second)
			now := time.Now().Format("2006-01-02 15:04:05.999999999")
			httpGet(now)
			fmt.Printf("异步消费的第%v条\n", i)
			ch <- true
		}(i)
	}
	for i := 1; i <= num; i++ {
		<-ch
	}
	fmt.Printf("%v条done\n", num)
}
