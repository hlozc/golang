package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func initGORM() *gorm.DB {
	// Database source name
	dsn := "root:123456@tcp(127.0.0.1:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"
	// gorm.Open() 用于创建一个连接 mysql 的实例
	// mysql.Open(dsn) 用来创建一个驱动实例，表示要连接的对象是什么，MySQL, PostgreSQL
	// &gorm.Config{} 表示 gorm 的全局配置，比如日志级别，默认事务等
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err == nil {
		fmt.Println("Connect mysql success")
	}
	return db
}

func main() {
	db = initGORM()
}
