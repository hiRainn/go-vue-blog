package front

import (
	"blog/model"
	"blog/pkg/errcode"
	"blog/pkg/status"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)


func Like(ctx *gin.Context) {
	var params map[string] interface{}
	if err := ctx.BindJSON(&params); err != nil {
		ctx.JSON(http.StatusOK,errcode.ParamError.GetH())
		return
	}
	article_id,_ := params["article_id"].(float64)
	comment_id,_:= params["comment_id"].(float64)
	if article_id == 0.00 && comment_id == 0.00 {
		ctx.JSON(http.StatusOK,errcode.ParamError.GetH())
		return
	}
	var l model.BlogLike
	l.ArticleId = int(article_id)
	l.CommentId = int(comment_id)
	l.Token = ctx.GetHeader("X-Token")
	l.Type = status.Like.GetCode()

	//repeat operate
	if l.GetUserOperate() {
		ctx.JSON(http.StatusOK,errcode.LikeRepeat.GetH())
		return
	}
	l.CreatedAt = time.Now().Unix()
	if id,err :=l.Add(); err != nil {
		ctx.JSON(http.StatusOK,err.GetH())
	} else {
		ctx.JSON(http.StatusOK,errcode.Ok.SetData(id))
	}

}

func DisLike(ctx *gin.Context) {
	var params map[string] interface{}
	if err := ctx.BindJSON(&params); err != nil {
		ctx.JSON(http.StatusOK,errcode.ParamError.GetH())
		return
	}
	article_id,_ := params["article_id"].(float64)
	comment_id,_:= params["comment_id"].(float64)
	if article_id == 0.00 && comment_id == 0.00 {
		ctx.JSON(http.StatusOK,errcode.ParamError.GetH())
		return
	}
	var l model.BlogLike
	l.ArticleId = int(article_id)
	l.CommentId = int(comment_id)
	l.Token = ctx.GetHeader("X-Token")
	l.Type = status.Dislike.GetCode()

	//repeat operate
	if l.GetUserOperate() {
		ctx.JSON(http.StatusOK,errcode.DislikeRepeat.GetH())
		return
	}

	l.CreatedAt = time.Now().Unix()
	if id,err :=l.Add(); err != nil {
		ctx.JSON(http.StatusOK,err.GetH())
	} else {
		ctx.JSON(http.StatusOK,errcode.Ok.SetData(id))
	}
}

func Report(ctx *gin.Context) {
	var params map[string] interface{}
	if err := ctx.BindJSON(&params); err != nil {
		ctx.JSON(http.StatusOK,errcode.ParamError.GetH())
		return
	}
	article_id,_ := params["article_id"].(float64)
	comment_id,_:= params["comment_id"].(float64)
	if article_id == 0.00 && comment_id == 0.00 {
		ctx.JSON(http.StatusOK,errcode.ParamError.GetH())
		return
	}
	var l model.BlogLike
	l.ArticleId = int(article_id)
	l.CommentId = int(comment_id)
	l.Token = ctx.GetHeader("X-Token")
	l.Type = status.Report.GetCode()

	//repeat operate
	if l.GetUserOperate() {
		ctx.JSON(http.StatusOK,errcode.ReportRepeat.GetH())
		return
	}

	l.CreatedAt = time.Now().Unix()
	if id,err :=l.Add(); err != nil {
		ctx.JSON(http.StatusOK,err.GetH())
	} else {
		ctx.JSON(http.StatusOK,errcode.Ok.SetData(id))
	}
}

func CancleLike(ctx *gin.Context) {
	var params map[string] interface{}
	if err := ctx.BindJSON(&params); err != nil {
		ctx.JSON(http.StatusOK,errcode.ParamError.GetH())
		return
	}
	article_id,_ := params["article_id"].(float64)
	comment_id,_:= params["comment_id"].(float64)
	if article_id == 0.00 && comment_id == 0.00 {
		ctx.JSON(http.StatusOK,errcode.ParamError.GetH())
		return
	}
	var l model.BlogLike
	l.ArticleId = int(article_id)
	l.CommentId = int(comment_id)
	l.Token = ctx.GetHeader("X-Token")
	l.Type = status.Like.GetCode()


	if l.GetUserOperate() {
		l.Delete()
		ctx.JSON(http.StatusOK,errcode.Ok.SetData(l.Id))
	} else {
		//fake operate
		ctx.JSON(http.StatusOK,errcode.Ok.SetData(l.Id))
	}

}

func CancleDislile(ctx *gin.Context) {
	var params map[string] interface{}
	if err := ctx.BindJSON(&params); err != nil {
		ctx.JSON(http.StatusOK,errcode.ParamError.GetH())
		return
	}
	article_id,_ := params["article_id"].(float64)
	comment_id,_:= params["comment_id"].(float64)
	if article_id == 0.00 && comment_id == 0.00 {
		ctx.JSON(http.StatusOK,errcode.ParamError.GetH())
		return
	}
	var l model.BlogLike
	l.ArticleId = int(article_id)
	l.CommentId = int(comment_id)
	l.Token = ctx.GetHeader("X-Token")
	l.Type = status.Dislike.GetCode()

	//
	if l.GetUserOperate() {
		l.Delete()
		ctx.JSON(http.StatusOK,errcode.Ok.SetData(l.Id))
	} else {
		//fake operate
		ctx.JSON(http.StatusOK,errcode.Ok.SetData(l.Id))
	}
}

func CancleReport(ctx *gin.Context) {
	var params map[string] interface{}
	if err := ctx.BindJSON(&params); err != nil {
		ctx.JSON(http.StatusOK,errcode.ParamError.GetH())
		return
	}
	article_id,_ := params["article_id"].(float64)
	comment_id,_:= params["comment_id"].(float64)
	if article_id == 0.00 && comment_id == 0.00 {
		ctx.JSON(http.StatusOK,errcode.ParamError.GetH())
		return
	}
	var l model.BlogLike
	l.ArticleId = int(article_id)
	l.CommentId = int(comment_id)
	l.Token = ctx.GetHeader("X-Token")
	l.Type = status.Report.GetCode()

	//
	if l.GetUserOperate() {
		l.Delete()
		ctx.JSON(http.StatusOK,errcode.Ok.SetData(l.Id))
	} else {
		//fake operate
		ctx.JSON(http.StatusOK,errcode.Ok.SetData(l.Id))
	}
}
