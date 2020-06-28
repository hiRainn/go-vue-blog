package model

import (
	"blog/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
var db *gorm.DB
func init() {
	fmt.Printf("model_init")
	conf := config.GetConf()
	conn := conf.DB_username + conf.DB_password + "@/" + conf.DB_dbname + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql",conn)
	if err != nil {
		defer db.Close()
	} else {
		fmt.Println(err)
	}
}

