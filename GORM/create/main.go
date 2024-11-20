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

// 增
func CreateUser(user *User) {
	tx := db.Create(user)

	// ID 会被设置进去，还有是否有错误码，以及受影响行数
	fmt.Println(user.ID, tx.Error, tx.RowsAffected)
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

func create(usr *User) {
	res := db.Create(usr)
	fmt.Println("Primary Key: ", usr.ID)
	fmt.Println("Error Info: ", res.Error)
	fmt.Println("Rows Affected Num: ", res.RowsAffected)
}

func createMulti(usr []*User) {
	res := db.Create(usr)

	fmt.Println("Error Info: ", res.Error)
	fmt.Println("Rows Affected Num: ", res.RowsAffected)
}

func createByHash(m map[string]interface{}) {
	db.Model(&User{}).Create(m)
}

func main() {
	db = initGORM()
	// 一次插入一个
	// usr := &User{
	// 	UserName: "张三",
	// 	Email:    "123@qq.com",
	// 	Password: "123",
	// }
	// create(usr)

	// // 一次插入多个
	// createMulti([]*User{
	// 	{
	// 		UserName: "李四",
	// 		Email:    "456@qq.com",
	// 		Password: "456",
	// 	},
	// 	{
	// 		UserName: "王五",
	// 		Email:    "789@qq.com",
	// 		Password: "789",
	// 	},
	// })

	createByHash(map[string]interface{}{
		"UserName": "赵六",
		"Email":    "000@qq.com",
		"Password": "000",
	})
}
