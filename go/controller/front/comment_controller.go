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
		//var res []tree
		//getTree(&list,&res,0)
		//func gettree2 faster then gettree
		test := make(map[int]model.AppCommentList)
		for k,v := range list {
			test[k] = v
		}
		var res2 []tree
		getTree2(&test,&res2,0)
		//ctx.JSON(http.StatusOK,errcode.Ok.SetData(map[string]interface{}{"list":res,"total":total,"counter1":counter1,"counter2":counter2}))
		ctx.JSON(http.StatusOK,errcode.Ok.SetData(map[string]interface{}{"list":res2,"total":total}))
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

var counter1 int = 0
var counter2 int = 0
type tree struct {
	Data model.AppCommentList `json:"data"`
	Children *[]tree `json:"children"`
}
func getTree(list *[]model.AppCommentList,res *[]tree,pid int) {
	for _,v := range *list {
		if v.Pid == pid {
			data := tree{Data: v, Children: &[]tree{}}
			*res = append(*res, data)
			getTree(list, data.Children, v.Id)
		}
		counter1++
	}
}

func getTree2(list *map[int]model.AppCommentList,res *[]tree,pid int) {
	for k,v := range *list {
		if v.Pid == pid {
			data := tree{Data: v, Children: &[]tree{}}
			*res = append(*res, data)
			delete(*list,k)
			getTree2(list, data.Children, v.Id)
		}
		counter2 ++
	}
}


