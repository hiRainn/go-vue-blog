package bac

import (
	"blog/model"
	"blog/pkg/errcode"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCategory(ctx *gin.Context) {

}

func AddCategory(ctx *gin.Context) {
	var params map[string] interface{}
	if err := ctx.BindJSON(&params); err != nil {
		ctx.JSON(http.StatusOK,errcode.ParamError.GetH())
		return ;
	}
	var cate model.BlogCate
	cate.CateName = params["name"].(string)
	//determine whether the name is empty
	if cate.CateName == "" {
		ctx.JSON(http.StatusOK,errcode.AddCateGoryNameError.GetH())
		return ;
	}
	//determine whether the name is repeat
	if cate.CheckNameRepeat() {
		ctx.JSON(http.StatusOK,errcode.AddCateGoryRepeat.GetH())
		return ;
	}
	//add cate
	res := cate.AddCate(cate)
	if res == 0 {
		if cate.CheckNameRepeat() {
			ctx.JSON(http.StatusOK,errcode.AddCateGoryRepeat.GetH())
			return ;
		}
	}
	ctx.JSON(http.StatusOK,errcode.Ok.GetH())
}
