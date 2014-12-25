package controllers

import (
	"pet/models"
	"pet/utils"
	"strconv"

	"github.com/astaxie/beego"
)

//获取通知接口
type MsgController struct {
	beego.Controller
}

func (c *MsgController) URLMapping() {
	//c.Mapping("GetAll", c.GetAll)
}

// @Title 获取我的动态列表
// @Description 获取我的动态列表
// @Param	offset	query	string	false	"列表索引"
// @Success 200 {object} models.Msg
// @Failure 403
// @router /me [get]
func (this *MsgController) Me() {
	data := make(map[string]interface{})
	if v, err := this.GetInt("offset"); err == nil {
		offset = int64(v)
	} else {
		offset = 0
	}

	userSession := this.GetSession("user").(models.Users)
	userIdStr := strconv.FormatInt(userSession.Id, 10)
	msgList := models.GetMsgPhotoApiData(userIdStr, offset, limit)

	if len(msgList) == 0 {
		data["Message"] = ""
	} else {

		data["Message"] = msgList
	}
	oneMore := models.GetMsgPhotoApiData(userIdStr, offset+limit, 1)
	if len(oneMore) == 0 {
		data["Has_more"] = 0
	} else {
		data["Has_more"] = 1
	}

	outPut := helper.Reponse(0, data, "")
	this.Data["json"] = outPut
	this.ServeJson()
}

// @Title 获取关注的人动态
// @Description 获取关注的人动态
// @Param	offset	query	string	false	"结果列表索引"
// @Success 200 {object} models.Timeline
// @Failure 403
// @router /following [get]
func (this *MsgController) GetFollowingPhotosTimeline() {

	data := make(map[string]interface{})

	if v, err := this.GetInt("offset"); err == nil {
		offset = int64(v)
	} else {
		offset = 0
	}

	userSession := this.GetSession("user").(models.Users)
	userId := userSession.Id

	l, err := models.GetFollowingTimeline(userId, offset, limit)
	oneMore, _ := models.GetFollowingTimeline(userId, offset+limit, 1)
	if len(oneMore) == 0 {
		data["Has_more"] = 0

	} else {
		data["Has_more"] = 1
	}
	if len(l) == 0 {
		data["Timeline"] = ""
	} else {
		data["Timeline"] = l
	}

	if err != nil {
		outPut := helper.Reponse(1, nil, err.Error())
		this.Data["json"] = outPut
	} else {
		outPut := helper.Reponse(0, data, "")
		this.Data["json"] = outPut
	}
	this.ServeJson()
}

// @Title 是否有新通知
// @Description 是否有新通知
// @Success 200 {object} models.Msg
// @Failure 403
// @router /has_notice [get]
func (this *MsgController) HasNotice() {
	resp := make(map[string]interface{})
	resp["HasNotice"] = true
	outPut := helper.Reponse(0, resp, "")
	this.Data["json"] = outPut
	this.ServeJson()
}
