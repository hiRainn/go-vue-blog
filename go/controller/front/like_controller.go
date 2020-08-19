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
	tx := model.DB().Begin()
	if id,err :=l.Add(tx); err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusOK,err.GetH())
	} else {
		//like comment
		if l.CommentId != 0 {
			var c model.BlogComment
			c.Id = l.CommentId
			if c.AddLikeNumber(tx) == false  {
				tx.Rollback()
				ctx.JSON(http.StatusOK,errcode.LikeError.GetH())
			} else {
				tx.Commit()
				ctx.JSON(http.StatusOK,errcode.Ok.SetData(id))
			}
		} else {
			//like article
			var a model.BlogArticle
			a.Id = l.ArticleId
			if a.AddLikeNumber(tx) == false {
				tx.Rollback()
				ctx.JSON(http.StatusOK,errcode.LikeError.GetH())
			} else {
				tx.Commit()
				ctx.JSON(http.StatusOK,errcode.Ok.SetData(id))
			}
		}
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
	tx := model.DB().Begin()
	if id,err :=l.Add(tx); err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusOK,err.GetH())
	} else {
		if l.CommentId != 0 {
			var c model.BlogComment
			c.Id = l.CommentId
			if c.AddUnlikeNumber(tx) == false {
				tx.Rollback()
				ctx.JSON(http.StatusOK,errcode.UnlikeError.GetH())
				return
			} else {
				tx.Commit()
				ctx.JSON(http.StatusOK,errcode.Ok.SetData(id))
				return
			}
		}
		tx.Commit()
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
	tx := model.DB().Begin()
	if id,err :=l.Add(tx); err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusOK,err.GetH())
	} else {
		if l.CommentId != 0 {
			var c model.BlogComment
			c.Id = l.CommentId
			if c.AddReportNumber(tx) == false {
				tx.Rollback()
				ctx.JSON(http.StatusOK,errcode.ReportError.GetH())
				return
			} else {
				tx.Commit()
				ctx.JSON(http.StatusOK,errcode.Ok.SetData(id))
				return
			}
		}
		tx.Commit()
		ctx.JSON(http.StatusOK,errcode.Ok.SetData(id))
	}
}

//discarded:  because there is no userId so can't determine whether it's the same person to do this
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

//discarded:  because there is no userId so can't determine whether it's the same person to do this
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

//discarded:  because there is no userId so can't determine whether it's the same person to do this
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
