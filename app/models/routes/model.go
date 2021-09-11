package routes

import (
	"gorbac/app/models"
	"gorbac/pkg/config"
	"gorm.io/gorm"
	"time"
)

func (m *Routes) TableName() string {
	prefix := config.GetString("database.mysql.prefix")
	table := "routes"
	return prefix + table
}

// Routes 菜单路由节点
type Routes struct {
	models.BaseModel
	Sort      int    `gorm:"type:int(11);default:0;comment:排序" json:"sort"`
	Type      string `gorm:"type:varchar(10);not null;default:page;comment:page-页面   api-接口" json:"type"`
	IsMenu    int    `gorm:"type:int(1);not null;default:1;comment:是否根菜单1-是 0-否" json:"is_menu"`
	Route     string `gorm:"type:varchar(255);comment:访问路由地址" json:"route" validate:"route"`
	Component string `gorm:"type:varchar(255);comment:页面组件地址" json:"component" validate:"component"`
	Name      string `gorm:"type:varchar(100);not null;comment:路由名称" json:"name" validate:"name"`
	Icon      string `gorm:"type:varchar(255);comment:icon图标" json:"icon"`
	ParentId  int    `gorm:"type:int(11);not null;default:0;comment:上级id" json:"parent_id" validate:"parent_id"`
	CreateBy  int    `gorm:"type:int(11);comment:创建者" json:"create_by"`
	Status    int    `gorm:"type:char(1);comment:1-已启用   0-未启用" json:"status"`
	models.BaseModelLast
}

// json 响应结构体定义，供查询数据引用

// JsonRoute 格式化返回菜单路由节点
type JsonRoute struct {
	Id        uint64         `json:"id"`
	Name      string         `json:"name"`
	Routes    string         `json:"routes"`
	Desc      string         `json:"desc"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// JsonRouteTree 格式化返回菜单路由树
type JsonRouteTree struct {
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
