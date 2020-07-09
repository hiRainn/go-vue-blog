package model

import (
	_ "blog/config"
	"blog/pkg/errcode"
	_ "github.com/jinzhu/gorm"
)

type BlogAuth struct{
	Id uint `json:"id";gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL;type:tinyint;UNSIGNED"`
	Username string `json:"username";gorm:"type:varchar(32);NOT NULL;DEFAULT:''"`
	Password string `json:"password";gorm:"type:char(64);NOT NULL;DEFAULT:''"`
	Nickname string `json:"nickname";gorm:"type:varchar(32);NOT NULL;DEFAULT:''"`
	Birthday int64 `json:"birthday";;gorm:"type:smallint(3);NOT NULL;DEFAULT:0"`
	Intro string `json:"intro";gorm:"type:varchar(500);NOT NULL;DEFAULT:''"`
}

//check auth
func (auth *BlogAuth) CheckAuth() *errcode.ERRCODE {
	res := db.Where(&auth).First(&auth)
	result := &errcode.ERRCODE{}
	if res.Error != nil  {
		if res.RecordNotFound() {
			result = errcode.PasswordError
		} else {
			result = errcode.DataBaseError
		}
	} else {
		result = nil
	}
	return result
}


func (auth *BlogAuth) InsertAuth() (uint, error)  {
	res := db.Create(&auth)
	return auth.Id,res.Error
}


