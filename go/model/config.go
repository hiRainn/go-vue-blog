package model

type BlogConfig struct {
	Id uint8 `json:"id";gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL;type:tinyint;UNSIGNED"`
	Key string `json:"key";gorm:"type:varchar(32);NOT NULL;DEFAULT:''"`
	Value string `json:"value";gorm:"type:varchar(64);NOT NULL;DEFAULT:''"`
}

func (conf *BlogConfig) Add() uint8 {
	var key uint8 = 0
	if db.Create(conf).Error == nil {
		key = conf.Id
	}
	return key
}

func (conf *BlogConfig) Find() ([]BlogConfig, error ) {
	var config []BlogConfig
	res := db.Find(&config)
	return config, res.Error
}

func (conf *BlogConfig) Update() {
	db.Table("blog_config").Where("`key` = ?",conf.Key).Update("value",conf.Value)
}

//determine if there is date in table
func (conf *BlogConfig) HasData() (bool,error) {
	var config BlogConfig
	res := db.First(&config)
	return res.RecordNotFound(),res.Error
}