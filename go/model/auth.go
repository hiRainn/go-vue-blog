package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type AUTH struct{
	gorm.Model
	Id int `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetList() AUTH {
	var auth AUTH
	db.First(&auth)
	return auth
}

func init()  {
	fmt.Println("123")
}

