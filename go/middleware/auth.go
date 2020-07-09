package middleware

import (
	"blog/pkg/errcode"
	"github.com/gin-gonic/gin"
	"net/http"

)

func CheckLogin() gin.HandlerFunc{
	return func(c *gin.Context){
		v,find := Cache.Get("token")
		token := c.GetHeader("X-Token")
		if !find || v != token {
			c.Abort()
			c.JSON(http.StatusOK,errcode.AuthError.GetH())
			return;
		}
		c.Next()
	}
}
