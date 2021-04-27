package main

import (
	"bishe/backend/method"

	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	method.Login()
	router := gin.Default()

	router.POST("/login", func(c *gin.Context) {
		c.String(http.StatusOK, "It works")
	})
	api := router.Group("/api")
	api.Use(method.AccessTokenMiddleware())
	api.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "It works")
	})

	router.Run(":8080")
}

/*
curl -H 'AuthToken: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTk2MTgwMjUsImlhdCI6MTYxOTUzMTYyNSwidXNlcl9pZCI6MSwid2hvIjoxfQ.0oPXhRNGbYRia0216iMzz0saBJgEfPDIeh-CfF6OW4U' 127.0.0.1:8080/api/
*/
