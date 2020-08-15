package model

import (
	"blog/pkg/errcode"
	"github.com/jinzhu/gorm"
)

type BlogTags struct {
	Id int `json:"id";gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL;type:tinyint;UNSIGNED"`
	TagName string `json:"tag_name";gorm:"type:varchar(32);NOT NULL;DEFAULT:''"`
	Num uint8 `json:"num";gorm:"type:smallint(3);unsigned;NOT NULL;DEFAULT:0"`
}

type SelectTags struct {
	Id int	`json:"id"`
	Name string `json:"name"`
	Num int `json:"num"`
}

func (t *BlogTags) GetInfo() *errcode.ERRCODE{
	if db.Where(t).First(t).Error != nil {
		return errcode.ParamError
	}
	return nil
}

//get all tags for select
func (t *BlogTags) GetSelectList() ([]SelectTags,*errcode.ERRCODE) {
	var list []SelectTags
	var err *errcode.ERRCODE
	res := db.Table("blog_tags").Select("id,tag_name as name,num").Order("num desc").Find(&list)
	if res.Error != nil {
		if res.RecordNotFound() {
			err = nil
		} else {
			err = errcode.DataBaseError
		}
	}
	return list,err
}

//determine whether category's is repeat
func (t *BlogTags) CheckNameRepeat() (bool,int) {
	if err := db.Where("tag_name = ?",t.TagName).First(t).Error; err != nil {
		return false,0
	}
	return true,t.Id
}
//determine whether article's number of this tag is 0
func (t *BlogTags) CheckNumEmpty(tag_name string) (bool,*errcode.ERRCODE) {
	var tag *BlogTags
	result := false
	var err *errcode.ERRCODE
	res := db.Where("name = ?", tag_name).First(tag)
	if res.Error != nil {
		if res.RecordNotFound() {
			err = errcode.DataNotExists
		}
		err = errcode.DataBaseError
	} else {
		if tag.Num == 0 {
			result = true
		}
	}
	return result, err
}

//add tag ，
func (t *BlogTags) AddCate(tx *gorm.DB) int {
	res := 0
	if err :=tx.Create(t).Error; err == nil {
		res =  t.Id
	}
	return res
}

//add 1 when a article choose these tags
func (c *BlogTags) SetIncNum(tx *gorm.DB,ids []int) bool {
	if err := tx.Table("blog_tags").Where("id in (?)", ids).UpdateColumn("num",gorm.Expr("num + ?",1)).Error; err !=nil {
		return false
	}
	return true
}
//reduce 1 when a article choose these tags
func (c *BlogTags) SetDecNum(tx *gorm.DB,ids []string) bool {
	if err := tx.Table("blog_tags").Where("id in (?)", ids).UpdateColumn("num",gorm.Expr("num - ?",1)).Error; err !=nil {
		return false
	}
	return true
}