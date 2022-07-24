package response

import (
	"gorm.io/gorm"
	"time"
)

type JsonRole struct { // 格式化返回角色列表
	Id        uint64         `json:"id"`
	Name      string         `json:"name"`
	Routes    string         `json:"routes"`
	Desc      string         `json:"desc"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type JsonRoleDetail struct { // 格式化返回角色详情
	AllRouteList   []AllRouteList `json:"all_route_list"`
	DefaultChecked interface{}    `json:"defaultChecked"`
}

type AllRouteList struct {
	Id       uint64         `json:"id"`
	IsMenu   string         `json:"is_menu"`
	Name     string         `json:"name"`
	ParentId int            `json:"parent_id"`
	Disabled bool           `json:"disabled"`
	Children []AllRouteList `json:"children" gorm:"-"`
}
