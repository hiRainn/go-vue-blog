package front

import (
	"blog/model"
	"blog/pkg/errcode"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)


//is_self eq 0 and sort for is_top desc sort desc
func GetArticleList(ctx *gin.Context) {
	page_num := ctx.DefaultQuery("page_size","0") //number of crticle per page
	p := ctx.DefaultQuery("p","1")
	cate_id := ctx.DefaultQuery("cate_id","0")
	tag:= ctx.DefaultQuery("tag","")

	//record view
	token := ctx.GetHeader("X-Token")
	go ViewRecord("article_list",ctx.ClientIP(),0,token)

	var page_size int
	if page_num == "" || page_num == "0" {
		page_size = 20
	} else {
		page_size,_ = strconv.Atoi(page_num)
	}
	//select condition where
	condition := make(map[string]interface{})
	if cate_id != "0" && cate_id != "" {
		condition["cate_id"],_ = strconv.Atoi(cate_id)
	}
	if len(tag) != 0 {
		var tags model.BlogTags
		tags.TagName = tag
		if err :=tags.GetInfo(); err != nil {
			ctx.JSON(http.StatusOK,err.GetH())
			return
		}
		condition["tag_id"] = tags.Id
	}
	//page information
	page := make(map[string]int)
	page["page_size"] = page_size
	page["page"],_ = strconv.Atoi(p)
	var art model.BlogArticle
	res,count,_ := art.GetAppArticleList(condition,page)

	//simple to process content format
	for i,_ := range res {
		res[i].Content = strings.Replace(res[i].Content,"#","",-1)
		res[i].Content = strings.Replace(res[i].Content,"-",",",-1)
	}

	ctx.JSON(http.StatusOK,errcode.Ok.SetData(map[string]interface{}{"list":res,"p":page["page"],"page_size":page_size,"total":count}))
}

//get article info by id  ---- content page
func GetArticle (ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusOK,errcode.ParamError.GetH())
		return
	}
	article_id ,_ := strconv.Atoi(id)
	token := ctx.GetHeader("X-Token")
	go ViewRecord("article",ctx.ClientIP(),article_id,token)
	var art model.BlogArticle
	art.Id = article_id
	if res,err := art.GetAppArticle(); err != nil {
		ctx.JSON(http.StatusOK,err.GetH())
		return
	} else {
		ctx.JSON(http.StatusOK,errcode.Ok.SetData(res))
	}
}

