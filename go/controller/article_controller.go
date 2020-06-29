package controller

import (
	"github.com/gin-gonic/gin"
)

func AddArticle(context *gin.Context) {

}

func GetArticles(context *gin.Context) {
	page := context.DefaultQuery("page","1")
	query := context.DefaultQuery("s","")
	context.String(200,page + query)
	

}
