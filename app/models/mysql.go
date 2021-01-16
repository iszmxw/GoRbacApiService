package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"log"
	"time"
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

// SetupDB 初始化数据库和 ORM
func SetupDB() {
	// 建立数据库连接池
	db := ConnectDB()

	// 命令行打印数据库请求的信息
	sqlDB, _ := db.DB()

	// 设置最大连接数
	sqlDB.SetMaxOpenConns(25)
	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(100)
	// 设置每个链接的过期时间
	sqlDB.SetConnMaxLifetime(time.Duration(5*60) * time.Second)
	// 创建和维护数据表结构
	migration(db)
}

func migration(db *gorm.DB) {
	// 自动迁移
	db.AutoMigrate(
		&Account{},
	)
}
