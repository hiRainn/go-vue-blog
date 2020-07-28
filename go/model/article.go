package model

import (
	"blog/pkg/errcode"
	"github.com/jinzhu/gorm"
)

type BlogArticle struct {
	Model
	Title string `json:"title";gorm:"type:varchar(100);DEFAULT:'';NOT NULL"`
	ModifyAt int64 `json:"modify_at";gorm:"type:int(10);default:0;not null"`
	CateId int `json:"cate_id";gorm:"type:int;not null;default:0"`
	Content string `json:"content";gorm:"type:varchar(5000);not null;default:''"`
	View int `json:"view";gorm:"type:int;not null;default:0"`
	IsTop uint8 `json:"is_top";gorm:"type:tinyint;unsigned;not null;default:0"`
	Sort uint8 `json:"sort";gorm:"type:tinyint;unsigned;not null;default:0"` // sort for articles  recommended
	TagsIds string `json:"tags_ids";gorm:"type varchar(100);not null;default:''"`

}

type AdminArticleList struct {
	Model
	Title string `json:"title"`
	CateName string `json:"cate_name"`
	Tags string `json:"tags"`
	Views int `json:"views"`
	Comments int `json:"comments"`
	NewViews uint16 `json:"new_views"`
	NewComments uint16 `json:"new_comments"`
	PostAt string `json:"post_at"`
}
const OFFSET = 15

//you must be in a transaction when post a article ,so must get handle of db
func(art *BlogArticle) AddArticle(tx *gorm.DB) (int,*errcode.ERRCODE){
	res := tx.Create(art)
	var e *errcode.ERRCODE
	id :=0
	if err:= res.Error;err != nil {
		e = errcode.AddArticleError
	} else {
		id = art.Id
	}
	return id,e
}

func (art *BlogArticle)  EditArticle(tx *gorm.DB) (bool,*errcode.ERRCODE) {
	if err := tx.Save(art).Error;err != nil {
		return false,errcode.EditSaveArticleError
	}
	return true,nil

}

func (art *BlogArticle) GetArticleById() *errcode.ERRCODE {
	res := db.Where("id = ?",art.Id).First(art)
	if res.Error != nil {
		if res.RecordNotFound() {
			return errcode.DataNotExists
		} else {
			return errcode.DataBaseError
		}
	}
	return nil
}

func (art *BlogArticle) GetArticleList () ([]AdminArticleList, error) {
	var list []AdminArticleList
	join_cate := "left join blog_cate bc on a.cate_id = bc.id"
	join_comment := "left join blog_comment c on a.id = c.article_id"
	join_tag := "left join blog_tags bt on FIND_IN_SET(bt.id,a.tags_ids)"
	join_view := "left join blog_view v on a.id = v.article_id"
	field := "a.id,a.title,bc.cate_name,group_concat(bt.tag_name) as tags,from_unixtime(a.created_at,'%Y-%m-%d %H:%i') as post_at,count(v.id) as views,count(c.id) as comments"
	res := db.Table("blog_article a").Select(field).Joins(join_cate).Joins(join_comment).Joins(join_tag).Joins(join_view).Group("a.id").Find(&list)
	return list, res.Error
}
