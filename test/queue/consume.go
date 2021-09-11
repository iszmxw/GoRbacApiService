package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)
// consume
const (
	testUrl           = "https://blog.54zm.com/mp3"
	proUrl            = "https://wxapp.yidaogz.cn/app/index.php?c=entry&a=wxapp&i=2&m=lingchi_b11&do=LuckSponsor&action=lottery_redis"
	routineCountTotal = 10 // 限制线程数
)

func main() {
	for true {
		syncConsume(routineCountTotal)
		fmt.Println("异步消费10条结束")
	}
}

func httpGet(now string) {
	resp, err := http.Get(proUrl)
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
