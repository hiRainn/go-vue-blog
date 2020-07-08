package utils

import (
	"blog/config"
	"crypto/md5"
	"encoding/hex"
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
