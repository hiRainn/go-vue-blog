package bac

import (
	"blog/model"
	"blog/pkg/errcode"
	"github.com/gin-gonic/gin"
	"net/http"
	//"blog/pkg/errcode"
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

}
