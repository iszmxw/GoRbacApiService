package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"log"
)

// DB gorm.DB 对象
var DB *gorm.DB

// ConnectDB 初始化模型
func ConnectDB() *gorm.DB {
	var err error

	// 初始化 MySQL 连接信息
	var (
		host     = "192.168.26.139"
		port     = "3306"
		database = "gorbac"
		username = "root"
		password = "root"
		charset  = "utf8mb4"
	)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s", username, password, host, port, database, charset, true, "Local")

	gormConfig := mysql.New(mysql.Config{
		DSN: dsn,
	})

	var level gormlogger.LogLevel
	// 读取不到数据也会显示
	level = gormlogger.Warn
	// 准备数据库连接池
	DB, err = gorm.Open(gormConfig, &gorm.Config{
		Logger: gormlogger.Default.LogMode(level),
	})
	if err != nil {
		log.Fatal(err)
	}
	return DB
}
