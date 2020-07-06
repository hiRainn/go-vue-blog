package bac

import (
	"blog/model"
	"github.com/gin-gonic/gin"
	"time"
)

//增加、修改
func AddArticle(context *gin.Context) {
	model1 :=model.Model{
		Id: 2,
		CreatedAt: time.Now().Unix(),
		Status: 0,
	}
	article := model.BlogArticle{
		Title: "文章1",
		Abstract: "摘要1",
		Content: "内容1",
		ModifyAt: time.Now().Unix(),
		Model:model1,
	}

	article.CreatedAt = time.Now().Unix()
	article.Status = 0
	model.AddArticle(&article)
	article.Id = 3
	model.AddArticle(&article)
	article.Id = 4
	model.AddArticle(&article)

}

func DelArticle(ctx *gin.Context) {

}

func GetArticles(ctx *gin.Context) {
	ctx.JSON(200,gin.H{"msg":"dsads"})
}
