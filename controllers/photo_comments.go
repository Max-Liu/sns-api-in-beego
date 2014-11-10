package controllers

import (
	"pet/models"
	"pet/utils"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

// 照片评论相关
type PhotoCommentsController struct {
	beego.Controller
}

func (this *PhotoCommentsController) URLMapping() {
	this.Mapping("Post", this.Post)
	this.Mapping("GetOne", this.GetOne)
	this.Mapping("GetAll", this.GetAll)
	this.Mapping("Put", this.Put)
	this.Mapping("Delete", this.Delete)
}

// @Title 评论
// @Description 评论照片
// @Param	photo_id	form	string 	true		"评论照片ID"
// @Param	content		form	string 	true		"评论内容"
// @Success 200 {int} models.PhotoComments.Id
// @Failure 403 body is empty
// @router / [post]
func (this *PhotoCommentsController) Post() {
	var v models.PhotoComments
	var err error
	valid := validation.Validation{}
	this.ParseForm(&v)

	photoIdStr := this.GetString("photo_id")
	photoId, _ := strconv.Atoi(photoIdStr)

	passed, _ := valid.Valid(&v)
	if !passed {
		outPut := helper.Reponse(1, nil, valid.Errors[0].Key+" "+valid.Errors[0].Message)
		this.Data["json"] = outPut
	} else {
		v.Photo, err = models.GetPhotosById(photoId)
		if err != nil {
			outPut := helper.Reponse(1, nil, err.Error())
			this.Data["json"] = outPut
			this.ServeJson()
			return
		}
		v.CreatedAt = time.Now()
		v.UpdatedAt = time.Now()
		userSession := this.GetSession("user").(models.Users)
		v.User = &userSession

		if _, err := models.AddPhotoComments(&v); err == nil {
			outPut := helper.Reponse(0, nil, "发表成功")
			this.Data["json"] = outPut
		} else {
			outPut := helper.Reponse(1, nil, err.Error())
			this.Data["json"] = outPut
		}
	}
	this.ServeJson()
}

// @Title Get
// @Description get PhotoComments by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.PhotoComments
// @Failure 403 :id is empty
// @router /:id [get]

func (this *PhotoCommentsController) GetOne() {
	//idStr := this.Ctx.Input.Params[":id"]
	//id, _ := strconv.Atoi(idStr)
	//v, err := models.GetPhotoCommentsById(id)
	//if err != nil {
	//this.Data["json"] = err.Error()
	//} else {
	//this.Data["json"] = v
	//}
	//this.ServeJson()
}

// @Title 获取评论列表
// @Description 获取评论列表
// @Param	photo_id	query	string	ture	"照片ID"
// @Param	offset		query	string	true	"查询结果索引"
// @Success 200 {object} models.PhotoComments
// @Failure 403
// @router / [get]
func (this *PhotoCommentsController) GetAll() {

	if v, err := this.GetInt("offset"); err == nil {
		offset = int64(v)
	}
	photoId := this.GetString("photo_id")

	query := make(map[string]string)
	query["photo_id"] = photoId

	l, err := models.GetAllPhotoComments(query, fields, sortby, order, offset, limit)
	if err != nil {
		outPut := helper.Reponse(1, nil, err.Error())
		this.Data["json"] = outPut
	} else {
		outPut := helper.Reponse(0, l, "")
		this.Data["json"] = outPut
	}
	this.ServeJson()
}

// @Title Update
// @Description update the PhotoComments
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.PhotoComments	true		"body for PhotoComments content"
// @Success 200 {object} models.PhotoComments
// @Failure 403 :id is not int
// @router /:id [put]

func (this *PhotoCommentsController) Put() {
	//idStr := this.Ctx.Input.Params[":id"]
	//id, _ := strconv.Atoi(idStr)
	//v := models.PhotoComments{Id: id}
	//json.Unmarshal(this.Ctx.Input.RequestBody, &v)
	//if err := models.UpdatePhotoCommentsById(&v); err == nil {
	//this.Data["json"] = "OK"
	//} else {
	//this.Data["json"] = err.Error()
	//}
	//this.ServeJson()
}

// @Title Delete
// @Description delete the PhotoComments
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]

func (this *PhotoCommentsController) Delete() {
	//idStr := this.Ctx.Input.Params[":id"]
	//id, _ := strconv.Atoi(idStr)
	//if err := models.DeletePhotoComments(id); err == nil {
	//this.Data["json"] = "OK"
	//} else {
	//this.Data["json"] = err.Error()
	//}
	//this.ServeJson()
}
