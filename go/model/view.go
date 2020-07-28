package model

type BlogView struct {
	Model
	ArticleId int `json:"article_id";gorm:"type:int;unsigned;not null;default:0"`
}
