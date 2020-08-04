package bac

import (
	"blog/model"
	"blog/pkg/errcode"
	"blog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

//if there is no data in config table,then init it
func CheckInit(ctx *gin.Context) {
	var conf model.BlogConfig
	not_data, err := conf.HasData()
	var result gin.H

	// there's data in table
	if err == nil {
		result = errcode.InitError.GetH()
	} else {
		// not data in table
		if not_data {
			result = errcode.Ok.GetH()
		} else {
			// data base error
			result = errcode.DataBaseError.GetH()
		}
	}

	ctx.JSON(http.StatusOK,result)
}

//init
func BacInit(ctx *gin.Context) {

	var auth model.BlogAuth

	var params map[string] interface{}
	if err := ctx.BindJSON(&params); err != nil {
		ctx.JSON(http.StatusOK,errcode.ParamError.GetH())
		return
	}
	//get username,nickname,intro
	auth.Username, _  = params["username"].(string)
	auth.Nickname, _  = params["nickname"].(string)
	auth.Intro,_ = params["intro"].(string)
	//get birthday
	date ,_:= params["birthday"].(string)
	if date == "" {
		auth.Birthday = 0
	} else {
		date = strings.Trim(date,"\r\n")
		timeLayout := "2006-01-02"
		times ,_ := time.ParseInLocation(timeLayout,date,time.Local)
		auth.Birthday = times.Unix()
	}
	//validate password
	auth.Password, _ = params["password"].(string)
	repeat,_ := params["repeat"].(string)

	if auth.Password == "" {
		ctx.JSON(http.StatusOK,errcode.EmptyPassError.GetH())
		return
	}
	if auth.Password != repeat {
		ctx.JSON(http.StatusOK,errcode.RepeatPassError.GetH())
		return
	}
	auth.Password = utils.PassEncry(auth.Password)

	title, _ := params["title"].(string)
	subtitle , _ := params["subtitle"].(string)
	lang := params["language"].(string)

	conf := []model.BlogConfig{
	 	{Key: "title",Value: title},
	 	{Key: "subtitle", Value: subtitle},
		{Key: "language", Value: lang},
	}
	if title == "" {
		ctx.JSON(http.StatusOK,errcode.BlogNameEmpty.GetH())
		return
	}
	//start translate
	tx := model.DB().Begin()
	if err1 := tx.Create(&auth).Error; err1 != nil {
		tx.Rollback()
		//ctx.JSON(http.StatusOK,gin.H{"data":date})
		ctx.JSON(http.StatusOK,errcode.InsertAuthError.GetH())
		return
	}
	for _ , v := range conf {
		c := &model.BlogConfig{Key: v.Key,Value: v.Value}
		if err2 := tx.Create(c).Error;err2 != nil {
			tx.Rollback()
			ctx.JSON(http.StatusOK,errcode.InsertConfError.GetH())
			return
		}
	}
	tx.Commit()

	ctx.JSON(http.StatusOK,errcode.Ok.GetH())
}
