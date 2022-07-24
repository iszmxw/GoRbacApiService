package response

import (
	"gorm.io/gorm"
	"time"
)

type JsonRoute struct { // 格式化返回菜单路由节点
	Id        uint64         `json:"id"`
	Name      string         `json:"name"`
	Routes    string         `json:"routes"`
	Desc      string         `json:"desc"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type JsonRouteTree struct { // 格式化返回菜单路由树
	Id        uint64          `json:"id"`
	Sort      int             `json:"sort"`
	Type      string          `json:"type"`
	IsMenu    int             `json:"is_menu"`
	Route     string          `json:"route"`
	Component string          `json:"component"`
	Name      string          `json:"name"`
	Icon      string          `json:"icon"`
	ParentId  int             `json:"parent_id"`
	CreateBy  int             `json:"create_by"`
	Status    int             `json:"status"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt gorm.DeletedAt  `json:"deleted_at"`
	Children  []JsonRouteTree `json:"children" gorm:"-"`
}
