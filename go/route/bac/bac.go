package bac

import (
	"blog/controller"
	"github.com/gin-gonic/gin"
)

//background route
func InitrouteBac() *gin.Engine {
	r := gin.Default()
	//all route in bac group
	bac_group := r.Group("bac")
	{
		bac_group.GET("index", func(c *gin.Context) {
			c.String(200,"todo")
		})
		//get menu
		bac_group.POST("menu",controller.GetMenu)
		bac_group.GET("blog_list", func(c *gin.Context) {
			c.String(200,"like")
		})
	}
	return r
}
