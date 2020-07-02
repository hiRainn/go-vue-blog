package controller

import (
	"blog/pkg/errcode"
	"github.com/gin-gonic/gin"
)

//get background menu
func GetMenu(ctx *gin.Context) {
	//举例直接返回获取菜单错误
	ctx.JSON(200,errcode.GetMenuError.GetH())
}
