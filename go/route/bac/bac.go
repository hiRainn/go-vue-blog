package bac

import (
	"blog/controller/bac"
	"github.com/gin-gonic/gin"
)

//background route
func InitrouteBac() *gin.Engine {
	r := gin.Default()
	//all route in bac group
	bac_group := r.Group("bac")
	{
		//------------- init ------------------------
		//check init
		bac_group.GET("check_init",bac.CheckInit)
		//init
		bac_group.POST("init",bac.BacInit)
		//------------auth && password ----------------
		//log in
		bac_group.POST("auth",bac.Auth)
		//change pass
		bac_group.POST("change_pass",bac.ChangePass)

		//------------article -------------
		//get bac article list
		bac_group.GET("article",bac.GetArticles)
		//add article
		bac_group.POST("article",bac.AddArticle)
		//update article
		bac_group.PUT("article/:id",bac.AddArticle)
		//delete article
		bac_group.DELETE("article/:id",bac.DelArticle)

		//-----------comment ---------------
		//get bac comment list
		bac_group.GET("comment",bac.GetComments)
		//add comment
		bac_group.POST("comment",bac.AddComment)
		//update comment
		bac_group.PUT("comment/:id",bac.AddComment)
		//delete comment
		bac_group.DELETE("comment/:id",bac.DelComment)


	}
	return r
}
