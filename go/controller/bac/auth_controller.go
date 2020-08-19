package bac

import (
	"blog/model"
	"blog/pkg/errcode"
	"blog/utils"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"math/rand"
	"net/http"
	"time"
	"blog/middleware"
)

//log in
func Auth(ctx *gin.Context) {
	var params map[string] interface{}
	if err := ctx.BindJSON(&params); err != nil {
		ctx.JSON(http.StatusOK,errcode.ParamError.GetH())
		return ;
	}
	username , _:= params["username"].(string)
	password ,_ := params["password"].(string)
	if username == "" || password == "" {
		ctx.JSON(http.StatusOK,errcode.ParamError.GetH())
		return ;
	}

	encry := utils.PassEncry(password)
	auth := model.BlogAuth{Username: username,Password: encry}

	//if auth fail ,return error
	if err := auth.CheckAuth();err != nil {
		ctx.JSON(http.StatusOK,err.GetH())
		return ;
	}

	//cache token for 3 days

	token := getToken(32)
	middleware.Cache.Set("token",token,cache.DefaultExpiration)
	data := map[string] interface{}{}
	data["token"] = token

	ctx.JSON(http.StatusOK,errcode.Ok.SetData(data))
}

//set token
func getToken(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRXTUVWXZY"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}


func ChangePass(ctx *gin.Context)  {
	var params map[string] interface{}
	if err := ctx.BindJSON(&params); err != nil {
		ctx.JSON(http.StatusOK,errcode.ParamError.GetH())
		return
	}
	newpass1 , _:= params["newpass1"].(string)
	newpass2 , _:= params["newpass2"].(string)
	password ,_ := params["password"].(string)
	if newpass1 == "" || password == "" || newpass1 != newpass2 {
		ctx.JSON(http.StatusOK,errcode.ParamError.GetH())
		return
	}

	encry := utils.PassEncry(password)
	auth := model.BlogAuth{Password: encry}

	//if auth fail ,return error
	if err := auth.CheckAuth();err != nil {
		ctx.JSON(http.StatusOK,err.GetH())
		return
	}

	newpass := utils.PassEncry(newpass1)
	auth.ChangePass(newpass)
	ctx.JSON(http.StatusOK,errcode.Ok.GetH())
}