package utils

import (
	"blog/config"
	"crypto/md5"
	"encoding/hex"
	"time"
)

func Md5(s string) string{
	h := md5.New()
	h.Write([]byte(s))
	encryptStr := h.Sum(nil)
	return hex.EncodeToString(encryptStr)
}

func PassEncry(pass string) string {
	conf:=config.GetConf()
	return Md5(pass + Md5(conf.Salt))
}

func YmdTotimestamp(date string) int64 {
	layout := "2006-01-02 15:04"
	loc,_ := time.LoadLocation("Local")
	the_time,err := time.ParseInLocation(layout,date,loc)
	if err != nil {
		return 0
	} else {
		return the_time.Unix()
	}
}
