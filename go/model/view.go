package model

type BlogView struct {
	Model
	ArticleId int `json:"article_id";gorm:"type:int;unsigned;not null;default:0"`
	Ip string `json:"ip";gorm:"type:varchar(64);default:'';not null;"`
	Module string `json:"module";gorm:"type:varchar(32);default:'';not null"`
	Token string `json:"token";gorm:"type:char(64);not null;default:''"`
}

func (v *BlogView) AddRecord() {
	db.Create(v)
}

func (v *BlogView) FindRecord() {
	db.Where(v).Order("created_at desc").First(v)
}

func (v *BlogView) StatViews(timestamp int) (int,error) {
	var views int
	return views,db.Table("blog_view").Where("created_at >= ?",timestamp).Count(&views).Error
}
