package model

import (
	"blog/pkg/errcode"
	"blog/pkg/status"
)

type BlogComment struct {
	Model
	ArticleId int `json:"article_id";gorm:"type:int;unsigned;not null;default:0"`
	Name string `json:"name";gorm:"type:varchar(64);not null;default:''"`
	Email string `json:"email";gorm:"type:varchar(128);not null;default:''"`
	Pid int `json:"pid";gorm:"type:int;unsigned;not null;default:0"`
	Content string `json:"content";gorm:"type:varchar(3000);not null;default:''"`
	IsView uint8 `json:"is_view";gorm:"type:tinyint;unsigned;not null;default:0"`
	IsAdmin uint8 `json:"is_admin";gorm:"type tinyint;unsigned;not null;default:0"`
	Level uint8 `json:"level";gorm:"type:tinyint;unsigned;not null;default:0"`
	Token string `json:"token";gorm:"type:char(64);not null;default:''"`
}

type AppCommentList struct {
	Id int `json:"id"`
	Status uint8 `json:"status"`
	Content string `json:"content"`
	Name string `json:"name"`
	Pid int `json:"pid"`
	CreatedAt int `json:"created_at"`
	IsAdmin uint8 `json:"is_admin"`
	Level uint8 `json:"level"`
}

//app get comment list
func (c *BlogComment) GetCommentsByArticleId(token string) ([]AppCommentList,int,*errcode.ERRCODE) {
	var app []AppCommentList
	var number int
	res := db.Table("blog_comment").Where("article_id = ?",c.ArticleId)
	if token == "" {
		res = res.Where("status = ?",0)
	} else {
		 want_to_see := []int8{int8(status.OK.GetCode()),int8(status.CommentCheck.GetCode()),int8(status.CommentNotPass.GetCode())}
		 res = res.Where("(status in (?) and token = ?) or status = 0",want_to_see,token)
	}

	res.Count(&number)
	if res.Find(&app).Error != nil {
		return app,0,errcode.GetCommentsError
	} else {
		return app,number,nil
	}
}

func (c *BlogComment) AddComment() *errcode.ERRCODE {
	if db.Create(c).Error != nil {
		return errcode.CommentError
	}
	return nil
}
