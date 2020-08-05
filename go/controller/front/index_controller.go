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

func GetCateMenu(ctx *gin.Context) {
	var c model.BlogCate
	if res,err := c.MenuCate();err != nil {
		ctx.JSON(http.StatusOK,errcode.SystemError.GetH())
		return
	} else {
		ctx.JSON(http.StatusOK,errcode.Ok.SetData(res))
	}
}

func ViewArticle(module, ip string, article_id int) {
	var view model.BlogView
	view.Ip = ip
	view.Module = module
	view.FindRecord()
	if view.Id == 0 || view.CreatedAt  + 10 < time.Now().Unix() {
		if article_id != 0 {
			view.ArticleId = article_id
		}
		view.Id=0
		view.CreatedAt = time.Now().Unix()
		view.AddRecord()
	}
}