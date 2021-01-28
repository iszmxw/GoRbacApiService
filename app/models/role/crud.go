package role

import "gorbac/pkg/mysql"

// Create 创建角色，通过 Role.ID 来判断是否创建成功
func Create(a Role) (err error) {
	if err = mysql.DB.Create(&a).Error; err != nil {
		return err
	}
	return nil
}

// GetOne 获取一条数据
func GetOne(where map[string]interface{}) (Role, error) {
	var role Role
	if err := mysql.DB.Where(where).First(&role).Error; err != nil {
		return role, err
	}
	return role, nil

}
