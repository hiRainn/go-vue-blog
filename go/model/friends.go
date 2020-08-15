package model

type BlogFriends struct {
	Id int `json:"id";gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	Name string `json:"name";gorm:"type:varchar(100);NOT NULL"`
	Link string `json:"link";gorm:"type:varchar(200);NOT NULL"`
	Sort uint8 `json:"sort";gorm:"type:tinyint;UNSIGNED;NOT NULL"`
}

func (f *BlogFriends) GetLink() ([]BlogFriends,error) {
	link := make([]BlogFriends,0)
	return link,db.Order("sort desc").Find(&link).Error
}
