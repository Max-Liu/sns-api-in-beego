package controllers

import (
	"pet/models"
	"pet/utils"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

// 喜欢照片相关操作
type LikesController struct {
	beego.Controller
}

func (this *LikesController) URLMapping() {
	this.Mapping("Post", this.Post)
	this.Mapping("GetOne", this.GetOne)
	this.Mapping("GetAll", this.GetAll)
	this.Mapping("Put", this.Put)
	this.Mapping("Delete", this.Delete)
}

// @Title 喜欢
// @Description 喜欢照片
// @Param	photo_id	form 	string	true		"照片ID"
// @Success 200 {int} models.Likes.Id
// @Failure 403 body is empty
// @router / [post]
func (this *LikesController) Post() {

	var v models.Likes
	valid := validation.Validation{}
	this.ParseForm(&v)

	photoIdStr := this.GetString("photo_id")
	photoId, _ := strconv.Atoi(photoIdStr)

	passed, _ := valid.Valid(&v)
	if !passed {
		outPut := helper.Reponse(1, nil, valid.Errors[0].Key+" "+valid.Errors[0].Message)
		this.Data["json"] = outPut
	} else {
		Photo, err := models.GetPhotosById(int64(photoId))
		if err != nil {
			outPut := helper.Reponse(1, nil, err.Error())
			this.Data["json"] = outPut
			this.ServeJson()
			return
		}

		like := new(models.Likes)
		like.CreatedAt = time.Now()
		like.UpdatedAt = time.Now()
		userSession := this.GetSession("user").(models.Users)
		like.User = &userSession
		like.Photo = Photo

		if _, err := models.AddLikes(like); err == nil {
			outPut := helper.Reponse(0, nil, "喜欢成功")
			this.Data["json"] = outPut
			models.Notice(userSession.Id, Photo.Id, 0)

		} else {

			outPut := helper.Reponse(1, nil, err.Error())
			this.Data["json"] = outPut
		}
	}
	this.ServeJson()
}

// @Title Get
// @Description get Likes by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Likes
// @Failure 403 :id is empty
// @router /:id [get]

func (this *LikesController) GetOne() {
	//idStr := this.Ctx.Input.Params[":id"]
	//id, _ := strconv.Atoi(idStr)
	//v, err := models.GetLikesById(id)
	//if err != nil {
	//this.Data["json"] = err.Error()
	//} else {
	//this.Data["json"] = v
	//}
	//this.ServeJson()
}

// @Title 我喜欢过的
// @Description 获取我喜欢过的照片列表
// @Param	user_id	query	string	false	"目标用户Id,为空代表获取当前用户"
// @Param	offset	query	string	false	"查询结果索引"
// @Success 200 {object} models.Likes
// @Failure 403
// @router / [get]
func (this *LikesController) GetAll() {

	query := make(map[string]string)
	if v, err := this.GetInt("offset"); err == nil {
		offset = int64(v)
	}

	fields := []string{"CreatedAt", "Photo"}

	if userId := this.GetString("user_id"); userId != "" {
		query["user_id"] = userId
	} else {
		userSession := this.GetSession("user").(models.Users)
		userId := strconv.FormatInt(userSession.Id, 10)
		query["user_id"] = userId
	}

	l, err := models.GetAllLikes(query, fields, sortby, order, offset, limit)

	var photoLikesDatas []*models.LikesApi
	var likes models.Likes

	for _, v := range l {
		likes.CreatedAt = v["CreatedAt"].(time.Time)
		likes.Photo, _ = models.GetPhotosById(v["Photo__Photo"].(int64))
		photoLikesData := models.ConverToLikedPhotoApiStruct(&likes)
		photoLikesDatas = append(photoLikesDatas, photoLikesData)
	}

	hasMore := hasMore(query, fields, offset, "likes")

	if err != nil {
		outPut := helper.Reponse(1, nil, err.Error())
		this.Data["json"] = outPut
	} else {
		data := make(map[string]interface{})
		data["has_more"] = hasMore
		if len(photoLikesDatas) == 0 {
			data["likes"] = ""
		} else {
			data["likes"] = photoLikesDatas
		}
		outPut := helper.Reponse(0, data, "")
		this.Data["json"] = outPut
	}
	this.ServeJson()
}

// @Title Update
// @Description update the Likes
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Likes	true		"body for Likes content"
// @Success 200 {object} models.Likes
// @Failure 403 :id is not int
// @router /:id [put]

func (this *LikesController) Put() {
	//idStr := this.Ctx.Input.Params[":id"]
	//id, _ := strconv.Atoi(idStr)
	//v := models.Likes{Id: id}
	//json.Unmarshal(this.Ctx.Input.RequestBody, &v)
	//if err := models.UpdateLikesById(&v); err == nil {
	//this.Data["json"] = "OK"
	//} else {
	//this.Data["json"] = err.Error()
	//}
	//this.ServeJson()
}

// @Title 取消喜欢
// @Description 取消喜欢
// @Param	id	path 	string	true		"照片ID"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (this *LikesController) Delete() {
	var v models.Likes
	idStr := this.Ctx.Input.Params[":id"]
	id, _ := strconv.ParseInt(idStr, 10, 0)

	valid := validation.Validation{}
	this.ParseForm(&v)
	passed, _ := valid.Valid(&v)
	if !passed {
		outPut := helper.Reponse(1, nil, valid.Errors[0].Key+" "+valid.Errors[0].Message)
		this.Data["json"] = outPut
	} else {
		userSession := this.GetSession("user").(models.Users)
		if _, err := models.DeleteLikedPhoto(userSession.Id, id); err == nil {
			outPut := helper.Reponse(0, nil, "取消喜欢成功")
			this.Data["json"] = outPut
		} else {
			outPut := helper.Reponse(1, nil, err.Error())
			this.Data["json"] = outPut
		}
	}
	this.ServeJson()
}

// @Title 获取喜欢某个照片的用户列表
// @Description 获取喜欢某个照片的用户列表
// @Param	photo_id	query 	string	true		"照片ID"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /users [get]
func (this *LikesController) UsersList() {
	photoId, _ := this.GetInt64("photo_id")
	userList, err := models.GetUsersByLikesPhoto(photoId)

	if err != nil {
		outPut := helper.Reponse(1, "", err.Error())
		this.Data["json"] = outPut
	} else {
		outPut := helper.Reponse(0, userList, "")
		this.Data["json"] = outPut
	}
	this.ServeJson()
}
