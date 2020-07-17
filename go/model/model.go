package model

import (
	"blog/config"
	_"blog/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
var db *gorm.DB

type Model struct{
	Id uint `json:"id";gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	Status uint `json:"status";gorm:"DEFAULT:0;type:tinyint;NOT NULL"`
	CreatedAt int64 `json:"create_at";gorm:"DEFAULT:0;NOT NULL;type:int(10)"`
}

func init() {
	conf := config.GetConf()
	var err error
	conn := conf.DB_username + ":" + conf.DB_password + "@/" + conf.DB_dbname + "?charset=utf8&parseTime=True&loc=Local"
	db, err = gorm.Open("mysql",conn)
	if err != nil {
		panic("failed to connect database")

	} else {
		db.SingularTable(true)
		db.AutoMigrate(&BlogArticle{},&BlogAuth{},&BlogConfig{},&BlogCate{})
		db.LogMode(true)
	}
}

func DB() *gorm.DB {
	return db
}


