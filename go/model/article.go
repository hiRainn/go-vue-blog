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

func GetArticleListByPage (page int) ([]BlogArticle, error) {
	if page <= 1 {
		page = 0
	} else {
		page = page -1
	}
	offset := OFFSET * page
	var article []BlogArticle
	res := db.Offset(offset).Limit(OFFSET).Order("id desc").Find(&article)
	return article, res.Error
}
