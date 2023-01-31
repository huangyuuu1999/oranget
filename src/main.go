package main

import (
	"time"
	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		now := time.Now()
		c.JSON(200, gin.H{
			"message": "pong",
			"name": "gin-webframework",
			"usage": "backend for orangear",
			"time": now.Format("2006-01-02 15:04:05.000 Mon Jan"),
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
