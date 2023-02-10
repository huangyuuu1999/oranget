package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/huangyuuu1999/oranget/src/package1"
	"net/http"
	"time"
)

type HeroInfo struct {
	Name   string
	Age    int
	Attack int
	Wise   int
}

func main() {
	r := gin.Default()
	r.GET("/api/ping", func(c *gin.Context) {
		now := time.Now()
		c.JSON(200, gin.H{
			"messAge": "pong",
			"Name":    "gin-webframework",
			"usAge":   "backend for orangear",
			"time":    now.Format("2006-01-02 15:04:05.000 Mon Jan"),
		})
	})
	r.GET("/api/info", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"info": []interface{}{
				HeroInfo{"朱元璋", 34, 76, 86},
				HeroInfo{"宋江", 45, 65, 82},
				HeroInfo{"洪秀全", 32, 89, 56},
				HeroInfo{"项羽", 28, 95, 75},
				HeroInfo{"张全蛋", 26, 56, 80},
				HeroInfo{"陈二", 28, 99, 65},
				HeroInfo{"吴佩孚", 46, 81, 79},
			},
		})
	})
	r.GET("/api/sites", func(c *gin.Context) {
		sites := package1.GetAllSites()
		fmt.Println(sites)
		c.JSON(http.StatusOK, gin.H{
			"info": sites,
		})
	})
	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080
}
