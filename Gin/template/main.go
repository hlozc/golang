package main

import (
	"github.com/gin-gonic/gin"
)

func staticIndex(c *gin.Context) {
	c.HTML(200, "index.tpl", gin.H{
		"Title":   "Welcome to Gin",
		"Header":  "This is a header",
		"Content": "This is a sample page generated using Gin and tmpl.",
		"Footer":  "© 2024 Your Project",
	})
}

func main() {
	r := gin.Default()
	// 加载 template 模板的目录, 或者说将指定的 template 文件都加载进来
	r.LoadHTMLGlob("../tpl/*")

	r.GET("/index", staticIndex)
	r.Run()
}
