package model

type BlogComment struct {
	Model
	ArticleId int `json:"article_id";gorm:"type:int;unsigned;not null;default:0"`
	Pid int `json:"pid";gorm:"type:int;unsigned;not null;default:0"`
	Content string `json:"content";gorm:"type:varchar(500);not null;default:''"`
	IsView uint8 `json:"is_view";gorm:"type:tinyint;unsigned;not null;default:0"`
}
