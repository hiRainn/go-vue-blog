package front

import (
	"blog/model"
	"blog/pkg/errcode"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//is_self eq 0 and sort for is_top desc sort desc
func GetArticle(ctx *gin.Context) {
	page_num := ctx.DefaultQuery("page_size","0") //number of crticle per page
	p := ctx.DefaultQuery("p","1")
	cate_id := ctx.DefaultQuery("cate_id","0")
	tag_id,_ := ctx.GetQueryArray("tag_id[]")
	title := ctx.DefaultQuery("title","")
	var page_size int
	if page_num == "" || page_num == "0" {
		page_size = 20
	} else {
		page_size,_ = strconv.Atoi(page_num)
	}
	//select condition where
	condition := make(map[string]interface{})
	condition["status"],_ = strconv.Atoi(ctx.DefaultQuery("status","0"))
	if cate_id != "0" && cate_id != "" {
		condition["cate_id"],_ = strconv.Atoi(cate_id)
	}
	if title != "" {
		condition["title"] = title
	}
	if len(tag_id) != 0 {
		condition["tag_id"] = tag_id
	}
	//page information
	page := make(map[string]int)
	page["page_size"] = page_size
	page["page"],_ = strconv.Atoi(p)
	var art model.BlogArticle
	res,count,_ := art.GetArticleList(condition,page)

	ctx.JSON(http.StatusOK,errcode.Ok.SetData(map[string]interface{}{"list":res,"p":page["page"],"page_size":page_size,"total":count}))
}
