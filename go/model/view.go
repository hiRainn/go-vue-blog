package model

type BlogView struct {
	Model
	ArticleId int `json:"article_id";gorm:"type:int;unsigned;not null;default:0"`
	Ip string `json:"ip";gorm:"type:varchar(64);default:'';not null;"`
	Module string `json:"module";gorm:"type:varchar(32);default:'';not null"`
}

func (v *BlogView) AddRecord() {
	db.Create(v)
}

func (v *BlogView) FindRecord() {
	db.Where(v).Order("created_at desc").First(v)
}
