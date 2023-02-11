package router

import (
	"github.com/gin-gonic/gin"
	"github.com/huangyuuu1999/oranget/src/package1"
)

func LoadRouter(r *gin.Engine) *gin.Engine {
	r.GET("/api/login", func(c *gin.Context) {
		c.String(200, "everything is ok now!")
	})
	r.PUT("/api/sites", package1.UpdateSite)

	return r
}
