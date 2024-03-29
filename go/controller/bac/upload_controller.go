package bac

import (
	"blog/config"
	"blog/pkg/errcode"
	"blog/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
	"strconv"
	"time"
)

func UploadArticleImg(ctx *gin.Context) {
	upload_key := "image"
	file , err := ctx.FormFile(upload_key)
	if err != nil {
		ctx.JSON(http.StatusOK,errcode.UploadFileError.GetH())
		return
	}
	conf := config.GetConf()
	savename := getUniqueName() + path.Ext(file.Filename)
	savepath := conf.UploadFilePath + savename
	getUrl := conf.UploadDomain + savename
	//savepath := "/static/upload/" + savename
	fmt.Println(savepath,savename)
	if ctx.SaveUploadedFile(file,savepath) != nil {
		ctx.JSON(http.StatusOK,errcode.UploadFileSaveError.GetH())
		return
	}

	res := map[string]interface {}{"url": getUrl,"filename":savename}
	ctx.JSON(http.StatusOK,errcode.Ok.SetData(res))
}

func UploadAvatarImg(ctx *gin.Context) {
	upload_key := "image"
	file , err := ctx.FormFile(upload_key)
	if err != nil {
		ctx.JSON(http.StatusOK,errcode.UploadFileError.GetH())
		return
	}
	conf := config.GetConf()
	savename := "avatar" + path.Ext(file.Filename)
	savepath := conf.UploadFilePath + savename
	getUrl := conf.UploadDomain + savename

	if ctx.SaveUploadedFile(file, savepath) != nil {
		ctx.JSON(http.StatusOK,errcode.UploadFileSaveError.GetH())
		return
	}

	res := map[string]interface {}{"url": getUrl,"filename":savename}
	ctx.JSON(http.StatusOK,errcode.Ok.SetData(res))
}

func UploadFriendImg(ctx *gin.Context) {
	upload_key := "image"
	file , err := ctx.FormFile(upload_key)
	if err != nil {
		ctx.JSON(http.StatusOK,errcode.UploadFileError.GetH())
		return
	}
	conf := config.GetConf()
	savename := getUniqueName() + path.Ext(file.Filename)
	savepath := conf.UploadFilePath + savename
	getUrl := conf.UploadDomain + savename

	if ctx.SaveUploadedFile(file, savepath) != nil {
		ctx.JSON(http.StatusOK,errcode.UploadFileSaveError.GetH())
		return
	}

	res := map[string]interface {}{"url": getUrl,"filename":savename}
	ctx.JSON(http.StatusOK,errcode.Ok.SetData(res))
}



func getUniqueName() string {
	timestr := strconv.Itoa(int(time.Now().Unix()))
	shuffled := utils.GetShuffledStr(10)
	encrystr := timestr + shuffled
	name := utils.Md5(encrystr)
	return name
}
