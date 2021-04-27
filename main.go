package main

import (
	"bishe/backend/method"

	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

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
