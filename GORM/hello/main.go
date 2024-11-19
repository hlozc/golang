package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Database source name
	dsn := "root:123456@tcp(127.0.0.1:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"
	// gorm.Open() 用于创建一个连接 mysql 的实例
	// mysql.Open(dsn) 告诉 GORM 使用 dsn 中的配置信息来连接数据库
	gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
