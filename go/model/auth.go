package model

import (
	_"blog/config"
	"fmt"
	_"github.com/jinzhu/gorm"
)

type BlogAuth struct{
	Id uint `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
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


