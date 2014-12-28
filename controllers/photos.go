package controllers

import (
	"os"
	"pet/models"
	helper "pet/utils"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/beego/redigo/redis"
)

var uploadPhotoPath string = "static/uploads/photos"
var err error

// 用户照片相关
type PhotosController struct {
	beego.Controller
}

func (this *PhotosController) URLMapping() {
	this.Mapping("Post", this.Post)
	this.Mapping("GetOne", this.GetOne)
	this.Mapping("GetAll", this.GetAll)
	this.Mapping("Put", this.Put)
	this.Mapping("Delete", this.Delete)
}

// @Title 发布照片
// @Description 发布照片
// @Param	photo		Form 	File	true		"用户发布的照片"
// @Param	title		form 	String	true		"照片描述"
// @Param	x		form 	String	true		"地理位置x"
// @Param	y		form 	String	true		"地理位置y"
// @Success 200 {int} models.Photos.Id
// @Failure 403 body is empty
// @router / [post]
func (this *PhotosController) Post() {
	var v models.Photos
	valid := validation.Validation{}
	this.ParseForm(&v)
	passed, _ := valid.Valid(&v)
	if !passed {
		outPut := helper.Reponse(1, nil, valid.Errors[0].Key+" "+valid.Errors[0].Message)
		this.Data["json"] = outPut
	} else {
		todayDateDir := "/" + helper.GetTodayDate()
		if _, err := os.Stat(uploadPhotoPath + todayDateDir); os.IsNotExist(err) {
			os.Mkdir(uploadPhotoPath+todayDateDir, 0777)
		}
		currentUser := this.GetSession("user").(models.Users)
		photoName := helper.GetGuid(currentUser.Id)
		dateSubdir := "/" + string(photoName[0]) + string(photoName[1])

		if _, err := os.Stat(uploadPhotoPath + todayDateDir + dateSubdir); os.IsNotExist(err) {
			os.Mkdir(uploadPhotoPath+todayDateDir+dateSubdir, 0777)
		}

		imagePath := uploadPhotoPath + todayDateDir + dateSubdir + "/" + photoName + ".jpg"

		err := this.SaveToFile("photo", imagePath)
		if err != nil {
			outPut := helper.Reponse(1, nil, err.Error())
			this.Data["json"] = outPut
		} else {
			v.Path = imagePath
			v.CreatedAt = time.Now()
			v.UpdatedAt = time.Now()
			v.Likes = 0
			v.User = &currentUser

			if id, err := models.AddPhotos(&v); err == nil {
				v.Id = id
				models.NoticeToFriendsTimeline(currentUser.Id, 0, v.Id, 2)
				data := models.ConverToPhotoApiStruct(&v)
				outPut := helper.Reponse(0, data, "创建成功")
				this.Data["json"] = outPut
			} else {
				outPut := helper.Reponse(1, nil, err.Error())
				this.Data["json"] = outPut
			}
		}
	}
	this.ServeJson()
}

// @Title 获取照片详情
// @Description 获取某张照片详情
// @Param	id		path 	string	true		"照片详情"
// @Success 200 {object} models.Photos
// @Failure 403 :id is empty
// @router /:id [get]
func (this *PhotosController) GetOne() {

	var data map[string]interface{} = make(map[string]interface{})

	idStr := this.Ctx.Input.Params[":id"]
	id, _ := strconv.ParseInt(idStr, 10, 0)
	photo, err := models.GetPhotosById(id)
	if err != nil {
		outPut := helper.Reponse(1, nil, err.Error())
		this.Data["json"] = outPut
		this.ServeJson()
		return
	}

	data["photo"] = models.ConverToPhotoApiStruct(photo)

	//userSession := this.GetSession("user").(models.Users)
	//userId := strconv.Itoa(userSession.Id)

	//query["user_id"] = userId
	//fields = []string{"Content", "CreatedAt", "User__name"}
	//data["comments"], err = models.GetAllPhotoComments(query, fields, sortby, order, offset, limit)

	outPut := helper.Reponse(0, data, "")
	this.Data["json"] = outPut
	this.ServeJson()
}

// @Title 获取照片
// @Description 获取照片列表
// @Param	sortby	query	string	false	"获取最新sortby=created_at;获取最热sortby=likes。默认获取最新"
// @Param	offset	query	string	false	"结果索引"
// @Param	myphoto	query	string	false	"为1是获取我的照片列表,默认获取全部"
// @Param	user_id	query	string	false	"获取某一个用户的照片列表,默认获取全部"
// @Success 200 {object} models.Photos
// @Failure 403
// @router / [get]
func (this *PhotosController) GetAll() {

	data := make(map[string]interface{})
	query := make(map[string]string)

	if v := this.GetString("myphoto"); err == nil {
		if v == "1" {
			userIdInt := this.GetSession("user").(models.Users).Id
			userIdStr := strconv.FormatInt(userIdInt, 10)
			query["user_id"] = userIdStr
		}
	}

	if v := this.GetString("user_id"); err == nil {
		if v != "" {
			query["user_id"] = v
		}
	}

	if v, err := this.GetInt("offset"); err == nil {
		offset = int64(v)
	}

	if v := this.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	fields := []string{"Title", "Path", "Likes", "CreatedAt", "Id", "User"}

	photos, err := models.GetAllPhotos(query, fields, sortby, order, offset, limit)
	oneMore, _ := models.GetAllPhotos(query, fields, sortby, order, offset+limit, 1)
	if len(oneMore) == 0 {
		data["Has_more"] = 0

	} else {
		data["Has_more"] = 1
	}

	var photoApiDatas []*models.PhotosApi
	var photo models.Photos

	userServerSession := this.GetSession("user")

	var currentUserSession models.Users
	if userServerSession == nil {
		currentUserSession = *new(models.Users)
	} else {
		currentUserSession = userServerSession.(models.Users)
	}

	for _, v := range photos {
		photo.Title = v["Title"].(string)
		photo.Id = v["Id"].(int64)
		photo.Likes = v["Likes"].(int64)
		photo.CreatedAt = v["CreatedAt"].(time.Time)
		photo.Path = v["Path"].(string)
		photo.User, _ = models.GetUsersById(v["User__User"].(int64))
		photoApiData := models.ConverToPhotoApiStruct(&photo, &currentUserSession)
		photoApiDatas = append(photoApiDatas, photoApiData)
	}
	if len(photoApiDatas) == 0 {
		data["Photos"] = ""

	} else {
		data["Photos"] = photoApiDatas
	}
	if err != nil {
		outPut := helper.Reponse(1, nil, err.Error())
		this.Data["json"] = outPut
		this.ServeJson()
		return
	}

	outPut := helper.Reponse(0, data, "")
	this.Data["json"] = outPut
	this.ServeJson()
}

// @Title Update
// @Description update the Photos
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Photos	true		"body for Photos content"
// @Success 200 {object} models.Photos
// @Failure 403 :id is not int
// @router /:id [put]

func (this *PhotosController) Put() {
	//idStr := this.Ctx.Input.Params[":id"]
	//id, _ := strconv.Atoi(idStr)
	//v := models.Photos{Id: id}
	//json.Unmarshal(this.Ctx.Input.RequestBody, &v)
	//if err := models.UpdatePhotosById(&v); err == nil {
	//this.Data["json"] = "OK"
	//} else {
	//this.Data["json"] = err.Error()
	//}
	//this.ServeJson()
}

// @Title Delete
// @Description delete the Photos
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]

func (this *PhotosController) Delete() {
	//idStr := this.Ctx.Input.Params[":id"]
	//id, _ := strconv.Atoi(idStr)
	//if err := models.DeletePhotos(id); err == nil {
	//this.Data["json"] = "OK"
	//} else {
	//this.Data["json"] = err.Error()
	//}
	//this.ServeJson()
}

func PushPhotoToFollowerTimelime(currentUserId, photoId int64) (err error) {
	redisAddress, _ := beego.Config("String", "redisServer", "")
	c, err := redis.Dial("tcp", redisAddress.(string))
	defer c.Close()
	if err != nil {
		beego.Error(err.Error())
	}
	currentUserIdStr := strconv.FormatInt(currentUserId, 10)
	result, err := c.Do("ZRANGE", "follower:"+currentUserIdStr, 0, -1)
	for _, userId := range result.([]interface{}) {
		targetUserIdStr := string(userId.([]uint8))
		targetUserId, _ := strconv.ParseInt(targetUserIdStr, 10, 0)
		err = models.NoticeToFriendsTimeline(currentUserId, targetUserId, photoId, 2)
		if err != nil {
			beego.Error(err.Error())
		}
	}
	return err
}

// @Title 获取Top10
// @Description 获取照片Top10
// @Success
// @Failure 403
// @router /top10 [get]
func (this *PhotosController) GetTop10() {

	photos, err := models.GetTop10()

	if err != nil {
		outPut := helper.Reponse(1, nil, err.Error())
		this.Data["json"] = outPut
	} else {
		outPut := helper.Reponse(0, photos, "")
		this.Data["json"] = outPut
	}
	this.ServeJson()
}
