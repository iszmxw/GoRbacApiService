package role

import (
	"gorbac/app/models"
	"gorbac/pkg/config"
	"gorm.io/gorm"
	"time"
)

func (m *Role) TableName() string {
	prefix := config.GetString("database.mysql.prefix")
	table := "role"
	return prefix + table
}

// Role 角色表
type Role struct {
	models.BaseModel
	Name   string `gorm:"type:varchar(100);not null;comment:角色名称" json:"name"`
	Routes string `gorm:"type:varchar(255);comment:路由id,该角色所具有的路由" json:"routes"`
	Desc   string `gorm:"type:varchar(255);comment:角色描述" json:"desc"`
	models.BaseModelLast
}

// json 响应结构体定义，供查询数据引用

// JsonRole 格式化返回角色列表
type JsonRole struct {
	Id        uint64         `json:"id"`
	Name      string         `json:"name"`
	Routes    string         `json:"routes"`
	Desc      string         `json:"desc"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// JsonRoleDetail 格式化返回角色详情
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
	Children []AllRouteList `json:"children" gorm:"-"`
}
