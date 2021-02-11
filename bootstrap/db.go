package bootstrap

import (
	"gorbac/app/models/account"
	"gorbac/app/models/login_log"
	"gorbac/app/models/operation_log"
	"gorbac/app/models/role"
	"gorbac/app/models/routes"
	"gorbac/pkg/config"
	"gorbac/pkg/mysql"
	"gorm.io/gorm"
	"time"
)

// SetupDB 初始化数据库和 ORM
func SetupDB() {
	// 建立数据库连接池
	db := mysql.ConnectDB()

	// 命令行打印数据库请求的信息
	sqlDB, _ := db.DB()

	// 设置最大连接数
	sqlDB.SetMaxOpenConns(config.GetInt("database.mysql.max_open_connections"))
	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(config.GetInt("database.mysql.max_idle_connections"))
	// 设置每个链接的过期时间
	sqlDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.mysql.max_life_seconds")) * time.Second)
	// 创建和维护数据表结构
	migration(db)
}

// 自动迁移
func migration(db *gorm.DB) {
	db.AutoMigrate(
		// Account 管理员表
		&account.Account{},
		// LoginLog 登录日志表
		&login_log.LoginLog{},
		// OperationLog 操作日志
		&operation_log.OperationLog{},
		// Role 角色表
		&role.Role{},
		// Route 菜单路由节点
		&routes.Routes{},
	)
}
