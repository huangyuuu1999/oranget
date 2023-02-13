package main

import (
	"github.com/gin-gonic/gin"
	"github.com/huangyuuu1999/oranget/src/package1"
	"github.com/huangyuuu1999/oranget/src/router"
	"time"
)

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
	r.GET("/api/info", package1.GetHero)
	r.GET("/api/sites", package1.GetSites)
	r.POST("/api/sites", package1.PostSite)

	r.POST("/api/bind", package1.TestBindJSON)
	r.DELETE("/api/sites", package1.DeleteSite)

	r = router.LoadRouter(r)

	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080
}
