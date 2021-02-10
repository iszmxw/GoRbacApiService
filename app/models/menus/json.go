package menus

import (
	"gorm.io/gorm"
	"time"
)

// json 响应结构体定义，供查询数据引用

// 格式化返回菜单路由节点
type JsonTreeMenus struct {
	Id        uint64          `json:"id"`
	Sort      int             `json:"sort"`
	Type      string          `json:"type"`
	IsMenu    int             `json:"is_menu"`
	Route     string          `json:"route"`
	Name      string          `json:"name"`
	Icon      string          `json:"icon"`
	ParentId  int             `json:"parent_id"`
	CreateBy  int             `json:"create_by"`
	Status    int             `json:"status"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt gorm.DeletedAt  `json:"deleted_at"`
	Children  []JsonTreeMenus `json:"children" gorm:"-"`
}
