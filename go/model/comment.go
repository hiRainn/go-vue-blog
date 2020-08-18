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
	FloorId int `json:"floor_id";gorm:"type:int;unsigned;not null;default:0"`
	Content string `json:"content";gorm:"type:varchar(3000);not null;default:''"`
	IsView uint8 `json:"is_view";gorm:"type:tinyint;unsigned;not null;default:0"`
	AdminView uint8 `json:"admin_view";gorm:"type:tinyint;unsigned;not null;default:0"`
	IsAdmin uint8 `json:"is_admin";gorm:"type tinyint;unsigned;not null;default:0"`
	Token string `json:"token";gorm:"type:char(64);not null;default:''"`
	LikeNumber int `json:"like_number";gorm:"type:int;NOT NULL;DEFAULT:0"`
	UnlikeNumber int `json:"unlike_number";gorm:"type:int;NOT NULL;DEFAULT:0"`
	ReportNumber int `json:"report_number";gorm:"type:int;NOT NULL;DEFAULT:0"`
}

type AppCommentList struct {
	Id int `json:"id"`
	Status uint8 `json:"status"`
	Content string `json:"content"`
	Name string `json:"name"`
	FloorId int `json:"floor_id"`
	Pid int `json:"pid"`
	CreatedAt string `json:"created_at"`
	IsAdmin uint8 `json:"is_admin"`
	LikeNumber int `json:"like_number"`
	UnlikeNumber int `json:"unlike_number"`
	Like uint8 `json:"like"`
	Unlike uint8 `json:"unlike"`
	Report uint8 `json:"report"`
}

type LatestComment struct {
	Id int `json:"id"`
	ArticleId int `json:"article_id"`
	ArticleTitle string `json:"article_title"`
	Content string `json:"content"`
	Name string `json:"name"`
}

//app get comment list
func (c *BlogComment) GetCommentsByArticleId(token string) ([]AppCommentList,int,*errcode.ERRCODE) {
	var app []AppCommentList
	var number int
	//comments check status is not allowed to see,but reported or deleted are allowed cause maybe they are replayed
	res := db.Table("blog_comment c")
	res = res.Select("from_unixtime(c.created_at,'%Y-%m-%d %H:%i') as created_at,c.*,l.like_number,l.like,d.unlike_number,d.unlike,r.report")
	res = res.Where("c.article_id = ?",c.ArticleId)
	no_see := status.CommentCheck.GetCode()
	if token == "" {
		res = res.Where("c.status <> ?",no_see)
	} else {
		res = res.Where("c.token = ? or c.status <> ?",token, no_see)
	}
	res = res.Joins("left join (select comment_id,count(type) as like_number,(select count(1) from blog_like where type = ? and token = ? and token <> '') as `like`  from blog_like where type = ? group by comment_id) l on c.id = l.comment_id",status.Like.GetCode(),token,status.Like.GetCode())
	res = res.Joins("left join (select comment_id,count(type) as unlike_number,(select count(1) from blog_like where type = ? and token = ? and token <> '') as `unlike` from blog_like where type = ? group by comment_id) d on c.id = d.comment_id",status.Dislike.GetCode(),token,status.Dislike.GetCode())
	res = res.Joins("left join (select comment_id,(select count(1) from blog_like where type = ? and token = ? and token <> '') as `report` from blog_like where type = ? group by comment_id) r on c.id = r.comment_id",status.Report.GetCode(),token,status.Report.GetCode())
	res.Count(&number)
	if res.Find(&app).Error != nil {
		return app,0,errcode.GetCommentsError
	} else {
		return app,number,nil
	}
}

func (c *BlogComment) GetFloorIdByPid(pid int) (int,*errcode.ERRCODE) {
	var comment BlogComment
	if db.Where("id = ?",pid).Find(&comment).Error != nil {
		return 0,errcode.DataBaseError
	}
	if comment.FloorId != 0 {
		return comment.FloorId,nil
	}
	return comment.Id,nil
}

func (c *BlogComment) AddComment() *errcode.ERRCODE {
	if db.Create(c).Error != nil {
		return errcode.CommentError
	}
	return nil
}

func (c * BlogComment) GetLastComment() ([]LatestComment,error) {
	latest := make([]LatestComment,0)
	res := db.Table("blog_comment c")
	res = res.Select("c.id,a.id as article_id,a.title as article_title,c.name,left(c.content,70) as content")
	res = res.Joins("left join blog_article a on c.article_id = a.id")
	res = res.Where("c.status = 0 and a.status = 0")
	res = res.Limit(10)
	res = res.Order("c.id desc").Find(&latest)
	return latest,res.Error
}

func (c *BlogComment) GetCommentNumber() (int,error) {
	var res int
	return res,db.Table("blog_comment").Where("status <> ? and article_id <> 0",status.CommentCheck.GetCode()).Count(&res).Error
}

func (c *BlogComment) GetMsgNumber() (int,error) {
	var res int
	return res,db.Table("blog_comment").Where("status <> ? and article_id = 0",status.CommentCheck.GetCode()).Count(&res).Error
}
