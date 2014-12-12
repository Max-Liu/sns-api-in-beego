package models

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/beego/redigo/redis"
)

type Msg struct {
	Id     int64
	Kind   int
	Object interface{}
}

type MsgPhoto struct {
	PhotoPath string
	Content   string
	UserId    int64
	CreatedAt int64
}
type MsgPhotoApi struct {
	PhotoPath string
	Content   string
	CreatedAt int64
	HeadImage string
}

func GetMsgPhotoApiData(userIdStr string, offset, limit int64) []*MsgPhotoApi {

	redisAddress, _ := beego.Config("String", "redisServer", "")
	c, err := redis.Dial("tcp", redisAddress.(string))
	defer c.Close()
	if err != nil {
		beego.Error(err.Error())
	}

	msgListInterface, err := c.Do("ZREVRANGE", "msg:"+userIdStr, offset, offset+limit)

	var msgList []*MsgPhotoApi

	for _, v := range msgListInterface.([]interface{}) {
		var msg Msg
		err := json.Unmarshal(v.([]uint8), &msg)
		if err != nil {
			fmt.Println("error:", err)
		}

		if msg.Kind == 0 {
			photoMsgMap := msg.Object.(map[string]interface{})
			msgPhotoApi := new(MsgPhotoApi)
			msgPhotoApi.PhotoPath = photoMsgMap["PhotoPath"].(string)
			user, _ := GetUsersById(int64(photoMsgMap["UserId"].(float64)))
			msgPhotoApi.Content = user.Name + "喜欢了你的照片"
			msgPhotoApi.HeadImage = user.Head
			msgPhotoApi.CreatedAt = int64(photoMsgMap["CreatedAt"].(float64))
			msgList = append(msgList, msgPhotoApi)
		}
		if msg.Kind == 1 {
			photoMsgMap := msg.Object.(map[string]interface{})
			msgPhotoApi := new(MsgPhotoApi)
			msgPhotoApi.PhotoPath = photoMsgMap["PhotoPath"].(string)
			user, _ := GetUsersById(int64(photoMsgMap["UserId"].(float64)))
			msgPhotoApi.Content = user.Name + "评论了你的照片:" + photoMsgMap["Content"].(string)
			msgPhotoApi.HeadImage = user.Head
			msgPhotoApi.CreatedAt = int64(photoMsgMap["CreatedAt"].(float64))
			msgList = append(msgList, msgPhotoApi)
		}
	}
	return msgList
}

func Notice(source, target int64, kind int, content ...string) (err error) {
	redisAddress, _ := beego.Config("String", "redisServer", "")
	c, err := redis.Dial("tcp", redisAddress.(string))
	defer c.Close()
	if err != nil {
		beego.Error(err.Error())
	}
	sourceUser, _ := GetUsersById(source)

	switch kind {
	case 0:
		{
			photo, _ := GetPhotosById(target)
			msgPhoto := &MsgPhoto{
				PhotoPath: photo.Path,
				UserId:    sourceUser.Id,
				CreatedAt: time.Now().Unix(),
			}
			msg := new(Msg)
			msg.Kind = 0
			msg.Object = msgPhoto
			b, _ := json.Marshal(msg)

			sourceUserIdStr := strconv.FormatInt(photo.User.Id, 10)
			_, err := c.Do("ZADD", "msg:"+sourceUserIdStr, time.Now().Unix(), string(b))
			if err != nil {
				beego.Error(err.Error())
			}
		}
	case 1:
		{
			photo, _ := GetPhotosById(target)
			msgPhoto := &MsgPhoto{
				PhotoPath: photo.Path,
				UserId:    sourceUser.Id,
				CreatedAt: time.Now().Unix(),
				Content:   content[0],
			}
			msg := new(Msg)
			msg.Kind = 1
			msg.Object = msgPhoto
			b, _ := json.Marshal(msg)

			sourceUserIdStr := strconv.FormatInt(photo.User.Id, 10)
			_, err := c.Do("ZADD", "msg:"+sourceUserIdStr, time.Now().Unix(), string(b))
			if err != nil {
				beego.Error(err.Error())
			}
		}
	}

	return err
}
