package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		he, _ := c.Cookie("hello")
		fmt.Println(he)

		c.SetCookie("hello", "world", 60*60*365*24, "", "", false, false)
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	router.Run()
}
