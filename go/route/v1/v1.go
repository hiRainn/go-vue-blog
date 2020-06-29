package v1

import (
	"github.com/gin-gonic/gin"
)

func InitrouteV1() *gin.Engine {
	r := gin.Default()
	r.GET("/heihei", func(c *gin.Context) {
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