package models

import (
	"encoding/json"
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

func Notice(source, target int64, kind int) (err error) {
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
	}

	return err
}
