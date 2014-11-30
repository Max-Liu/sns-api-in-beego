package controllers

import (
	"encoding/json"
	"fmt"
	"pet/models"
	"pet/utils"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/beego/redigo/redis"
)

// oprations for Msg
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
	if v, err := this.GetInt("offset"); err == nil {
		offset = int64(v)
	}

	redisAddress, _ := beego.Config("String", "redisServer", "")
	c, err := redis.Dial("tcp", redisAddress.(string))
	defer c.Close()
	if err != nil {
		beego.Error(err.Error())
	}
	userSession := this.GetSession("user").(models.Users)
	userIdStr := strconv.FormatInt(userSession.Id, 10)
	msgListInterface, err := c.Do("ZREVRANGE", "msg:"+userIdStr, offset, offset+limit)

	var msgList []*models.MsgPhoto

	for _, v := range msgListInterface.([]interface{}) {
		//if v.(models.Msg).Kind == 0 {
		var msg models.Msg
		err := json.Unmarshal(v.([]uint8), &msg)
		if err != nil {
			fmt.Println("error:", err)
		}
		if msg.Kind == 0 {
			photoMsgMap := msg.Object.(map[string]interface{})
			msgPhoto := new(models.MsgPhoto)
			msgPhoto.PhotoPath = photoMsgMap["PhotoPath"].(string)
			user, _ := models.GetUsersById(int64(photoMsgMap["UserId"].(float64)))
			msgPhoto.Content = user.Name + "喜欢了你的照片"
			msgPhoto.UserId = int64(photoMsgMap["UserId"].(float64))
			msgPhoto.CreatedAt = int64(photoMsgMap["CreatedAt"].(float64))
			msgList = append(msgList, msgPhoto)
		}
	}
	outPut := helper.Reponse(0, msgList, "")
	this.Data["json"] = outPut
	this.ServeJson()
}
