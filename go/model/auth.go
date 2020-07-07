package model

import (
	_"blog/config"
	"fmt"
	_"github.com/jinzhu/gorm"
)

type BlogAuth struct{
	Id uint `json:"id";gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL;type:tinyint;UNSIGNED"`
	Username string `json:"username";gorm:"type:varchar(32);NOT NULL;DEFAULT:''"`
	Password string `json:"password";gorm:"type:char(64);NOT NULL;DEFAULT:''"`
	Nickname string `json:"nickname";gorm:"type:varchar(32);NOT NULL;DEFAULT:''"`
	Birthday int64 `json:"birthday";;gorm:"type:smallint(3);NOT NULL;DEFAULT:0"`
	Intro string `json:"intro";gorm:"type:varchar(500);NOT NULL;DEFAULT:''"`
}

func CheckAuth(auth BlogAuth) bool {
	res := db.Where(&auth).First(&auth)
	fmt.Println(res)
	result := true
	if res.RecordNotFound()  {
		result = false
	}
	return result
}

func GetList()  BlogAuth {
	var auth BlogAuth
	db.Find(&auth)
	return auth
}

func InsertAuth(auth BlogAuth) {
	db.Create(&auth)
}


