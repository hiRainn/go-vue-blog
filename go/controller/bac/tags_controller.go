package bac

import (
	"blog/model"
	"blog/pkg/errcode"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTags(ctx *gin.Context) {

}

func TagsForSelect(ctx *gin.Context) {
	var cate *model.BlogTags
	list , err := cate.GetSelectList()

	if err != nil {
		ctx.JSON(http.StatusOK,err.GetH())
		return ;
	}
	var data = errcode.Ok

	ctx.JSON(http.StatusOK,data.SetData(map[string]interface{}{"list":list}))
}

