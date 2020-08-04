package route

import (
	"blog/route/bac"
	"github.com/gin-gonic/gin"
	b "blog/controller/bac"
	f "blog/controller/front"

)

var r *gin.Engine
func Initroute() *gin.Engine {
	r = bac.InitrouteBac()

	r.Static("/static","static")
	//------------- init ------------------------
	//check init
	r.GET("/bac/check_init",b.CheckInit)
	//init
	r.POST("/bac/init",b.BacInit)
	//------------auth && password ----------------
	//log in
	r.POST("/bac/auth",b.Auth)

	//---------------route for app ----------------
	//get base info
	r.GET("info",f.Index)
	//get cate_article
	r.GET("cate_article",f.GetCateMenu)
	r.GET("article", f.GetArticle)


	return r
}