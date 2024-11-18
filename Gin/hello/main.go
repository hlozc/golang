package main

import (
	"github.com/gin-gonic/gin"
)

func sayHello(c *gin.Context) {
	// gin.H 就是 map[string]any 的哈希表
	// c.JSON 会将这里序列化之后的结果写入到 Response 中
	c.JSON(200, gin.H{
		"message": "hello gin",
	})
}

func main() {
	r := gin.Default()

	r.GET("/hello", sayHello)

	r.Run()
}
