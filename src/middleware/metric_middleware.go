package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func CostTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		nowTime := time.Now()

		c.Next()

		costTime := time.Since(nowTime)

		url := c.Request.URL.String()
		fmt.Printf("[CUSTOM-LOG]the request URL %s cost %v\n", url, costTime)
	}
}
