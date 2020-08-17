package model

import "blog/pkg/status"

type BlogLike struct {
	Model
	ArticleId int `json:"article_id"`
	CommentId int `json:"comment_id"`
	Type uint8 `json:"type"`
	Token string `json:"token"`
}

func (l *BlogLike) GetIndexLike() (int) {
	like := status.Like.GetCode()
	var res int
	db.Raw("select (select count(1) from blog_like where comment_id = 0 and type = ?)  + (select count(1) from blog_like l left join blog_comment c on l.comment_id = c.id  where comment_id <> 0 and c.is_admin = 1 and type = ?)",like,like).Count(&res)
	return res
}
