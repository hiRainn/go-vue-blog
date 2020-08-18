package model

import (
	"blog/pkg/errcode"
	"blog/pkg/status"
)

type BlogLike struct {
	Model
	ArticleId int `json:"article_id";gorm:"type:int;NOT NULL;DEFAULT:0"`
	CommentId int `json:"comment_id";gorm:"type:int;NOT NULL;DEFAULT:0"`
	Type uint8 `json:"type";gorm:"type:tinyint;NOT NULL;DEFAULT:0"`
	Token string `json:"token";gorm:"type:varchar(64);NOT NULL;DEFAULT:0"`
}

func (l *BlogLike) GetIndexLike() (int) {
	like := status.Like.GetCode()
	var res int
	db.Raw("select (select count(1) from blog_like where comment_id = 0 and type = ?)  + (select count(1) from blog_like l left join blog_comment c on l.comment_id = c.id  where comment_id <> 0 and c.is_admin = 1 and type = ?)",like,like).Count(&res)
	return res
}

func (l *BlogLike) GetUserOperate() bool {
	if l.Token == "" {
		return false
	}
	if db.Where("token = ? and type = ? and article_id = ? and comment_id = ?",l.Token , l.Type, l.ArticleId, l.CommentId).First(l).RecordNotFound() {
		return false
	}

	return true
}

func (l *BlogLike) Add() (int,*errcode.ERRCODE) {
	res := db.Create(l)
	if res.Error != nil {
		return 0,errcode.LikeError
	}
	return l.Id,nil
}

func (l *BlogLike) Delete()  {
	db.Delete(l)
}
