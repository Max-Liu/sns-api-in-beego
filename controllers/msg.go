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
	c.Mapping("GetAll", c.GetAll)
}

// @Title 获取通知列表
// @Description 获取通知列表
// @Param	offset	query	string	false	"列表索引"
// @Success 200 {object} models.Msg
// @Failure 403
// @router / [get]
func (this *MsgController) GetAll() {
	data := make(map[string]interface{})
	if v, err := this.GetInt("offset"); err == nil {
		offset = int64(v)
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
