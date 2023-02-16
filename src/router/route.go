package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huangyuuu1999/oranget/src/common"
	"github.com/huangyuuu1999/oranget/src/model"
	"github.com/huangyuuu1999/oranget/src/package1"
)

func LoadRouter(r *gin.Engine) *gin.Engine {
	r.GET("/api/login", func(c *gin.Context) {
		c.String(200, "everything is ok now!")
	})
	r.PUT("/api/sites", package1.UpdateSite)

	r.POST("/api/auth/register", func(c *gin.Context) {
		db := common.InitDB()
		// 获取参数
		username := c.PostForm("username")
		password := c.PostForm("password")
		phone := c.PostForm("phone")

		if len(password) < 6 {
			c.JSON(422, gin.H{
				"msg": "密码不得小于 6 位",
			})
			return
		}
		if len(phone) != 11 {
			c.JSON(422, gin.H{
				"msg": "手机号必须为11位数",
			})
			return
		}
		var user model.UserBasic
		db.Where("username = ?", username).First(&user)
		if user.ID != 0 {
			c.JSON(http.StatusFailedDependency, gin.H{
				"msg": "已经存在的用户，无法注册",
			})
			return
		}
		db.Create(&model.UserBasic{
			Username: username,
			Password: password,
			Phone:    phone,
		})
		c.JSON(200, gin.H{
			"msg": "注册成功",
		})

	})
	return r
}
