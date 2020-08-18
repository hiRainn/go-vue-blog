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
	r.GET("article", f.GetArticleList)
	r.GET("article/:id",f.GetArticle)
	r.GET("comment/:id",f.GetArticleComment)
	r.POST("comment",f.PostComment)
	r.GET("message",f.GetMessage)
	r.POST("message",f.PostComment)

	//right_info
	r.GET("friends",f.GetFriendsLink)
	r.GET("latest_comments",f.GetLatestComments)
	r.GET("tags",f.GetTagsNum)
	r.GET("most",f.GetClickMostArticle)
	r.GET("stat",f.GetStat)
	r.GET("filing",f.GetFiling)

	//like/dislike/report
	r.POST("like",f.Like)
	r.POST("dislike",f.DisLike)
	r.POST("report",f.Report)
	r.DELETE("like",f.CancleLike)
	r.DELETE("dislike",f.CancleDislile)
	r.DELETE("report",f.CancleReport)
	r.POST("like_article",f.Like)
	r.GET("about",f.AboutMe)



	return r
}