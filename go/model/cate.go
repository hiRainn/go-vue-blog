package model

import (
	"blog/pkg/errcode"
	"github.com/jinzhu/gorm"
)

type BlogCate struct {
	Id int `json:"id";gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL;type:tinyint;UNSIGNED"`
	CateName string `json:"cate_name";gorm:"type:varchar(32);NOT NULL;DEFAULT:''"`
	Num uint8 `json:"num";gorm:"type:smallint(3);unsigned;NOT NULL;DEFAULT:0"`
}

type SelectCate struct {
	Id int	`json:"id"`
	Name string `json:"name"`
}

//get tag_list by condition
func (c *BlogCate) GetList(condition map[string] interface{}) ([]*BlogCate,*errcode.ERRCODE) {
	var list []*BlogCate
	var err *errcode.ERRCODE
	res := db.Where(condition).Find(list)
	if res.Error != nil {
		if res.RecordNotFound() {
			err = nil
		} else {
			err = errcode.DataBaseError
		}
	}
	return list,err
}
//get all categories for select
func (c *BlogCate) GetSelectList() ([]*SelectCate,*errcode.ERRCODE) {
	var list []*SelectCate
	var err *errcode.ERRCODE
	res := db.Table("blog_cate").Select("id,cate_name as name").Order("id desc").Find(&list)
	if res.Error != nil {
		if res.RecordNotFound() {
			err = nil
		} else {
			err = errcode.DataBaseError
		}
	}
	return list,err
}
//add 1 when a article choose this category
func (c *BlogCate) SetIncNum(tx *gorm.DB) bool {
	if err := tx.Model(c).Where("id = ?", c.Id).UpdateColumn("num",gorm.Expr("num + ?",1)).Error; err !=nil {
		return false
	}
	return true
}
//reduce 1 when a article choose this category
func (c *BlogCate) SetDecNum(tx *gorm.DB) bool {
	if err := tx.Model(c).Where("id = ?", c.Id).UpdateColumn("num",gorm.Expr("num - ?",1)).Error; err !=nil {
		return false
	}
	return true
}

//determine whether category's is repeat
func (c *BlogCate) CheckNameRepeat() bool {
	if err := db.Where("cate_name = ?",c.CateName).First(c).Error; err != nil {
		return false
	}
	return true
}
//add category
func (c *BlogCate) AddCate(cate *BlogCate) int {
	res := 0
	if err :=db.Create(cate).Error; err == nil {
		res =  cate.Id
	}
	return res
}
