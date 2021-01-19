package models

import (
	"gorbac/pkg/mysql"
)

// Account 管理员表
type Account struct {
	BaseModel
	Username string `gorm:"type:varchar(50);not null;comment:用户名" json:"username" validate:"username"`
	Password string `gorm:"type:varchar(255);comment:密码" json:"password" validate:"password"`
	Level    int    `gorm:"type:int(2);default:2;comment:账号类型，1为管理员，2为商户" json:"level"`
	RoleId   int    `gorm:"type:int(11);not null;default:2;comment:(角色id)" json:"role_id"`
	Mobile   string `gorm:"type:varchar(14);comment:手机号码" json:"mobile"`
	Status   int    `gorm:"type:int(2);not null;default:1;comment:状态1：为正常 -1：为冻结" json:"status"`
	BaseModelLast
}

// Create 创建账户，通过 User.ID 来判断是否创建成功
func Create(a Account) (err error) {
	if err = mysql.DB.Create(&a).Error; err != nil {
		return err
	}
	return nil
}

// GetOne 获取一条数据
func GetOne(where map[string]interface{}) (Account, error) {
	var account Account
	if err := mysql.DB.Where(where).First(&account).Error; err != nil {
		return account, err
	}
	return account, nil

}
