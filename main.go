package main

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	router.GET("/index.htm", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	router.Run() // listen and serve on 0.0.0.0:8080
}
