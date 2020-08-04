package bac

import (
	"blog/model"
	"blog/pkg/errcode"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCategory(ctx *gin.Context) {

}

func CateForSelect(ctx *gin.Context) {
	var cate *model.BlogCate
	list , err := cate.GetSelectList()

	if err != nil {
		ctx.JSON(http.StatusOK,err.GetH())
		return ;
	}
	var data = errcode.Ok

	ctx.JSON(http.StatusOK,data.SetData(map[string]interface{}{"list":list}))
}

func AddCategory(ctx *gin.Context) {
	var params map[string] interface{}
	if err := ctx.BindJSON(&params); err != nil {
		ctx.JSON(http.StatusOK,errcode.ParamError.GetH())
		return ;
	}
	cate := new(model.BlogCate)
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
	res := cate.AddCate()
	if res == 0 {
		ctx.JSON(http.StatusOK,errcode.AddCateGoryError.GetH())
		return ;
	}

	ctx.JSON(http.StatusOK,errcode.Ok.SetData(map[string]interface{}{"id":res}))
}
