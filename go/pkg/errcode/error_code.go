package errcode

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type ERRCODE struct {
	code int  //错误码
	msg string  //错误信息
	data interface{}
	err error   //debug错误信息
}

func (err *ERRCODE) GetErr() {
	fmt.Printf("Error - %d, message : %s ,error : %s",err.code, err.msg, err.err)
}

func (err *ERRCODE) SetOk(data interface{}) gin.H {
	return gin.H{"code": 0 ,"msg": "ok","data": data}
}

func (err *ERRCODE) SetData(data interface{}) gin.H {
	if err.code != 0 {
		err.data = data
	}
	return gin.H{"code":err.code,"msg":err.msg,"data":data}
}

// return gin.H
func (err * ERRCODE) GetH() gin.H {
	return gin.H{"code":err.code, "msg":err.msg}
}
//设计原则，code为5位数组成
//第一位代表错误级别，如1-系统错误，2-一般功能错误，3-数据库错误，
//第二三位代表模块，如01-article，02-auth，03-init
//第四五位代表错误类型，如01-密码错误，02-发布失败

//模块及常用错误对应表
/*
----第一位----
1-系统错误
2-数据库错误
3-一般功能错误
4-调用外部api错误
----第二三位----
01-init
02-article
03-auth
04-conf

----第四五位----
01-密码错误
02-上传文件过大
03-获取菜单失败
 */

//常用错误码举例
var (
	Ok = &ERRCODE{code: 0, msg: "ok"}
	// ------------  system error   -----------
	SystemError = &ERRCODE{code: 10000,msg: "system error"}
	ErrParams = &ERRCODE{code: 10001, msg: "error params"}
	FileOverMax = &ERRCODE{code: 10101, msg: "over size"}
	ParamError = &ERRCODE{code: 10301,msg: "params error"}

	// ------------- database error ------------
	DataBaseError = &ERRCODE{code : 20101, msg: "database conect error"}
	DataNotExists = &ERRCODE{code: 20102,msg: "data is not exists"}
	InsertAuthError = &ERRCODE{code: 20301,msg: "insert auth error"}
	InsertConfError = &ERRCODE{code: 20401,msg: "insert conf error"}

	//-------------admin functions error  ------------
	AuthError = &ERRCODE{code: 300001,msg: "token is not exists"}
	InitError = &ERRCODE{code:30101, msg:"init error"}
	EmptyPassError = &ERRCODE{code:30301, msg: "empty password"}
	RepeatPassError = &ERRCODE{code:30302, msg: "password repeat error"}
	BlogNameEmpty = &ERRCODE{code: 30303, msg: "blog name not null"}
	PasswordError = &ERRCODE{code: 30304,msg: "username or pasword error"}
	AddCateGoryNameError = &ERRCODE{code : 30305, msg: "catagory's name is null"}
	AddCateGoryRepeat = &ERRCODE{code: 30306,msg: "name is repaet"}
	AddCateGoryError = &ERRCODE{code : 30307, msg: "add category error"}
	AddTagNameError = &ERRCODE{code: 30308,msg: "tag's name is null"}
	AddTagError = &ERRCODE{code: 30309,msg: "add tag error"}
	AddArticleError = &ERRCODE{code: 30310,msg: "add article error"}
	EditSaveArticleError = &ERRCODE{code: 30311,msg: "edit article error"}
	ArticleTitleEmpty = &ERRCODE{code: 30311,msg: "title is empty"}
	ArticleCateEmpty = &ERRCODE{code: 30312, msg:"category is empty"}
	ArticleContentEmpty = &ERRCODE{code: 30313,msg: "content is empty"}
	SetIncTagsError = &ERRCODE{code: 30314,msg: "tag setInc error"}
	SetIncCateError = &ERRCODE{code: 30315,msg: "cate setInc error"}
	SetDecTagsError = &ERRCODE{code: 30316,msg: "tag setDec error"}
	SetDecCateError = &ERRCODE{code: 30317,msg: "cate setDec error"}
	EditArticleError = &ERRCODE{code: 30318,msg: "edit article error"}
	DeleteArticleError = &ERRCODE{code: 30319,msg: "delete article error"}

	UploadFileError = &ERRCODE{code: 30398,msg: "upload file error"}
	UploadFileSaveError = &ERRCODE{code: 30399,msg: "upload file save error"}

	//-------------app functions error  ------------
	GetCommentsError = &ERRCODE{code: 90201,msg: "get comments error"}
	CommentError =&ERRCODE{code: 90202,msg: "comment faild"}
	LikeError = &ERRCODE{code: 90203,msg: "like fail"}
	LikeRepeat = &ERRCODE{code: 90204,msg: "like repeat"}
	DislikeRepeat = &ERRCODE{code: 90205,msg: "dislike repeat"}
	ReportRepeat = &ERRCODE{code: 90206,msg: "repoet repeat"}
	UnlikeError = &ERRCODE{code: 90207,msg: "like fail"}
	ReportError = &ERRCODE{code: 90208,msg: "like fail"}
)



/*
更为严谨的错误码设计原则上会有更长的划分，其中会包含具体第几位代表哪个业务哪个功能的错误代码
考虑到实际debug的时候，测试人员或者程序员测试的时候，知道错误发生的具体请求，所以5位已经足够大部分的使用场景
 */
