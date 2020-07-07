package bac

import (
	"blog/model"
	"blog/pkg/errcode"
	"github.com/gin-gonic/gin"
	"net/http"
	"blog/utils"
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
	//get username
	auth.Username = ctx.DefaultPostForm("username","")

	//get birthday
	date := ctx.DefaultPostForm("birthday","")
	if date == "" {
		auth.Birthday = 0
	} else {
		date = date + " 00:00:00"  //fomart
		timeLayout := "2020-07-07 12:00:00"
		loc, _ := time.LoadLocation("Local")
		times ,_ := time.ParseInLocation(timeLayout,date,loc)
		auth.Birthday = times.Unix()
	}
	//validate password
	password := ctx.DefaultPostForm("password","")
	repeat := ctx.DefaultPostForm("repeat","")
	if password == "" {

	}
}
