package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	router := r.Group("/api")
	{
		router.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "quickstart-v2",
			})
		})

		router.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong-v2",
			})
		})
	}

	r.Run(":8080")
}
