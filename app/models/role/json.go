package role

import (
	"gorm.io/gorm"
	"time"
)

// json 响应结构体定义，供查询数据引用

// 格式化返回角色列表
type JsonRole struct {
	Id        uint64         `json:"id"`
	Name      string         `json:"name"`
	Routes    string         `json:"routes"`
	Desc      string         `json:"desc"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// 格式化返回角色详情
type JsonRoleDetail struct {
	AllRouteList   []AllRouteList `json:"all_route_list"`
	DefaultChecked interface{}    `json:"defaultChecked"`
}

type AllRouteList struct {
	Id       uint64         `json:"id"`
	IsMenu   string         `json:"is_menu"`
	Name     string         `json:"name"`
	ParentId int            `json:"parent_id"`
	Disabled bool           `json:"disabled"`
	Children interface{} `json:"children"`
}
