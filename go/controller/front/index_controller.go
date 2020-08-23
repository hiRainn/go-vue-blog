package front

import (
	"blog/model"
	"blog/pkg/errcode"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//get information index page
func Index(ctx *gin.Context) {

	token := ctx.GetHeader("X-Token")
	go ViewRecord("home",ctx.ClientIP(),0,token)

	//get blog's information
	var conf model.BlogConfig
	res := make(map[string]interface{})
	if confs , err := conf.Find(); err != nil {
		ctx.JSON(http.StatusOK,errcode.SystemError.GetH())
		return
	} else {
		tmp := make(map[string]interface{})
		for _,v := range confs {
			tmp[v.Key] = v.Value
		}
		res["blog_info"] = tmp
	}
	//get author's information
	var auth model.BlogAuth
	if auths,err := auth.GetInfo(); err != nil {
		ctx.JSON(http.StatusOK,errcode.SystemError.GetH())
		return
	} else {
		res["author"] = auths
	}
	ctx.JSON(http.StatusOK,errcode.Ok.SetData(res))
}

func GetIndexArticle(ctx *gin.Context) {
	var a model.BlogArticle
	list,_ := a.GetAppHomeArticle()
	ctx.JSON(http.StatusOK,errcode.Ok.SetData(list))
}


func GetStat(ctx *gin.Context) {
	res := make(map[string]interface{},0)
	//get article_nums
	var a model.BlogArticle
	res["article_num"],_ = a.GetAppArticleNum()
	//first article
	_ = a.GetFirstArticle()
	res["first_day"] = a.CreatedAt
	//get like number
	var l model.BlogLike
	res["like_number"] = l.GetIndexLike()

	//get comments number
	var c model.BlogComment
	res["comments_number"],_ = c.GetCommentNumber()
	res["msg_number"],_ = c.GetMsgNumber()

	//get stat
	var v model.BlogView
	timeStr := time.Now().Format("2006-01-02")
	t, _ := time.Parse("2006-01-02", timeStr)
	today0Clock := int(t.Unix())
	res["today"] ,_ = v.StatViews(today0Clock)
	res["week"],_ = v.StatViews(today0Clock - 7 * 86400)
	res["month"],_ = v.StatViews(today0Clock - 30 * 86400)
	res["year"],_ = v.StatViews(today0Clock - 365 * 86400)
	res["total"],_ = v.StatViews(0)

	ctx.JSON(http.StatusOK,errcode.Ok.SetData(res))

}

func GetCateMenu(ctx *gin.Context) {
	var c model.BlogCate
	if res,err := c.MenuCate();err != nil {
		ctx.JSON(http.StatusOK,errcode.SystemError.GetH())
		return
	} else {
		//get is_self = 1 article
		var a model.BlogArticle
		if list := a.GetSelfArticle();list != nil {
			for _,v := range list {
				for k,_ := range res {
					if res[k].Id == v.CateId {
						res[k].Num --
						break
					}
				}
			}
		}
		ctx.JSON(http.StatusOK,errcode.Ok.SetData(res))
	}
}

func GetLatestComments(ctx *gin.Context) {
	var c model.BlogComment
	//fmt.Println("debug")
	latest,_ := c.GetLastComment()
	//fmt.Println(latest)
	//fmt.Println("debug_over")
	ctx.JSON(http.StatusOK,errcode.Ok.SetData(latest))
}

func GetFriendsLink(ctx *gin.Context) {
	var f model.BlogFriends
	link,_ := f.GetLink()
	ctx.JSON(http.StatusOK,errcode.Ok.SetData(link))
}

func GetTagsNum(ctx *gin.Context) {
	var t model.BlogTags
	tags,_ := t.GetSelectList()
	ctx.JSON(http.StatusOK,errcode.Ok.SetData(tags))
}

func GetClickMostArticle(ctx *gin.Context) {
	var a model.BlogArticle
	most,_ := a.GetClickMost()
	ctx.JSON(http.StatusOK,errcode.Ok.SetData(most))
}

func GetFiling(ctx *gin.Context) {
	token := ctx.GetHeader("X-Token")
	go ViewRecord("filing",ctx.ClientIP(),0,token)
	var a model.BlogArticle
	list,_ := a.GetAllArticles()
	ctx.JSON(http.StatusOK,errcode.Ok.SetData(list))
}

func AboutMe(ctx *gin.Context) {
	var a model.BlogAuth
	about,_ := a.About()
	token := ctx.GetHeader("X-Token")
	go ViewRecord("about",ctx.ClientIP(),0,token)
	ctx.JSON(http.StatusOK,errcode.Ok.SetData(about.Intro))
}

func ViewRecord(module, ip string, article_id int,token string) {
	var view model.BlogView
	view.Ip = ip
	view.ArticleId = article_id
	view.Module = module
	view.FindRecord()

	if view.Id == 0 || view.CreatedAt  + 1800 < time.Now().Unix() {
		view.Token = token
		view.Id=0
		view.CreatedAt = time.Now().Unix()
		view.AddRecord()
		if view.ArticleId != 0 {
			var a model.BlogArticle
			a.Id = view.ArticleId
			a.AddView()
		}
	}
}