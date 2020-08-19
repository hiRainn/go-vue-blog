package bac

import (
	"blog/model"
	"blog/pkg/errcode"
	"blog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetBlogSetting(c *gin.Context) {
	var conf model.BlogConfig
	list,_ := conf.Find()
	c.JSON(http.StatusOK,errcode.Ok.SetData(list))
}

func SetBlogSetting(c *gin.Context) {
	var params map[string] interface{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusOK,errcode.ParamError.GetH())
		return
	}
	var conf model.BlogConfig
	for key,value := range params {
		conf.Key = key
		v,_ := value.(string)
		conf.Value = v
		conf.Update()
	}
	c.JSON(http.StatusOK,errcode.Ok.SetData(conf))
}

func GetUserSetting(c *gin.Context) {
	var a model.BlogAuth
	a.GetAllInfo()
	c.JSON(http.StatusOK,errcode.Ok.SetData(a))
}

func SetUserSetting(c *gin.Context) {
	var params map[string] interface{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusOK,errcode.ParamError.GetH())
		return
	}
	the_time := params["birthday"].(string)
	if the_time != "" {
		params["birthday"] = utils.YmdTotimestamp(the_time)
	}
	var a model.BlogAuth
	a.SetInfo(params)
	c.JSON(http.StatusOK,errcode.Ok.SetData(a))
}
