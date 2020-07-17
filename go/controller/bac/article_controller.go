package bac

import (
	"blog/model"
	"blog/pkg/errcode"
	"github.com/gin-gonic/gin"
	"net/http"
)

//增加、修改
func AddArticle(ctx *gin.Context) {
	var article model.BlogArticle

	var params map[string] interface{}
	if err := ctx.BindJSON(&params); err != nil {
		ctx.JSON(http.StatusOK,errcode.ParamError.GetH())
		return ;
	}

}


//articles list
func GetArticles(ctx *gin.Context) {

	ctx.JSON(http.StatusOK,errcode.Ok.GetH())
}


func DelArticle(ctx *gin.Context) {

}
