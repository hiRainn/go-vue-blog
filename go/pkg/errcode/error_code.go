package errcode

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type ERRCODE struct {
	code int  //错误码
	msg string  //错误信息
	err error   //debug错误信息
}

func (err *ERRCODE) getErr() {
	fmt.Printf("Error - %d, message : %s ,error : %s",err.code, err.msg, err.err)
}

//返回gin.H格式
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
01-article
02-auth
03-init
----第四五位----
01-密码错误
02-上传文件过大
03-获取菜单失败
 */

//常用错误码举例
var (
	Ok = &ERRCODE{code: 0, msg: "ok"}
	// ------------  系统错误  -----------
	ErrParams = &ERRCODE{code: 10001, msg: "参数错误"}
	FileOverMax = &ERRCODE{code: 10101, msg: "上传文件过大"}

	//------------- 一般功能错误  ------------
	GetMenuError = &ERRCODE{code: 20303, msg: "获取菜单失败"}
)



/*
更为严谨的错误码设计原则上会有更长的划分，其中会包含具体第几位代表哪个业务哪个功能的错误代码
考虑到实际debug的时候，测试人员或者程序员测试的时候，知道错误发生的具体请求，所以5位已经足够大部分的使用场景
 */
