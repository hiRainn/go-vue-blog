package bac

import (
	"blog/model"
	"blog/pkg/errcode"
	"blog/utils"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"net/http"
	"time"
)

//log in
func Auth(ctx *gin.Context) {
	params := map[string] interface{}{}
	if err := ctx.BindJSON(&params); err != nil {
		ctx.JSON(http.StatusOK,errcode.ParamError.GetH())
		return ;
	}
	username := params["username"].(string)
	password := params["password"].(string)
	if username == "" || password == "" {
		ctx.JSON(http.StatusOK,errcode.ParamError.GetH())
		return ;
	}
	auth := &model.BlogAuth{}

	encry := utils.PassEncry(password)
	if encry != password {
		ctx.JSON(http.StatusOK,errcode.PasswordError.GetH())
		return ;
	}

	c := cache.New(72*time.Hour, 10*time.Minute)
	token := setToken()
	c.Set("token",token,cache.DefaultExpiration)

	var data map[string] interface{}
	data["token"] = token

	ctx.JSON(http.StatusOK,errcode.Ok.SetData(data))
}

func setToken() string {
	return "dsadsacdakdsakdsadsa"
}


func ChangePass(ctx *gin.Context)  {
	
}