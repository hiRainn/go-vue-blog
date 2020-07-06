package route

import (
	"blog/route/bac"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine
func Initroute() *gin.Engine {
	r = bac.InitrouteBac()

	r.Static("/static","static")
	r.GET("/article", func(c *gin.Context) {
		c.String(200,"hello ///")
	})

	//配置auth group
	auth_group := r.Group("auth")
	{
		auth_group.GET("like", func(c *gin.Context) {
			c.String(200,"like")
		})
	}

	return r
}