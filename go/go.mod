module blog

go 1.13

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/jinzhu/gorm v1.9.14
	github.com/patrickmn/go-cache v2.1.0+incompatible
)

replace github.com/jinzhu/gorm v1.9.14 => github.com/go-gorm/gorm v1.9.12
