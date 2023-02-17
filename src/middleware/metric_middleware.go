package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/huangyuuu1999/oranget/src/common"
	"github.com/huangyuuu1999/oranget/src/model"
)

func CostTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		nowTime := time.Now()

		c.Next()

		costTime := time.Since(nowTime) / 10e3 // μs

		url := c.Request.URL.String()
		ip := c.ClientIP()
		method := c.Request.Method

		record := model.VisitRecord{Time: nowTime.Unix(), Path: url, CostTime: int(costTime), IP: ip, Method: method}
		db := common.InitDB()
		defer func() {
			dbSQL, err := db.DB()
			if err != nil {
				panic(err)
			}
			dbSQL.Close()
		}()
		db.Create(&record)
	}
}
