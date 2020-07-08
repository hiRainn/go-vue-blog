package config

type Conf struct {
	//database config
	DB_username string
	DB_password string
	DB_dbname string
	DB_port string
	DB_prefix string
	Salt string
}

var conf Conf

func init() {
	conf.DB_username = "root"
	conf.DB_password = "0325"
	conf.DB_dbname = "blog"
	conf.DB_port = "3306"
	conf.DB_prefix = "blog"
	conf.Salt = "@9d*1md103"
}

func SetConf(key string,val string) Conf {
	switch key {
	case "db_username" : conf.DB_username = val
	case "db_password" : conf.DB_password = val
	case "db_dbname" : conf.DB_dbname = val
	case "db_port" : conf.DB_port = val
	case "db_prefix" : conf.DB_prefix = val
	case "salt" : conf.Salt = val
	}

	return conf
}

func GetConf() Conf {
	return conf
}


