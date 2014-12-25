package models

import (
	"encoding/json"
	"fmt"
	"pet/utils"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/beego/redigo/redis"
	"github.com/davecgh/go-spew/spew"
)

type Msg struct {
	Id     int64
	Kind   int
	Object interface{}
}

type MsgInCache struct {
	UserId    int64
	TargetId  int64
	PhotoId   int64
	CreatedAt int64
}
type MsgMeApi struct {
	PhotoPath string
	Content   string
	CreatedAt string
	HeadImage string
	UserId    int64
	Photo     *PhotosApi
}

type MsgTimelineApi struct {
	Photo                   *PhotosApi
	FollowingUser           *UsersApi
	FollowingUserTargetUser *UsersApi
	CreatedAt               string
	Content                 string
}

func GetMsgToMe(userIdStr string, offset, limit int64) []*MsgMeApi {

	redisAddress, _ := beego.Config("String", "redisServer", "")
	c, err := redis.Dial("tcp", redisAddress.(string))
	defer c.Close()
	if err != nil {
		beego.Error(err.Error())
	}

	msgListInterface, err := c.Do("ZREVRANGE", "msg:"+userIdStr, offset, offset+limit)

	var msgList []*MsgMeApi

	for _, v := range msgListInterface.([]interface{}) {
		var msg Msg
		err := json.Unmarshal(v.([]uint8), &msg)
		if err != nil {
			fmt.Println("error:", err)
		}

		if msg.Kind == 0 {
			msgMap := msg.Object.(map[string]interface{})
			msgMeApi := new(MsgMeApi)
			spew.Dump(msgMap)

			photoId := int64(msgMap["PhotoId"].(float64))
			photo, _ := GetPhotosById(photoId)
			photoApi := ConverToPhotoApiStruct(photo)

			msgMeApi.PhotoPath = photo.Path

			user, _ := GetUsersById(int64(msgMap["UserId"].(float64)))

			msgMeApi.Content = user.Name + "喜欢了你的照片"
			msgMeApi.HeadImage = user.Head
			msgMeApi.UserId = user.Id
			msgMeApi.CreatedAt = helper.GetTimeAgo(int64(msgMap["CreatedAt"].(float64)))
			msgMeApi.Photo = photoApi
			msgList = append(msgList, msgMeApi)
		}
		if msg.Kind == 1 {
			msgMap := msg.Object.(map[string]interface{})
			msgMeApi := new(MsgMeApi)

			photoId := int64(msgMap["PhotoId"].(float64))
			photo, _ := GetPhotosById(photoId)
			photoApi := ConverToPhotoApiStruct(photo)

			msgMeApi.PhotoPath = photo.Path

			user, _ := GetUsersById(int64(msgMap["UserId"].(float64)))

			commentId := int64(msgMap["TargetId"].(float64))
			comment, _ := GetPhotoCommentsById(commentId)
			msgMeApi.Content = user.Name + "评论了你的照片:" + comment.Content
			msgMeApi.HeadImage = user.Head
			msgMeApi.UserId = user.Id
			msgMeApi.CreatedAt = helper.GetTimeAgo(int64(msgMap["CreatedAt"].(float64)))
			msgMeApi.Photo = photoApi
			msgList = append(msgList, msgMeApi)
		}
		if msg.Kind == 3 {
			msgMap := msg.Object.(map[string]interface{})
			msgMeApi := new(MsgMeApi)

			user, _ := GetUsersById(int64(msgMap["UserId"].(float64)))

			msgMeApi.Content = user.Name + "喜欢了你关注了"
			msgMeApi.HeadImage = user.Head
			msgMeApi.UserId = user.Id
			msgMeApi.CreatedAt = helper.GetTimeAgo(int64(msgMap["CreatedAt"].(float64)))
			msgList = append(msgList, msgMeApi)
		}
	}
	return msgList
}

func GetFollowingTimeline(currentUserId int64, offset int64, limit int64) ([]*MsgTimelineApi, error) {
	redisAddress, _ := beego.Config("String", "redisServer", "")
	c, err := redis.Dial("tcp", redisAddress.(string))
	defer c.Close()
	if err != nil {
		beego.Error(err.Error())
	}
	currentUserIdStr := strconv.FormatInt(currentUserId, 10)
	msgListInterface, err := c.Do("ZREVRANGE", "ftm:"+currentUserIdStr, offset, offset+limit)

	var msgList []*MsgTimelineApi

	spew.Dump(msgListInterface)
	for _, v := range msgListInterface.([]interface{}) {
		var msg Msg
		err := json.Unmarshal(v.([]uint8), &msg)
		if err != nil {
			fmt.Println("error:", err)
		}

		if msg.Kind == 0 {
			msgMap := msg.Object.(map[string]interface{})
			msgApi := new(MsgTimelineApi)

			photoId := int64(msgMap["PhotoId"].(float64))
			photo, _ := GetPhotosById(photoId)
			msgApi.Photo = ConverToPhotoApiStruct(photo)

			msgApi.CreatedAt = helper.GetTimeAgo(int64(msgMap["CreatedAt"].(float64)))

			sourceUserId := int64(msgMap["UserId"].(float64))
			sourceUser, _ := GetUsersById(sourceUserId)
			msgApi.FollowingUser = ConverToUserApiStruct(sourceUser)

			msgApi.Content = fmt.Sprintf("%s 喜欢了一张照片", sourceUser.Name)
			msgList = append(msgList, msgApi)
		}
		if msg.Kind == 3 {
			msgMap := msg.Object.(map[string]interface{})
			msgApi := new(MsgTimelineApi)

			userId := int64(msgMap["UserId"].(float64))
			targetUserId := int64(msgMap["TargetUserId"].(float64))

			sourceUser, _ := GetUsersById(userId)
			targetUser, _ := GetUsersById(targetUserId)

			msgApi.FollowingUser = ConverToUserApiStruct(sourceUser)
			msgApi.FollowingUserTargetUser = ConverToUserApiStruct(targetUser)

			msgApi.CreatedAt = helper.GetTimeAgo(int64(msgMap["CreatedAt"].(float64)))
			msgApi.Content = fmt.Sprintf("%s 关注了用户 %s", sourceUser.Name, targetUser.Name)
			msgList = append(msgList, msgApi)
		}

		if msg.Kind == 2 {
			msgMap := msg.Object.(map[string]interface{})
			msgApi := new(MsgTimelineApi)

			photoId := int64(msgMap["PhotoId"].(float64))
			photo, _ := GetPhotosById(photoId)

			msgApi.Photo = ConverToPhotoApiStruct(photo)
			userId := int64(msgMap["UserId"].(float64))
			sourceUser, _ := GetUsersById(userId)
			msgApi.FollowingUser = ConverToUserApiStruct(sourceUser)
			msgApi.CreatedAt = helper.GetTimeAgo(int64(msgMap["CreatedAt"].(float64)))
			msgApi.Content = fmt.Sprintf("%s 上传了一张照片", sourceUser.Name)
			msgList = append(msgList, msgApi)
		}
	}
	return msgList, nil
}

func NoticeToFriendsTimeline(currentUserId, targetUserId, photoId int64, kind int, content ...string) (err error) {

	redisAddress, _ := beego.Config("String", "redisServer", "")
	c, err := redis.Dial("tcp", redisAddress.(string))
	defer c.Close()
	if err != nil {
		beego.Error(err.Error())
	}

	var msgInCache *MsgInCache

	msgInCache = &MsgInCache{
		UserId:    currentUserId,
		TargetId:  targetUserId,
		PhotoId:   photoId,
		CreatedAt: time.Now().Unix(),
	}

	msg := new(Msg)
	msg.Kind = kind
	msg.Object = msgInCache
	b, _ := json.Marshal(msg)

	//get User's followers
	currentUserIdStr := strconv.FormatInt(currentUserId, 10)
	result, err := c.Do("ZRANGE", "follower:"+currentUserIdStr, 0, -1)
	for _, userId := range result.([]interface{}) {
		followerUserIdStr := string(userId.([]uint8))
		_, err = c.Do("ZADD", "ftm:"+followerUserIdStr, time.Now().Unix(), string(b))
		if err != nil {
			beego.Error(err.Error())
		}
	}

	//friends timeline
	if err != nil {
		beego.Error(err.Error())
	}
	return err
}

func Notice(source, target int64, kind int, ext ...int64) (err error) {
	redisAddress, _ := beego.Config("String", "redisServer", "")
	c, err := redis.Dial("tcp", redisAddress.(string))
	defer c.Close()
	if err != nil {
		beego.Error(err.Error())
	}

	switch kind {
	case 0:
		{
			photo, _ := GetPhotosById(target)
			msgPhoto := &MsgInCache{
				UserId:    source,
				PhotoId:   target,
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
			msgPhoto := &MsgInCache{
				UserId:    source,
				TargetId:  ext[0],
				PhotoId:   target,
				CreatedAt: time.Now().Unix(),
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
	case 3:
		{
			msgPhoto := &MsgInCache{
				UserId:    source,
				TargetId:  target,
				CreatedAt: time.Now().Unix(),
			}
			msg := new(Msg)
			msg.Kind = 3
			msg.Object = msgPhoto
			b, _ := json.Marshal(msg)

			sourceUserIdStr := strconv.FormatInt(target, 10)

			_, err := c.Do("ZADD", "msg:"+sourceUserIdStr, time.Now().Unix(), string(b))
			if err != nil {
				beego.Error(err.Error())
			}
		}
	}

	return err
}
