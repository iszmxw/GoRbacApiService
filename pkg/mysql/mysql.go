package mysql

import (
	"fmt"
	"gorbac/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
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

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: 500 * time.Millisecond, // 慢 SQL 阈值 500ms （如打印所有sql设置1ms）
			LogLevel:      logger.Warn,            // Log level 读取不到数据也会显示
			Colorful:      false,                  // 禁用彩色打印
		},
	)

	// 准备数据库连接池
	DB, err = gorm.Open(gormConfig, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   prefix, // 表名前缀，`User` 的表名应该是 `go_users`
			SingularTable: true,   // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `go_user`
		},
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal(err)
	}
	return DB
}
