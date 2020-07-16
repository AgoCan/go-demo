package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": 1,
		})

	})
	r.POST("login", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"he": 1,
		})
	})
	r.Run(":9000")
}
