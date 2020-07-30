package bac

import (
	"blog/pkg/errcode"
	"blog/utils"
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
	savename := getUniqueName() + path.Ext(file.Filename)
	savepath := "/static/upload/"+ savename

	if ctx.SaveUploadedFile(file,"." + savepath) != nil {
		ctx.JSON(http.StatusOK,errcode.UploadFileSaveError.GetH())
		return
	}

	host := "http://" + ctx.Request.Host
	res := map[string]interface {}{"url": host + savepath,"filename":savename}
	ctx.JSON(http.StatusOK,errcode.Ok.SetData(res))
}

func getUniqueName() string {
	timestr := strconv.Itoa(int(time.Now().Unix()))
	shuffled := utils.GetShuffledStr(10)
	encrystr := timestr + shuffled
	name := utils.Md5(encrystr)
	return name
}
