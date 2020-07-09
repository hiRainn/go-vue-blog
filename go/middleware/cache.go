package middleware

import (
	"blog/config"
	"github.com/patrickmn/go-cache"
)
var Cache *cache.Cache

func init() {
	conf := config.GetConf()
	Cache = cache.New(conf.CacheExpireTime,conf.CacheFlashTime)
}



