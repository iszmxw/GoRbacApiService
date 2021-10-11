package main

import (
	"fmt"
	"gorbac/pkg/validator"
	"reflect"
)

func main() {
	pageInfo := PageInfo{Page: 1000, PageSize: 1000}
	t := reflect.TypeOf(pageInfo)
	fmt.Println(t)
	fmt.Println(t)
	fmt.Println(t)
	fmt.Println(t)
	fmt.Println(t)
	err := service.Validate.Struct(pageInfo) //注意导入validator包
	if err != nil {
		r := service.Translate(err, pageInfo)
		fmt.Println(Errors{
			Code:   -1,
			Msg:    "请求参数验证失败",
			Errors: r,
		})
		return
	}
}

//PageInfo 分页请求数据结构
type PageInfo struct {
	Page     int `json:"page"  validate:"required,min=1" label:"页码"`
	PageSize int `json:"pageSize" validate:"required,max=100" label:"每页大小"`
}

//Errors 错误响应
type Errors struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Errors interface{} `json:"errors"`
}
