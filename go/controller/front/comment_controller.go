package front

import (
	"blog/model"
	"blog/pkg/errcode"
	"blog/pkg/status"
	"blog/service/SensitiveWords/tool"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func GetArticleComment(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusOK,errcode.ErrParams)
		return
	}
	var c model.BlogComment
	c.ArticleId,_ = strconv.Atoi(id)
	if list,total,err := c.GetCommentsByArticleId();err!=nil {
		ctx.JSON(http.StatusOK,err.GetH())
		return
	} else {
		ctx.JSON(http.StatusOK,errcode.Ok.SetData(map[string]interface{}{"list":list,"total":total}))
	}
}

func PostComment(ctx *gin.Context){
	var params map[string] interface{}
	if err := ctx.BindJSON(&params); err != nil {
		ctx.JSON(http.StatusOK,errcode.ParamError.GetH())
		return
	}
	article_id,_ := params["article_id"].(float64)
	if article_id == 0.00 {
		ctx.JSON(http.StatusOK,errcode.ParamError.SetData(article_id))
		return
	}
	name,_ := params["name"].(string)
	email,_ := params["email"].(string)
	pid,_ := params["pid"].(float64)
	content,_ := params["content"].(string)
	var c model.BlogComment

	c.ArticleId = int(article_id)
	c.Pid = int(pid)
	c.Name = name
	c.Email = email
	c.Content = content
	c.Status = status.CommentCheck.GetCode()
	c.CreatedAt = time.Now().Unix()


	_,ok := tool.GetMap().CheckSensitive(c.Content)
	if ok == false {
		c.Status = 0
	}
	if err := c.AddComment();err!= nil {
		ctx.JSON(http.StatusOK,err.GetH())
	} else {
		ctx.JSON(http.StatusOK,errcode.Ok.SetData(c.Status))
	}

}
