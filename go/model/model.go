package model

import (
	"blog/config"
	_"blog/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
var db *gorm.DB

func InitDb() *gorm.DB {
	conf := config.GetConf()
	var err error
	conn := conf.DB_username + ":" + conf.DB_password + "@/" + conf.DB_dbname + "?charset=utf8&parseTime=True&loc=Local"
	db, err = gorm.Open("mysql",conn)
	if err != nil {
		panic("failed to connect database")

	} else {
		db.SingularTable(true)
		db.LogMode(true)
	}

	return db
}

