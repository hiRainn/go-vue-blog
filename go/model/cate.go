package model

import "github.com/jinzhu/gorm"

type BlogCate struct {
	Id int `json:"id";gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL;type:tinyint;UNSIGNED"`
	CateName string `json:"cate_name";gorm:"type:varchar(32);NOT NULL;DEFAULT:''"`
	Num uint8 `json:"num";gorm:"type:smallint(3);unsigned;NOT NULL;DEFAULT:0"`
}

func (c *BlogCate) AddNUm(id int) bool {
	if err := db.Model(c).Where("id = ?", id).UpdateColumn("num",gorm.Expr("num + ?",1)).Error; err !=nil {
		return false
	}
	return false
}

func (c *BlogCate) CheckNameRepeat() bool {
	if err := db.First(c).Error; err != nil {
		return true
	}
	return false
}

func (c *BlogCate) AddCate(cate BlogCate) int {
	res := 0
	if err :=db.Create(cate).Error; err != nil {
		res =  cate.Id
	}
	return res
}
