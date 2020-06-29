package controller

import (
	"blog/model"
	"github.com/gin-gonic/gin"
	"net/http"
)


func CheckPass(context *gin.Context) {
	username := context.DefaultQuery("username","heihei")
	password := context.DefaultQuery("password","123456")
	user := model.BlogAuth{Username: username,Password: password}
	auth := model.CheckAuth(user)
	var str string
	if auth == true {
		str = "通过结果"
	} else {
		str = "滚蛋"
	}
	context.String(http.StatusOK,str)
}
