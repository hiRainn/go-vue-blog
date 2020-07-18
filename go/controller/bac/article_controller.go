package bac

import (
	"blog/model"
	"blog/pkg/errcode"
	"github.com/gin-gonic/gin"
	"net/http"
)

//增加、修改
func AddArticle(ctx *gin.Context) {
	//var article model.BlogArticle

	var params map[string] interface{}
	if err := ctx.BindJSON(&params); err != nil {
		ctx.JSON(http.StatusOK,errcode.ParamError.GetH())
		return ;
	}
	//determine params
	art := new(model.BlogArticle)
	art.Title = params["title"].(string)
	if art.Title == "" {
		ctx.JSON(http.StatusOK,errcode.ArticleTitleEmpty.GetH())
		return ;
	}
	art.CateId = params["cate_id"].(int)
	if art.CateId == 0 {
		ctx.JSON(http.StatusOK,errcode.ArticleCateEmpty.GetH())
		return ;
	}
	art.Content = params["content"].(string)
	if art.Content == "" {
		ctx.JSON(http.StatusOK,errcode.ArticleContentEmpty.GetH())
		return ;
	}
	art.IsTop = params["is_top"].(uint8)
	if art.IsTop != 0 {
		art.Sort = params["sort"].(uint8)
	}
	// start transaction
	tx := model.DB().Begin()
	//do with tags
	var tags []int
	var tag *model.BlogTags
	for _ , val := range params["tags"].([]interface{}) {
		//if get int , then it is the id of tags
		if value1, ok1 := val.(int); ok1 == true {
			tags = append(tags,value1)
		}
		//if get string ,then need to created
		if value2, ok2 := val.(string); ok2 == true {
			//determine whether is exitst,
			if value2 == "" {
				tx.Rollback()
				ctx.JSON(http.StatusOK,errcode.AddTagNameError.GetH())
				return ;
			}
			tag.TagName = value2
			repeat, id := tag.CheckNameRepeat()
			if repeat == true {
				tags = append(tags,id)
			} else {
				// create tags and return id
				tag_id := tag.AddCate(tx)
				// add fail
				if tag_id == 0 {
					tx.Rollback()
					ctx.JSON(http.StatusOK,errcode.AddTagError.GetH())
					return ;
				}
				tags = append(tags,tag_id)
			}
		}
	}
	//set tags num + 1
	if tag.SetIncNum(tx,tags) == false {
		tx.Rollback()
		ctx.JSON(http.StatusOK,errcode.SetIncTagsError.GetH())
		return ;
	}
	// set atricle tag_ids



}


//articles list
func GetArticles(ctx *gin.Context) {

	ctx.JSON(http.StatusOK,errcode.Ok.GetH())
}


func DelArticle(ctx *gin.Context) {

}
