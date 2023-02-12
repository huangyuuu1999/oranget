package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/huangyuuu1999/oranget/src/package1"
	"github.com/huangyuuu1999/oranget/src/router"
	"net/http"
	"time"
)

type HeroInfo struct {
	Name        string
	Age         int
	Attack      int
	Wise        int
	DynamicTags []string
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
				HeroInfo{"朱元璋", 34, 76, 86, []string{"和尚", "陈友谅", "岳父🐂"}},
				HeroInfo{"宋江", 45, 65, 82, []string{"黑厮", "及时雨"}},
				HeroInfo{"洪秀全", 32, 89, 56, []string{"农民", "太平天国"}},
				HeroInfo{"项羽", 28, 95, 75, []string{"大力举鼎", "学万人敌"}},
				HeroInfo{"张全蛋", 26, 56, 80, []string{"foxxxx"}},
				HeroInfo{"陈二", 28, 99, 65, []string{"打工仔", "小镇做题家"}},
				HeroInfo{"吴佩孚", 46, 81, 79, []string{"军阀", "奉系", "孙传芳"}},
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
	r.POST("/api/sites", package1.PostSite)

	r.POST("/api/bind", package1.TestBindJSON)
	r.DELETE("/api/sites", package1.DeleteSite)

	r = router.LoadRouter(r)

	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080
}
