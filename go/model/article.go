package model

type BlogArticle struct {
	Model
	Title string `json:"title";gorm:"type:varchar(100);DEFAULT:'';NOT NULL"`
	Abstract string `json:"abstract";gorm:"type:varchar(500);default:'';"`
	ModifyAt uint `json:"modify_at";gorm:"type:int;default:0"`
	Content string `json:"content";gorm:"type:varchar(5000)"`
}
const OFFSET = 15

func AddArticle(article *BlogArticle) error{
	 res := db.Create(article)
	 return res.Error
}

func GetArticleListByPage (page int) (BlogArticle, error) {
	if page <= 1 {
		page = 1
	} else {
		page = page -1
	}
	offset := OFFSET * page
	var article BlogArticle
	res := db.Offset(offset).Limit(OFFSET).Order("created_at desc").Find(&article)
	return article, res.Error
}
