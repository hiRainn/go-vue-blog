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
	"html"
)

func GetArticleComment(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusOK,errcode.ErrParams)
		return
	}
	var c model.BlogComment
	token := ctx.GetHeader("X-Token")
	c.ArticleId,_ = strconv.Atoi(id)
	if list,total,err := c.GetCommentsByArticleId(token);err!=nil {
		ctx.JSON(http.StatusOK,err.GetH())
		return
	} else {
		//process the data to [{id,children:[{id,children:[]}]

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
	submit,_ := params["submit"].(bool)
	token := ctx.GetHeader("X-Token")
	var c model.BlogComment


	c.ArticleId = int(article_id)
	c.Pid = int(pid)
	c.Name = html.EscapeString(name)
	c.Email = html.EscapeString(email)
	c.Content =  html.EscapeString(content)
	c.Status = status.CommentCheck.GetCode()
	c.CreatedAt = time.Now().Unix()
	c.Token = token

	//sensitive words check
	msg := tool.GetMap().FindAllSensitive(c.Content)
	if len(msg) == 0 {
		c.Status = 0
	}

	if len(msg) != 0 && submit == false {
		ctx.JSON(http.StatusOK,errcode.Ok.SetData(map[string]interface{}{"status":c.Status,"msg":msg,"submit":submit}))
		return
	}
	if err := c.AddComment();err!= nil {
		ctx.JSON(http.StatusOK,err.GetH())
	} else {
		ctx.JSON(http.StatusOK,errcode.Ok.SetData(map[string]interface{}{"status":c.Status,"id":c.Id,"submit":submit}))
	}
}

type tree struct {
	id int
	pid int
	data model.AppCommentList
	children *[]tree
}
func getTree(list *[]model.AppCommentList,res *map[int]tree,pid int) *map[int]tree {
	for k,v := range *list {
		if v.Pid == 0 {
			(*res)[v.Id] = tree{id: v.Id,pid: v.Pid,data: v,children: &[]tree{}}
			//delete this key
			*list = append((*list)[:k], (*list)[k+1:]...)
			getTree(list,res,v.Id)
		} else {

		}
	}
	return res
}


