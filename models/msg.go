package models

import (
	"encoding/json"
	"fmt"
	"pet/utils"
	"reflect"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/beego/redigo/redis"
	"github.com/davecgh/go-spew/spew"
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
	PhotoId   int64
}
type MsgPhotoApi struct {
	PhotoPath string
	Content   string
	CreatedAt string
	HeadImage string
	UserId    int64
	Photo     *PhotosApi
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
			msgPhotoApi.CreatedAt = helper.GetTimeAgo(int64(photoMsgMap["CreatedAt"].(float64)))
			msgPhotoApi.UserId = user.Id

			if reflect.ValueOf(photoMsgMap["Id"]).Kind().String() == "int64" {
				photo, _ := GetPhotosById(photoMsgMap["Id"].(int64))
				msgPhotoApi.Photo = ConverToPhotoApiStruct(photo)
			}
			msgList = append(msgList, msgPhotoApi)
		}
		if msg.Kind == 1 {
			photoMsgMap := msg.Object.(map[string]interface{})
			spew.Dump(photoMsgMap)
			msgPhotoApi := new(MsgPhotoApi)
			msgPhotoApi.PhotoPath = photoMsgMap["PhotoPath"].(string)
			if reflect.ValueOf(photoMsgMap["Id"]).Kind().String() == "int64" {
				photo, _ := GetPhotosById(photoMsgMap["Id"].(int64))
				msgPhotoApi.Photo = ConverToPhotoApiStruct(photo)
			}

			user, _ := GetUsersById(int64(photoMsgMap["UserId"].(float64)))
			msgPhotoApi.Content = user.Name + "评论了你的照片:" + photoMsgMap["Content"].(string)
			msgPhotoApi.HeadImage = user.Head
			msgPhotoApi.UserId = user.Id
			msgPhotoApi.CreatedAt = helper.GetTimeAgo(int64(photoMsgMap["CreatedAt"].(float64)))
			msgList = append(msgList, msgPhotoApi)
		}
	}
	return msgList
}

func GetFollowingMsgPhotos(userId int64, offset int64, limit int64) ([]*MsgPhotoApi, error) {
	var photoApiDatas []*MsgPhotoApi
	var photo Photos
	redisAddress, _ := beego.Config("String", "redisServer", "")
	c, err := redis.Dial("tcp", redisAddress.(string))
	defer c.Close()
	if err != nil {
		beego.Error(err.Error())
	}
	userIdStr := strconv.FormatInt(userId, 10)
	result, err := c.Do("LRANGE", "ptm:"+userIdStr, offset, offset+limit)

	if err != nil {
		beego.Error(err.Error())
	}

	if reflect.TypeOf(result).String() == "[]interface {}" {
		if reflect.ValueOf(result).Len() == 0 {
			return photoApiDatas, nil

		}
	}

	var photoIdList []string
	for _, photoId := range result.([]interface{}) {

		photoIdList = append(photoIdList, string(photoId.([]uint8)))

	}
	o := orm.NewOrm()
	qs := o.QueryTable("photos")
	var lists []orm.Params
	qs.Filter("id__in", photoIdList).Values(&lists)

	msgPhoto := new(MsgPhotoApi)

	for _, v := range lists {
		photo.CreatedAt = v["CreatedAt"].(time.Time)
		photo.Id = v["Id"].(int64)
		photo.Likes = v["Likes"].(int64)
		photo.Path = v["Path"].(string)
		photo.Title = v["Title"].(string)
		photo.User, _ = GetUsersById(v["User"].(int64))

		photoApiData := ConverToPhotoApiStruct(&photo)

		msgPhoto.PhotoPath = photoApiData.Path
		msgPhoto.CreatedAt = photoApiData.CreatedAt
		msgPhoto.HeadImage = photo.User.Head
		msgPhoto.Content = fmt.Sprintf("%s上传了一张照片", photo.User.Name)
		msgPhoto.UserId = photo.User.Id

		if reflect.ValueOf(v["Id"]).Kind().String() == "int64" {
			photo, _ := GetPhotosById(v["Id"].(int64))
			msgPhoto.Photo = ConverToPhotoApiStruct(photo)
		}

		photoApiDatas = append(photoApiDatas, msgPhoto)

	}
	return photoApiDatas, err
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
