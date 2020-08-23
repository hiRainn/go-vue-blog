package model

import (
	"blog/pkg/errcode"
	"blog/pkg/status"
	"github.com/jinzhu/gorm"
)

type BlogArticle struct {
	Model
	Title string `json:"title";gorm:"type:varchar(100);DEFAULT:'';NOT NULL"`
	ModifyAt int64 `json:"modify_at";gorm:"type:int(10);default:0;not null"`
	CateId int `json:"cate_id";gorm:"type:int;not null;default:0"`
	Content string `json:"content";gorm:"type:text;not null;default:''"`
	IsTop uint8 `json:"is_top";gorm:"type:tinyint;unsigned;not null;default:0"`
	IsOriginal uint8 `json:"is_original";gorm:"type:tinyint;unsigned;not null;default:0"`
	ReprintFrom string `json:"reprint_from";gorm:"type varchar(200);not null;default:''"`
	AllowComment uint8 `json:"allow_comment";gorm:"type:tinyint;unsigned;not null;default:0"`
	IsSelf uint8 `json:"is_self";gorm:"type:tinyint;unsigned;not null;default:0"`
	Sort uint8 `json:"sort";gorm:"type:tinyint;unsigned;not null;default:0"` // sort for articles  recommended
	TagsIds string `json:"tags_ids";gorm:"type varchar(100);not null;default:''"`
	Views int `json:"views";gorm:"type:int;NOT NULL;DEFAULT:0"`
	Comments int `json:"comments";gorm:"type:int;NOT NULL;DEFAULT:0"`
	LikeNumber int `json:"like_number";gorm:"type:int;NOT NULL;DEFAULT:0"`


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
	IsSelf uint8 `json:"is_self"`
	IsTop uint8 `json:"is_top"`
	LikeNumber int `json:"like_number"`
	IsOriginal uint8 `json:"is_original"`
}

type AppArticleList struct {
	Model
	Title string `json:"title"`
	CateName string `json:"cate_name"`
	CateId uint8 `json:"cate_id"`
	TagName string `json:"tag_name"`
	PostAt string `json:"post_at"`
	ModifyAt string `json:"modify_at";gorm:"type:int(10);default:0;not null"`
	Content string `json:"content"`
	Views int `json:"views"`
	Comments int `json:"comments"`
	IsOriginal uint8 `json:"is_original"`
	IsTop uint8 `json:"is_top"`
	AllowComment uint8 `json:"allow_comment"`
	LikeNumber int `json:"like_number"`
	Like int8 `json:"like"`
}

type ClickMost struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Num int `json:"num"`
}

type TimeLine struct {
	Id int `json:"id"`
	Title string `json:"title"`
	CreatedAt string `json:"created_at"`
}

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

//admin get article to edit
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

	field := "a.id,a.title,bc.cate_name,bt.tag_name as tags"
	field = field + ",from_unixtime(a.created_at,'%Y-%m-%d %H:%i') as post_at"
	field = field + ",a.is_self,a.is_original,is_top,a.views,a.comments,a.like_number"
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
	res = res.Joins("left join (select a.id,group_concat(bt.tag_name) as tag_name from blog_article a left join blog_tags bt on find_in_set(bt.id,a.tags_ids) group by a.id) as bt on bt.id = a.id")

	res = res.Group("a.id")
	var count int
	res.Count(&count)
	res = res.Offset(offset).Limit(page["page_size"]).Order("a.id desc").Find(&list)

	return list,count,res.Error
}

func (art *BlogArticle) GetAppArticleList(condition map[string]interface{}, page map[string]int) ([]AppArticleList,int,error) {
	var list []AppArticleList
	field := "a.id,a.title,bc.cate_name,bt.tag_name,left(a.content,200) as content,a.is_original,is_top,a.cate_id"
	field = field + ",from_unixtime(a.created_at,'%Y-%m-%d %H:%i') as post_at,allow_comment,a.views,a.comments,a.like_number"
	field = field + ",from_unixtime(a.modify_at,'%Y-%m-%d %H:%i') as modify_at"
	offset := (page["page"] - 1) * page["page_size"]
	res := db.Table("blog_article a").Select(field)
	if condition["title"] != nil {
		res = res.Where("`title` like ?","%" + condition["title"].(string) + "%")
	}
	if condition["cate_id"] != nil {
		res = res.Where("`cate_id` = ?",condition["cate_id"].(int))
	}
	if condition["tag_id"] != nil {
		res = res.Where("find_in_set(?,tags_ids)",condition["tag_id"].(int))
	}
	res = res.Where("a.is_self = 0 and a.status = 0")
	res = res.Joins("left join blog_cate bc on a.cate_id = bc.id")
	res = res.Joins("left join (select a.id,group_concat(bt.tag_name) as tag_name from blog_article a left join blog_tags bt on find_in_set(bt.id,a.tags_ids) group by a.id) as bt on bt.id = a.id")
	res = res.Group("a.id")
	var count int
	res.Count(&count)
	res = res.Offset(offset).Limit(page["page_size"]).Order("a.is_top desc,a.sort desc,a.created_at desc").Find(&list)

	return list,count,res.Error
}

func (art *BlogArticle) GetAppHomeArticle() ([]AppArticleList,error){
	var list []AppArticleList
	field := "a.id,a.title,bc.cate_name,bt.tag_name,left(a.content,200) as content,a.is_original,is_top,a.cate_id"
	field = field + ",from_unixtime(a.created_at,'%Y-%m-%d %H:%i') as post_at,allow_comment,a.views,a.comments,a.like_number"
	field = field + ",from_unixtime(a.modify_at,'%Y-%m-%d %H:%i') as modify_at"
	res := db.Table("blog_article a").Select(field)

	res = res.Where("a.is_self = 0 and a.status = 0")
	res = res.Joins("left join blog_cate bc on a.cate_id = bc.id")
	res = res.Joins("left join (select a.id,group_concat(bt.tag_name) as tag_name from blog_article a left join blog_tags bt on find_in_set(bt.id,a.tags_ids) group by a.id) as bt on bt.id = a.id")
	res = res.Group("a.id")
	res = res.Limit(10).Order("a.created_at desc").Find(&list)

	return list,res.Error
}

func (art *BlogArticle) GetAppArticle(token string) (*AppArticleList,*errcode.ERRCODE) {
	app := new(AppArticleList)
	field := "a.id,a.title,bc.cate_name,bt.tag_name,a.content,a.is_original,is_top,a.cate_id"
	field = field + ",from_unixtime(a.created_at,'%Y-%m-%d %H:%i') as post_at,allow_comment"
	field = field + ",from_unixtime(a.modify_at,'%Y-%m-%d %H:%i') as modify_at,a.like_number,l.like"
	field = field + ",a.views,a.comments"
	res := db.Table("blog_article a").Select(field)
	res = res.Where("a.id = ? and is_self  = 0",art.Id)
	res = res.Joins("left join blog_cate bc on a.cate_id = bc.id")
	res = res.Joins("left join (select a.id,group_concat(bt.tag_name) as tag_name from blog_article a left join blog_tags bt on find_in_set(bt.id,a.tags_ids) group by a.id) as bt on bt.id = a.id")
	res = res.Joins("left join (select article_id,(select count(1) from blog_like where type = ? and token = ? and article_id = ? and comment_id = 0 and token <> '') as `like`  from blog_like where type = ? and article_id = ? and comment_id = 0 group by article_id) l on a.id = l.article_id",status.Like.GetCode(),token,art.Id,status.Like.GetCode(),art.Id)
	res = res.Group("a.id")
	if err := res.Find(app).Error; err != nil {
		return app,errcode.DataBaseError
	}
	return app,nil
}

func (art *BlogArticle) GetClickMost() ([]ClickMost,error) {
	most := make([]ClickMost,0)
	res := db.Table("blog_article a ").Select("a.id,a.title,count(v.id) as num")
	res = res.Joins("left join blog_view v on a.id = v.article_id").Where("a.is_self = 0")
	res = res.Order("num desc").Limit(10).Group("a.id").Find(&most)
	return most,res.Error
}

func (art *BlogArticle) GetAppArticleNum() (int,error) {
	res := 0
	return res,db.Table("blog_article").Where("status = 0 and is_self = 0").Count(&res).Error
}

func (art *BlogArticle) GetFirstArticle() error {
	return db.Order("created_at asc").Where("created_at <> 0 and is_self = 0").First(art).Error
}

func (art *BlogArticle) GetAllArticles() ([]TimeLine,error) {
	list := make([]TimeLine,0)
	res := db.Table("blog_article")
	res = res.Select("id,title,from_unixtime(created_at,'%Y-%m-%d %H:%i') as created_at")
	res = res.Where("is_self = 0 and status = 0")
	res = res.Order("created_at desc").Find(&list)
	return list,res.Error
}

func (art *BlogArticle) AddView() bool {
	if err := db.Table("blog_article").Where("id = ?", art.Id).UpdateColumn("views",gorm.Expr("views + ?",1)).Error; err !=nil {
		return false
	}
	return true
}

func (art *BlogArticle) AddLikeNumber(tx *gorm.DB) bool {
	if err := tx.Table("blog_article").Where("id = ?", art.Id).UpdateColumn("like_number",gorm.Expr("like_number + ?",1)).Error; err !=nil {
		return false
	}
	return true
}

func (art *BlogArticle) AddCommentNumber(tx *gorm.DB) bool {
	if err := tx.Table("blog_article").Where("id = ?", art.Id).UpdateColumn("comments",gorm.Expr("comments + ?",1)).Error; err !=nil {
		return false
	}
	return true
}

type Self struct {
	Id int `json:"id"`
	CateId int `json:"cate_id"`
}

func (art *BlogArticle) GetSelfArticle() []Self {
	self := make([]Self,0)
	db.Table("blog_article").Where("is_self = 1").Find(&self)
	return self
}
