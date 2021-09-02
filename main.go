package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/gin-hello-world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world!",
		})
	})
	router.Run(":8180")

}
