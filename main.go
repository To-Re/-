package main

import (
	"bishe/backend/method"

	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(Cors())

	// 注册
	router.POST("/register", method.Register)
	// 登录
	router.POST("/login", method.Login)

	api := router.Group("/api")
	{
		api.Use(method.AccessTokenMiddleware())

		// api/teacher
		teacher := api.Group("/teacher")
		{
			teacher.GET("/", func(c *gin.Context) {
				c.String(http.StatusOK, "teacher")
			})
		}

		// api/student
		student := api.Group("/student")
		{
			student.GET("/", func(c *gin.Context) {
				c.String(http.StatusOK, "student")
			})
		}
	}

	router.Run(":8080")
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "*")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
