package account

import (
	"gorbac/app/models"
)

// Account 管理员表
type Account struct {
	models.BaseModel
	Username string `gorm:"type:varchar(50);not null;comment:用户名" json:"username" validate:"username"`
	Password string `gorm:"type:varchar(255);comment:密码" json:"password" validate:"password"`
	Level    int    `gorm:"type:int(2);default:2;comment:账号类型，1为管理员，2为商户" json:"level"`
	RoleId   int    `gorm:"type:int(11);not null;default:2;comment:(角色id)" json:"role_id"`
	Mobile   string `gorm:"type:varchar(14);comment:手机号码" json:"mobile"`
	Status   int    `gorm:"type:int(2);not null;default:1;comment:状态1：为正常 -1：为冻结" json:"status"`
	models.BaseModelLast
}
