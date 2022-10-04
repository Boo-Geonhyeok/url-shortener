package main

import (
	"github.com/Boo-Geonhyeok/url/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/", handlers.GenerateURL)
	r.GET("/r/:url", handlers.Redirect)
	r.Run()

}
