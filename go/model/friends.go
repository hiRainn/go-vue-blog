package model

type BlogFriends struct {
	Id int `json:"id";gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	Name string `json:"name";gorm:"type:varchar(100);NOT NULL"`
	Link string `json:"link";gorm:"type:varchar(200);NOT NULL"`
	Avatar string `json:"avatar";gorm:"type:varchar(300);NOT NULL"`
	Sort uint8 `json:"sort";gorm:"type:tinyint;UNSIGNED;NOT NULL"`
}

func (f *BlogFriends) GetLink() ([]BlogFriends,error) {
	link := make([]BlogFriends,0)
	return link,db.Order("sort desc").Find(&link).Error
}

func (f *BlogFriends) AddFriend() (int,error) {
	return f.Id,db.Create(f).Error
}

func (f *BlogFriends) EditFriend() error {
	return db.Table("blog_friends").Where("id = ?",f.Id).Update(f).Error
}

func (f *BlogFriends) DelFriend() error {
	return db.Table("blog_friends").Where("id = ?" ,f.Id).Delete(f).Error
}
