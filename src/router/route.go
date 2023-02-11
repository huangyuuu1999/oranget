package router

import "github.com/gin-gonic/gin"

func LoadRouter(r *gin.Engine) *gin.Engine {
	r.GET("/api/login", func(c *gin.Context) {
		c.String(200, "everything is ok now!")
	})
	return r
}
