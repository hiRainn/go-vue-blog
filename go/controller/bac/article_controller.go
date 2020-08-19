package bac

import (
	"blog/model"
	"blog/pkg/errcode"
	"blog/pkg/status"
	"blog/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//post
func PostArticle(ctx *gin.Context) {
	//var article model.BlogArticle
	var params map[string] interface{}
	if err := ctx.BindJSON(&params); err != nil {
		ctx.JSON(http.StatusOK,errcode.ParamError.GetH())
		return
	}
	//determine params
	art := new(model.BlogArticle)
	if err :=determineArt(art,params); err != nil {
		ctx.JSON(http.StatusOK,err.GetH())
		return
	}
	cate := new(model.BlogCate)
	cate.Id = int(params["cate_id"].(float64))
	// start transaction
	tx := model.DB().Begin()
	//add category's num + 1
	if cate.SetIncNum(tx) == false {
		tx.Rollback()
		ctx.JSON(http.StatusOK,errcode.SetIncCateError.GetH())
		return
	}
	//do with tags
	tags,err := getTagsIds(tx,params)
	if err != nil {
		ctx.JSON(http.StatusOK,err.GetH())
		return
	}
	//set tags num + 1
	tag  := new(model.BlogTags)
	if len(tags) != 0 {
		if tag.SetIncNum(tx,tags) == false {
			tx.Rollback()
			ctx.JSON(http.StatusOK,errcode.SetIncTagsError.GetH())
			return
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
	//add article
	article_id, err := art.AddArticle(tx)
	if err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusOK,err.GetH())
		return
	} else {
		tx.Commit()
		ctx.JSON(http.StatusOK,errcode.Ok.SetData(map[string]interface{}{"id":article_id}))
	}

}

func GetArticleInfo(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusOK,errcode.ParamError.GetH())
		return
	}
	var art model.BlogArticle
	art.Id,_ = strconv.Atoi(id)
	if art.Id == 0 {
		ctx.JSON(http.StatusOK,errcode.ParamError.GetH())
		return
	}
	err :=art.GetArticleById()
	if err != nil {
		ctx.JSON(http.StatusOK,err.GetH())
		return
	}
	ctx.JSON(http.StatusOK,errcode.Ok.SetData(art))
}

//edit
func EditArticle(ctx *gin.Context) {
	//var article model.BlogArticle
	var params map[string] interface{}
	if err := ctx.BindJSON(&params); err != nil {
		ctx.JSON(http.StatusOK,errcode.ParamError.GetH())
		return
	}
	art := new(model.BlogArticle)
	art.Id = int(params["id"].(float64))
	if art.Id == 0 {
		ctx.JSON(http.StatusOK,errcode.ParamError.GetH())
		return
	}
	if err :=art.GetArticleById(); err != nil {
		ctx.JSON(http.StatusOK,err.GetH())
		return
	}
	cate := new(model.BlogCate)
	tag := new(model.BlogTags)
	//start transaction
	tx := model.DB().Begin()
	cate.Id = art.CateId
	//reduce cate num
	if cate.SetDecNum(tx) == false {
		tx.Rollback()
		ctx.JSON(http.StatusOK,errcode.SetDecCateError.GetH())
		return
	}
	//reduce tags num
	tag_array := strings.Split(art.TagsIds,",")
	if tag.SetDecNum(tx,tag_array) == false {
		tx.Rollback()
		ctx.JSON(http.StatusOK,errcode.SetDecTagsError.GetH())
		return
	}
	//set art.TagsIds to empty
	art.TagsIds = ""
	// get params
	if err :=determineArt(art,params); err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusOK,err.GetH())
		return
	}
	//set new cate num + 1
	cate.Id = art.CateId
	if cate.SetIncNum(tx) == false {
		tx.Rollback()
		ctx.JSON(http.StatusOK,errcode.SetIncCateError.GetH())
		return
	}
	//get new tags
	tags,err := getTagsIds(tx,params)
	if err != nil {
		ctx.JSON(http.StatusOK,err.GetH())
		return
	}
	//set tags num + 1
	if len(tags) != 0 {
		if tag.SetIncNum(tx,tags) == false {
			tx.Rollback()
			ctx.JSON(http.StatusOK,errcode.SetIncTagsError.GetH())
			return
		}
		//set struct article's tags, id,id,id ...
		for _, v :=  range tags {
			art.TagsIds = art.TagsIds + strconv.Itoa(v) + ","
		}
	}

	art.TagsIds = strings.Trim(art.TagsIds,",")
	art.ModifyAt = time.Now().Unix()
	//edit article
	if tx.Where("id = ?",art.Id).Save(art).Error != nil {
		tx.Rollback()
		ctx.JSON(http.StatusOK,errcode.EditArticleError.GetH())
		return
	}
	tx.Commit()
	ctx.JSON(http.StatusOK,errcode.Ok.GetH())


}

func PostSaveArticle(ctx *gin.Context) {
	//var article model.BlogArticle
	var params map[string] interface{}
	if err := ctx.BindJSON(&params); err != nil {
		ctx.JSON(http.StatusOK,errcode.ParamError.GetH())
		return
	}
	art := new(model.BlogArticle)
	art.Id = int(params["id"].(float64))
	if art.Id == 0 {
		ctx.JSON(http.StatusOK,errcode.ParamError.GetH())
		return
	}
	if err :=art.GetArticleById(); err != nil {
		ctx.JSON(http.StatusOK,err.GetH())
		return
	}
	cate := new(model.BlogCate)
	tag := new(model.BlogTags)
	//start transaction
	tx := model.DB().Begin()

	// get params
	if err :=determineArt(art,params); err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusOK,err.GetH())
		return
	}
	art.Status = status.OK.GetCode()
	//set cate num + 1
	cate.Id = art.CateId
	if cate.SetIncNum(tx) == false {
		tx.Rollback()
		ctx.JSON(http.StatusOK,errcode.SetIncCateError.GetH())
		return
	}
	//get new tags
	tags,err := getTagsIds(tx,params)
	if err != nil {
		ctx.JSON(http.StatusOK,err.GetH())
		return
	}
	//set TagsIds to empty
	art.TagsIds = ""
	//set tags num + 1
	if len(tags) != 0 {
		if tag.SetIncNum(tx,tags) == false {
			tx.Rollback()
			ctx.JSON(http.StatusOK,errcode.SetIncTagsError.GetH())
			return
		}
		//set struct article's tags, id,id,id ...
		for _, v :=  range tags {
			art.TagsIds = art.TagsIds + strconv.Itoa(v) + ","
		}
	}


	art.TagsIds = strings.Trim(art.TagsIds,",")
	//edit article
	if tx.Where("id = ?",art.Id).Save(art).Error != nil {
		tx.Rollback()
		ctx.JSON(http.StatusOK,errcode.EditArticleError.GetH())
		return
	}
	tx.Commit()
	ctx.JSON(http.StatusOK,errcode.Ok.GetH())
}


//save article to drafts
func SaveArticle(ctx *gin.Context) {
	//var article model.BlogArticle
	var params map[string] interface{}
	if err := ctx.BindJSON(&params); err != nil {
		ctx.JSON(http.StatusOK,errcode.ParamError.GetH())
		return
	}
	//determine params
	art := new(model.BlogArticle)
	if err :=determineArt(art,params); err != nil {
		ctx.JSON(http.StatusOK,err.GetH())
		return
	}
	art.Status = status.ArticleSave.GetCode()

	// start transaction
	tx := model.DB().Begin()
	//do with tags
	tags,err := getTagsIds(tx,params)
	if err != nil {
		ctx.JSON(http.StatusOK,err.GetH())
		return
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

	//determine if is in article table
	art.Id = int(params["id"].(float64))
	if art.Id != 0 {
		//edit article
		_,err := art.EditArticle(tx)
		if err != nil {
			tx.Rollback()
			ctx.JSON(http.StatusOK,err.GetH())
			return
		} else {
			tx.Commit()
			ctx.JSON(http.StatusOK,errcode.Ok.SetData(map[string]interface{}{"id":art.Id}))
		}
	} else {
		//add article
		article_id, err := art.AddArticle(tx)
		if err != nil {
			tx.Rollback()
			ctx.JSON(http.StatusOK,err.GetH())
			return
		} else {
			tx.Commit()
			ctx.JSON(http.StatusOK,errcode.Ok.SetData(map[string]interface{}{"id":article_id}))
		}
	}
}

//determine params
func determineArt(art *model.BlogArticle,params map[string] interface{}) *errcode.ERRCODE {
	art.Title = params["title"].(string)
	if art.Title == "" {
		return errcode.ArticleTitleEmpty
	}
	art.CateId = int(params["cate_id"].(float64))
	if art.CateId == 0 {
		return errcode.ArticleCateEmpty
	}
	art.Content = params["content"].(string)
	if art.Content == "" {
		return errcode.ArticleContentEmpty
	}
	top:= params["is_top"].(bool)
	if top {
		art.IsTop = 1
		switch params["sort"].(type) {
		case nil:
			art.Sort = 0
		case float64:
			art.Sort = uint8(params["sort"].(float64))
		case string:
			tmp,_ := strconv.Atoi(params["sort"].(string))
			art.Sort = uint8(tmp)
		}
	} else {
		art.IsTop = 0
		art.Sort = 0
	}
	//is original
	original,_ := params["is_original"].(bool)
	if original {
		art.IsOriginal = 1
		art.ReprintFrom = ""
	} else {
		art.IsOriginal = 0
		art.ReprintFrom = params["reprint_from"].(string)
	}
	//only see by yourself
	self := params["is_self"].(bool)
	if self {
		art.IsSelf = 1
	} else {
		art.IsSelf = 0
	}
	//if comment is forbidden
	allow_comment := params["allow_comment"].(bool)
	if allow_comment {
		art.AllowComment = 1
	} else {
		art.AllowComment = 0
	}

	//do with created_at
	the_time := params["create_at"].(string)
	if the_time == "" {
		art.CreatedAt = time.Now().Unix()
	} else {
		art.CreatedAt = utils.YmdhmTotimestamp(the_time)
	}
	return nil
}

func getTagsIds(tx *gorm.DB,params map[string]interface{}) ([]int,*errcode.ERRCODE) {
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
				return tags,errcode.AddTagNameError
			}
			tag.TagName = value2
			repeat, id := tag.CheckNameRepeat()
			if repeat == true {
				tags = append(tags,id)
			} else {
				// create tags and return id
				tag.Id = 0
				tag_id := tag.AddCate(tx)
				// add fail
				if tag_id == 0 {
					return tags,errcode.AddTagError
				}
				tags = append(tags,tag_id)
			}
		}
	}
	return tags,nil
}

//articles list
func GetArticles(ctx *gin.Context) {
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


func DelArticle(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusOK,errcode.ParamError)
		return
	}
	var art model.BlogArticle
	art.Id,_ = strconv.Atoi(id)
	err :=art.GetArticleById()
	if err != nil {
		ctx.JSON(http.StatusOK,err.GetH())
		return
	}
	//if article is draft just delete
	if art.Status == status.ArticleSave.GetCode() {
		if err := art.DelArticleById(); err != nil {
			ctx.JSON(http.StatusOK,err.GetH())
			return
		} else {
			ctx.JSON(http.StatusOK,errcode.Ok.GetH())
		}
	} else {
		//must delete tag number and category number
		cate := new(model.BlogCate)
		tag := new(model.BlogTags)
		//start transaction
		tx := model.DB().Begin()
		cate.Id = art.CateId
		//reduce cate num
		if cate.SetDecNum(tx) == false {
			tx.Rollback()
			ctx.JSON(http.StatusOK,errcode.SetDecCateError.GetH())
			return
		}
		//reduce tags num
		tag_array := strings.Split(art.TagsIds,",")
		if tag.SetDecNum(tx,tag_array) == false {
			tx.Rollback()
			ctx.JSON(http.StatusOK,errcode.SetDecTagsError.GetH())
			return
		}

		if tx.Delete(art).Error != nil {
			tx.Rollback()
			ctx.JSON(http.StatusOK,errcode.DeleteArticleError.GetH())
			return
		}
		tx.Commit()
		ctx.JSON(http.StatusOK,errcode.Ok.GetH())
	}

}
