package main

import (
	"encoding/xml"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/gin-hello-world-json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world!",
		})
	})
	router.GET("/gin-hello-world-xml", func(c *gin.Context) {
		c.XML(200, Message{
			MessageText: "Hello World!",
		})
	})

	router.Run(":8180")

}

type Message struct {
	XMLName     xml.Name `xml:"message"`
	MessageText string   `xml:"messageText,attr"`
}
