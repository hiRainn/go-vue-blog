package middleware

import (
	"blog/config"
	"blog/pkg/errcode"
	"github.com/gin-gonic/gin"
	"net/http"

)

func CheckLogin() gin.HandlerFunc{
	return func(c *gin.Context){
		conf := config.GetConf()
		//if debug is false
		if !conf.Debug {
			v,find := Cache.Get("token")
			token := c.GetHeader("X-Token")
			if !find || v != token {
				c.Abort()
				c.JSON(http.StatusOK,errcode.AuthError.GetH())
				return;
			}
			c.Next()
		} else {
			c.Next()
		}

	}
}
