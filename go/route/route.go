package route

import (
	"blog/controller"
	v1 "blog/route/v1"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine
func Initroute() *gin.Engine {
	r = v1.InitrouteV1()

	r.Static("/static","static")
	r.GET("/", func(c *gin.Context) {
		c.String(200,"hello ///")
	})

	//配置auth group
	auth_group := r.Group("auth")
	{
		auth_group.GET("check_pass",controller.CheckPass )
		auth_group.GET("like", func(c *gin.Context) {
			c.String(200,"like")
		})
	}
	article_group := r.Group("article")
	{
		article_group.GET("/",controller.GetArticles)
	}

	return r
}