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


func GetStat(ctx *gin.Context) {
	res := make(map[string]interface{},0)
	//get article_nums
	var a model.BlogArticle
	res["article_num"],_ = a.GetAppArticleNum()

}

func GetCateMenu(ctx *gin.Context) {
	var c model.BlogCate
	if res,err := c.MenuCate();err != nil {
		ctx.JSON(http.StatusOK,errcode.SystemError.GetH())
		return
	} else {
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
	}
}