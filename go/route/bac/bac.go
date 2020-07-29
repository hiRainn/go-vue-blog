package bac

import (
	"blog/controller/bac"
	"blog/middleware"
	"github.com/gin-gonic/gin"
)

//background route
func InitrouteBac() *gin.Engine {
	r := gin.Default()
	//all route in bac group
	bac_group := r.Group("bac").Use(middleware.CheckLogin())
	{
		//change pass
		bac_group.POST("change_pass",bac.ChangePass)

		//------------article -------------
		//get bac article list
		bac_group.GET("article",bac.GetArticles)
		//get bac article info
		bac_group.GET("article/:id",bac.GetArticleInfo)
		//add article
		bac_group.POST("article",bac.PostArticle)
		//save article
		bac_group.POST("save_article",bac.SaveArticle)
		//post save article
		bac_group.PUT("save_article",bac.PostSaveArticle)
		//update article
		bac_group.PUT("article",bac.EditArticle)
		//delete article
		bac_group.DELETE("article/:id",bac.DelArticle)
		//image
		bac_group.POST("article/upload",bac.UploadArticleImg)

		//-----------category-----------------------
		//get category list,category admin page
		bac_group.GET("cate",bac.GetCategory)
		//get all categories for select
		bac_group.GET("cate_for_select",bac.CateForSelect)
		//add category
		bac_group.POST("cate",bac.AddCategory)

		//-------------tags -------------------------
		//get category list,category admin page
		bac_group.GET("tags",bac.GetTags)
		//get all categories for select
		bac_group.GET("tags_for_select",bac.TagsForSelect)



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
