package bac

import (
	"blog/model"
	"blog/pkg/errcode"
	"blog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
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
	cate := new(model.BlogCate)
	art.Title = params["title"].(string)
	if art.Title == "" {
		ctx.JSON(http.StatusOK,errcode.ArticleTitleEmpty.GetH())
		return ;
	}
	art.CateId = int(params["cate_id"].(float64))
	cate.Id = int(params["cate_id"].(float64))
	if art.CateId == 0 {
		ctx.JSON(http.StatusOK,errcode.ArticleCateEmpty.GetH())
		return ;
	}

	art.Content = params["content"].(string)
	if art.Content == "" {
		ctx.JSON(http.StatusOK,errcode.ArticleContentEmpty.GetH())
		return ;
	}
	top := params["is_top"].(bool)
	if top {
		art.IsTop = 1
	}
	art.Sort = uint8(params["sort"].(float64))

	//do with created_at
	the_time := params["create_at"].(string)
	if the_time == "" {
		art.CreatedAt = time.Now().Unix()
	} else {
		art.CreatedAt = utils.YmdTotimestamp(the_time)
	}

	// start transaction
	tx := model.DB().Begin()
	//do with tags
	var tags []int
	tag  := new(model.BlogTags)
	for _ , val := range params["tags"].([]interface{}) {
		//if get int , then it is the id of tags
		if value1, ok1 := val.(float64); ok1 == true {
			tags = append(tags,int(value1))
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
	if len(tags) != 0 {
		if tag.SetIncNum(tx,tags) == false {
			tx.Rollback()
			ctx.JSON(http.StatusOK,errcode.SetIncTagsError.GetH())
			return ;
		}
	}

	//set struct article's tags, id,id,id ...
	//determine if tags is []
	if len(tags) == 0 {
		art.TagsIds = ""
	} else {
		for _, v :=  range tags {
			art.TagsIds = art.TagsIds + strconv.Itoa(v) + ","
		}
	}
	art.TagsIds = strings.Trim(art.TagsIds,",")
	//add category's num + 1
	if cate.SetIncNum(tx) == false {
		tx.Rollback()
		ctx.JSON(http.StatusOK,errcode.SetIncCateError.GetH())
		return ;
	}

	//add article
	article_id, err := art.AddArticle(tx)
	if err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusOK,err.GetH())
		return ;
	} else {
		tx.Commit()
		ctx.JSON(http.StatusOK,errcode.Ok.SetData(map[string]interface{}{"id":article_id}))
	}

}




//articles list
func GetArticles(ctx *gin.Context) {

	ctx.JSON(http.StatusOK,errcode.Ok.GetH())
}


func DelArticle(ctx *gin.Context) {

}
