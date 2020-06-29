package model

type BlogArticle struct {
	Model
	Title string `json:"title";gorm:"type:varchar(100);DEFAULT:'';NOT NULL"`
	Abstract string `json:"abstract";gorm:"type:varchar(500);default:'';"`
	Content string `json:"content";gorm:"type:varchar(5000)"`
}

func GetArticleByPage (page int) *BlogAuth {
	
}
