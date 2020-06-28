package route

import (
	"github.com/gin-gonic/gin"
)

func Initroute() *gin.Engine {
	r := gin.Default()
	r.Static("/static","static")
	r.GET("/", func(c *gin.Context) {
		c.String(200,"hello heihei")
	})
	v1group := r.Group("v1")
	{
		v1group.GET("todo", func(c *gin.Context) {
			c.String(200,"todo")
		})
		v1group.GET("like", func(c *gin.Context) {
			c.String(200,"like")
		})
	}
	return r
}