package account

import (
	"gorbac/app/models"
	"gorm.io/gorm"
)

// Account 管理员表
type Account struct {
	gorm.Model
	Username string `gorm:"type:varchar(50);not null;comment '用户名'"`
	Password string `gorm:"type:varchar(255);comment '密码'" valid:"password"`
	Level    int    `gorm:"type:int(2);default 2;comment '账号类型，1为管理员，2为商户'"`
	RoleId   int    `gorm:"type:int(11);not null;default 2;comment '角色id'"`
	Mobile   string `gorm:"type:varchar(14);default null;comment '手机号码'"`
	Status   int    `gorm:"type:int(2);not null;default 1;comment '状态1：为正常 -1：为冻结'"`
}

// Create 创建账户，通过 User.ID 来判断是否创建成功
func (account *Account) Create() (err error) {
	if err = models.DB.Create(&account).Error; err != nil {
		return err
	}
	return nil
}
