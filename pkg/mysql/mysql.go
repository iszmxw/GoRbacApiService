package mysql

import (
	"fmt"
	"gorbac/pkg/config"
	MysqlLog "gorbac/pkg/logger"
	"gorbac/pkg/logger/zapgorm2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

// DB gorm.DB 对象
var DB *gorm.DB

// ConnectDB 初始化模型
func ConnectDB() *gorm.DB {
	var err error

	// 初始化 MySQL 连接信息
	var (
		host     = config.GetString("database.mysql.host")
		port     = config.GetString("database.mysql.port")
		database = config.GetString("database.mysql.database")
		username = config.GetString("database.mysql.username")
		password = config.GetString("database.mysql.password")
		prefix   = config.GetString("database.mysql.prefix")
		charset  = config.GetString("database.mysql.charset")
	)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s", username, password, host, port, database, charset, true, "Local")

	gormConfig := mysql.New(mysql.Config{
		DSN: dsn,
	})

	// 追踪mysql日志
	logger := zapgorm2.New(MysqlLog.Logger)
	logger.SetAsDefault()
	// 准备数据库连接池
	DB, err = gorm.Open(gormConfig, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   prefix, // 表名前缀，`User` 的表名应该是 `go_users`
			SingularTable: true,   // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `go_user`
		},
		Logger: logger,
	})
	if err != nil {
		log.Fatal(err)
	}
	return DB
}
