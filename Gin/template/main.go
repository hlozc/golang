package main

import (
	"encoding/json"
	"fmt"
	"net/http"

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

// 自定义 Go 中间件 拦截器
func Interrupter() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 通过自定义的中间件，设置的值，在后续处理只要调用了这个中间件的都可以拿到这里的参数
		c.Set("usersession", "userid-1")
		// 如果通过了就放行
		c.Next()
		// 否则就阻断
		c.Abort()
	}
}

func main() {
	r := gin.Default()

	// 这里添加全局中间件
	r.Use()
	// 加载静态页面
	r.LoadHTMLGlob("../tpl/*")
	// 加载静态资源
	r.Static("/static", "../static")

	r.GET("/index", staticIndex)
	// 加个 冒号 就可以将这里的 userid 和 username 提取出来了
	// 通过 Query 获取查询字符串里面的数据
	r.GET("/user/info/:userid/:username", func(c *gin.Context) {
		userid := c.Param("userid")
		username := c.Param("username")
		page := c.Query("page")
		c.JSON(200, gin.H{
			"userid":   userid,
			"username": username,
			"page":     page,
		})
	})
	// 获取前端传递过来的 json 数据
	r.POST("/json_req", func(c *gin.Context) {
		jsonData, _ := c.GetRawData()
		c.JSON(200, jsonData)

		// 反序列化
		var m map[string]interface{}
		_ = json.Unmarshal(jsonData, &m)
	})
	r.POST("/user/form", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		c.JSON(200, gin.H{
			"username": username,
			"password": password,
		})
	})
	// 路由 重定向
	r.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})
	// 404, 返回一个 404 页面
	r.NoRoute(func(ctx *gin.Context) {

	})

	// 路由组
	userGroup := r.Group("/user")
	{
		userGroup.GET("/add")
		userGroup.GET("/login")
	}

	orderGroup := r.Group("/order")
	{
		orderGroup.GET("/add")
		orderGroup.GET("/query")
	}

	// 通过 拦截器 ，去获取 中间件 里面的值
	r.GET("/user/info", Interrupter(), func(c *gin.Context) {
		if _, ok := c.MustGet("usersession").(string); ok {
			fmt.Println("拦截器拦截到的值")
		}
	})

	// 文件上传类问题
	r.POST("upload", func(c *gin.Context) {
		// 从请求和 前端 字段中获取这个文件
		f, _ := c.FormFile("f1")
		// 保存文件
		c.SaveUploadedFile(f, "./")
	})

	// 多个文件的保存类问题
	r.POST("/uploads", func(c *gin.Context) {
		// 获取解析之后的表单
		form, _ := c.MultipartForm()
		// 获取 form 表单里面的文件对象
		files := form.File["file"]

		for _, file := range files {
			c.SaveUploadedFile(file, "./")
		}
	})

	r.Run()
}
