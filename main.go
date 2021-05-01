package main

import (
	"bishe/backend/config"
	"bishe/backend/dal"
	klassMethod "bishe/backend/method/klass"
	questionMethod "bishe/backend/method/question"
	userMethod "bishe/backend/method/user"
	"bishe/backend/util"
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	/* 配置 */
	config := config.GetDbInfo()
	/* 连数据库 */
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username, config.Password, config.Host, config.Port, config.Database,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("db init panic")
	}
	dal.NewDal(db)

	/* 路由 */
	router := gin.Default()
	router.Use(Cors())

	// 注册
	router.POST("/user/register", userMethod.Register)
	// 登录
	router.POST("/user/login", userMethod.Login)
	user := router.Group("/user")
	{
		user.Use(util.AccessTokenMiddleware())
		user.GET("/info", userMethod.UserInfo)
		user.POST("/logout", userMethod.Logout)
	}

	api := router.Group("/api")
	{
		api.Use(util.AccessTokenMiddleware())

		// api/teacher
		teacher := api.Group("/teacher")
		{
			teacher.Use(util.CheckTeacherAuth())
			klass := teacher.Group("/klass")
			{
				klass.GET("/list", klassMethod.KlassList)
				klass.POST("/create", klassMethod.KlassCreate)
				klass.POST("/update", klassMethod.KlassUpdate)
				klass.GET("/detail", klassMethod.KlassDetail)
			}
			question := teacher.Group("question")
			{
				question.GET("/list", questionMethod.QuestionList)
				question.POST("/create", questionMethod.QuestionCreate)
			}
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
