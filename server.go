package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	web := gin.Default()

	web.Static("/static", "./assets")
	web.StaticFile("/", "./assets/index.html")

	// web.LoadHTMLGlob("templates/*")
	// web.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.tmpl", gin.H{
	// 		"title": "Paper",
	// 	})
	// })

	web.Run(":3000")
}
