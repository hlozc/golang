package main

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	ID        int       `gorm:"primaryKey"` // 声明为主键
	UserName  string    `gorm:"not null;unique;column:username"`
	Email     string    `gorm:"not null;unique;column:email"`
	Password  string    `gorm:"not null;column:password"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func initGORM() *gorm.DB {
	// Database source name
	dsn := "root:123456@tcp(127.0.0.1:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"
	// gorm.Open() 用于创建一个连接 mysql 的实例
	// mysql.Open(dsn) 用来创建一个驱动实例，表示要连接的对象是什么，MySQL, PostgreSQL
	// &gorm.Config{} 表示 gorm 的全局配置，比如日志级别，默认事务等
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("DB connect fail")
		os.Exit(0)
	}

	// 自动针对 User 结构体建表，默认是 users
	db.AutoMigrate(&User{})
	return db
}

// Take, First, Last 都是基于 Id 来进行搜索的
func selectTest() {
	// var usr User
	var users []*User
	// 1.
	// db.Where("UserName", "张三").Take(&usr)

	// 2.
	// db.Where("Username = ?", "张三").Take(&usr)

	// 3. Model() 接口用来选择一个 Model，同时也可以在里面限制条件
	// db.Model(&User{}).Where("UserName != ?", "张三").Find(&users)

	// 4.
	// db.First(&usr)
	// db.Take(&usr)

	// 5. Hash
	// db.Where(map[string]interface{}{
	// 	"UserName": "张三",
	// 	"ID":       1,
	// }).Take(&usr)

	// 6.
	db.Find(&users, "id > ? and username != ?", 1, "张三")

	for i := 0; i < len(users); i++ {
		fmt.Println(users[i])
	}
}

func main() {
	db = initGORM()
	selectTest()
}
