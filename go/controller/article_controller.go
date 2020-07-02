package controller

import (
	"blog/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

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

func GetArticles(context *gin.Context) {
	page := context.DefaultQuery("page","1")
	query := context.DefaultQuery("s","")
	fmt.Println(page,query)
	data,err := model.GetArticleListByPage(1)
	if err != nil {
		ErrReturn(200,gin.H{"code":200,"msg":"未知错误","data":data},context)
	} else {
		ErrReturn(201,gin.H{"code":201,"msg":"未知错误","data":data},context)
	}



}
