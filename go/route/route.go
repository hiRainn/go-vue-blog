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
	r.GET("api/bac/check_init",b.CheckInit)
	//init
	r.POST("api/bac/init",b.BacInit)
	//------------auth && password ----------------
	//log in
	r.POST("api/bac/auth",b.Auth)

	//---------------route for app ----------------
	//get base info
	r.GET("api/info",f.Index)
	//get cate_article
	r.GET("api/cate_article",f.GetCateMenu)
	r.GET("api/article", f.GetArticleList)
	r.GET("api/article/:id",f.GetArticle)
	r.GET("api/comment/:id",f.GetArticleComment)
	r.POST("api/comment",f.PostComment)
	r.GET("api/message",f.GetMessage)
	r.POST("api/message",f.PostComment)

	//right_info
	r.GET("api/friends",f.GetFriendsLink)
	r.GET("api/latest_comments",f.GetLatestComments)
	r.GET("api/tags",f.GetTagsNum)
	r.GET("api/most",f.GetClickMostArticle)
	r.GET("api/stat",f.GetStat)
	r.GET("api/filing",f.GetFiling)

	//like/dislike/report
	r.POST("api/like",f.Like)
	r.POST("api/dislike",f.DisLike)
	r.POST("api/report",f.Report)
	r.DELETE("api/like",f.CancleLike)
	r.DELETE("api/dislike",f.CancleDislile)
	r.DELETE("api/report",f.CancleReport)
	r.POST("api/like_article",f.Like)
	r.GET("api/about",f.AboutMe)
	r.GET("api/index_article",f.GetIndexArticle)



	return r
}