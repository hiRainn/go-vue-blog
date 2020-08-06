module blog

go 1.13

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/jinzhu/gorm v1.9.14
	github.com/patrickmn/go-cache v2.1.0+incompatible
	gopkg.in/TomatoMr/SensitiveWords.v1 v1.0.0-20180624094620-0b5f75ba0d56
)

replace github.com/jinzhu/gorm v1.9.14 => github.com/go-gorm/gorm v1.9.12
