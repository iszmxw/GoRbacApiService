package account

import (
	"gorbac/app/models"
	"gorbac/pkg/config"
	"gorbac/pkg/utils/times"
	"gorm.io/gorm"
)

func (m *Account) TableName() string {
	prefix := config.GetString("database.mysql.prefix")
	table := "account"
	return prefix + table
}

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

// json 响应结构体定义，供查询数据引用

// JsonAccount 格式化返回登录日志
type JsonAccount struct {
	Id        uint64         `json:"id"`
	Username  string         `json:"username"`
	Level     int            `json:"level"`
	RoleId    int            `json:"role_id"`
	RoleName  string         `json:"role_name"`
	Mobile    string         `json:"mobile"`
	Status    int            `json:"status"`
	CreatedAt times.MyTime   `json:"created_at"`
	UpdatedAt times.MyTime   `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
