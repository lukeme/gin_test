package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化Gin引擎
	r := gin.Default()

	// 定义一个GET路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	// 定义一个POST路由
	r.POST("/submit", func(c *gin.Context) {
		type RequestBody struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}
		var requestBody RequestBody
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"name": requestBody.Name,
			"age":  requestBody.Age,
		})
	})

	// 启动Gin服务器，默认在0.0.0.0:8080
	r.Run()
}
