package account

import "gorbac/pkg/mysql"

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
