package bac

import (
	"blog/model"
	"github.com/gin-gonic/gin"
)

//if there is no data in config table,then init it
func CheckInit(ctx *gin.Context) {
	conf := &model.BlogConfig{}
	res , err  := conf.Find()
	if err != nil {
		ctx.JSON(200,gin.H{"code":200,"msg":err.Error()})
	} else {
		ctx.JSON(200,gin.H{"code":201,"data":res})
	}
}

func BacInit(ctx *gin.Context) {

}
