package controller

import "github.com/gin-gonic/gin"

func GetContent(str string) {
	var c *gin.Context
	c.String(200,str)
}