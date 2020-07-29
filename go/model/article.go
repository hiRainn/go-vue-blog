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
	IsSelf uint8 `json:"is_self";gorm:"type:tinyint;unsigned;not null;default:0"`
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
	res := db.First(art)
	if res.Error != nil {
		if res.RecordNotFound() {
			return errcode.DataNotExists
		} else {
			return errcode.DataBaseError
		}
	}
	return nil
}

func (art *BlogArticle) DelArticleById() *errcode.ERRCODE {
	res := db.Where("id = ?",art.Id).Delete(art)
	if res.Error != nil {
		return errcode.DeleteArticleError
	}
	return nil
}

func (art *BlogArticle) GetArticleList (condition map[string]interface{}, page map[string]int) ([]AdminArticleList, int, error) {
	var list []AdminArticleList

	field := "a.id,a.title,bc.cate_name,group_concat(bt.tag_name) as tags,from_unixtime(a.created_at,'%Y-%m-%d %H:%i') as post_at,count(v.id) as views,count(c.id) as comments"
	offset := (page["page"] - 1) * page["page_size"]
	res := db.Table("blog_article a").Select(field)
	if condition["title"] != nil {
		res = res.Where("`title` like ?","%" + condition["title"].(string) + "%")
	}
	if condition["cate_id"] != nil {
		res = res.Where("`cate_id` = ?",condition["cate_id"].(int))
	}
	if condition["tag_id"] != nil {
		for _,v := range condition["tag_id"].([]string) {
			res = res.Where("find_in_set(?,tags_ids)",v)
		}
	}
	res = res.Where("a.status = ?", condition["status"].(int))
	res = res.Joins("left join blog_cate bc on a.cate_id = bc.id")
	res = res.Joins("left join blog_comment c on a.id = c.article_id")
	res = res.Joins("left join blog_tags bt on FIND_IN_SET(bt.id,a.tags_ids)")
	res = res.Joins("left join blog_view v on a.id = v.article_id")
	res = res.Group("a.id")
	var count int
	res.Count(&count)
	res = res.Offset(offset).Limit(page["page_size"]).Order("a.id desc").Find(&list)

	return list,count,res.Error
}
