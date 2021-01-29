package route

import (
	"gorbac/app/models"
	"gorbac/pkg/config"
)

// Route 菜单路由节点
type Route struct {
	models.BaseModel
	IsMenu   string `gorm:"type:int(1);not null;default:1;comment:是否根菜单1-是 0-否" json:"is_menu"`
	Route    string `gorm:"type:varchar(255);comment:访问路由地址" json:"route"`
	Name     string `gorm:"type:varchar(100);not null;comment:路由名称" json:"name"`
	ParentId string `gorm:"type:int(11);not null;default:0;comment:上级id" json:"parent_id"`
	models.BaseModelLast
}

func (Route) TableName() string {
	prefix := config.GetString("database.mysql.prefix")
	table := "route"
	return prefix + table
}
