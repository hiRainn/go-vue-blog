package bac

import (
	"blog/model"
	"blog/pkg/errcode"
	"blog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetBlogSetting(c *gin.Context) {
	var conf model.BlogConfig
	list,_ := conf.Find()
	c.JSON(http.StatusOK,errcode.Ok.SetData(list))
}

func SetBlogSetting(c *gin.Context) {
	var params map[string] interface{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusOK,errcode.ParamError.GetH())
		return
	}
	var conf model.BlogConfig
	for key,value := range params {
		conf.Key = key
		v,_ := value.(string)
		conf.Value = v
		conf.Update()
	}
	c.JSON(http.StatusOK,errcode.Ok.SetData(conf))
}

func GetUserSetting(c *gin.Context) {
	var a model.BlogAuth
	a.GetAllInfo()
	c.JSON(http.StatusOK,errcode.Ok.SetData(a))
}

func SetUserSetting(c *gin.Context) {
	var params map[string] interface{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusOK,errcode.ParamError.GetH())
		return
	}
	the_time := params["birthday"].(string)
	if the_time != "" {
		params["birthday"] = utils.YmdTotimestamp(the_time)
	}
	var a model.BlogAuth
	a.SetInfo(params)
	c.JSON(http.StatusOK,errcode.Ok.SetData(a))
}

func GetFriendsLink(c *gin.Context) {
	var f model.BlogFriends
	link,_ := f.GetLink()
	c.JSON(http.StatusOK,errcode.Ok.SetData(link))
}

func DelFriend(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK,errcode.ParamError)
		return
	}
	f := new(model.BlogFriends)
	f.Id,_ = strconv.Atoi(id)
	if f.Id == 0 {
		c.JSON(http.StatusOK,errcode.ParamError.GetH())
		return
	}
	if err := f.DelFriend();err!= nil {
		c.JSON(http.StatusOK,errcode.ParamError.GetH())
	} else {
		c.JSON(http.StatusOK,errcode.Ok.GetH())
	}
}

func AddFriends(c *gin.Context) {
	var params map[string] interface{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusOK,errcode.ParamError.GetH())
		return
	}
	f := new(model.BlogFriends)
	f.Id,_ = params["id"].(int)
	f.Name,_ = params["name"].(string)
	f.Link,_ = params["link"].(string)
	f.Avatar,_ = params["avatar"].(string)
	sort1,_ := params["sort"].(string)
	sort2,_ := strconv.Atoi(sort1)
	f.Sort = uint8(sort2)
	var err error
	if f.Id == 0 {
		_,err = f.AddFriend()
	} else {
		err = f.EditFriend()
	}
	if err != nil {
		c.JSON(http.StatusOK,errcode.ErrParams.GetH())
	} else {
		c.JSON(http.StatusOK,errcode.Ok.SetData(f.Id))
	}
}


