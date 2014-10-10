package controllers

import (
	"encoding/json"
	"log"
	"os"
	"pet/models"
	"strconv"
	"strings"
	"time"
	"web"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

var uploadPhotoPath string = "static/uploads/photos"
var err error

// oprations for Photos
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

// @Title Post
// @Description create Photos
// @Param	body		body 	models.Photos	true		"body for Photos content"
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
				v.Id = int(id)
				outPut := helper.Reponse(0, v, "创建成功")
				this.Data["json"] = outPut
			} else {
				outPut := helper.Reponse(1, nil, err.Error())
				this.Data["json"] = outPut
			}
		}
	}
	this.ServeJson()
}

// @Title Get
// @Description get Photos by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Photos
// @Failure 403 :id is empty
// @router /:id [get]
func (this *PhotosController) GetOne() {

	var data map[string]interface{} = make(map[string]interface{})
	idStr := this.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	fields = []string{"user__name"}
	data["photo"], err = models.GetPhotosById(id)

	if err != nil {
		outPut := helper.Reponse(1, nil, err.Error())
		this.Data["json"] = outPut
		this.ServeJson()
	}
	userSession := this.GetSession("user").(models.Users)
	userId := strconv.Itoa(userSession.Id)

	query["user_id"] = userId
	fields = []string{"Content", "CreatedAt", "User__name"}
	log.Println(fields)
	data["comments"], err = models.GetAllPhotoComments(query, fields, sortby, order, offset, limit)

	outPut := helper.Reponse(0, data, "")
	this.Data["json"] = outPut
	this.ServeJson()
}

// @Title Get All
// @Description get Photos
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Photos
// @Failure 403
// @router / [get]
func (this *PhotosController) GetAll() {

	userIdInt := this.GetSession("user").(models.Users).Id
	userIdStr := strconv.Itoa(userIdInt)

	query := make(map[string]string)
	query["user_id"] = userIdStr

	// fields: col1,col2,entity.col3
	if v := this.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}

	if v, err := this.GetInt("offset"); err == nil {
		offset = v
	}

	fields = []string{"Title", "Path", "Likes", "CreatedAt", "Id"}

	photos, err := models.GetMyPhotos(query, fields, sortby, order, offset, limit)
	if err != nil {
		outPut := helper.Reponse(1, nil, err.Error())
		this.Data["json"] = outPut
		this.ServeJson()
		return
	}

	outPut := helper.Reponse(0, photos, "")
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
	idStr := this.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	v := models.Photos{Id: id}
	json.Unmarshal(this.Ctx.Input.RequestBody, &v)
	if err := models.UpdatePhotosById(&v); err == nil {
		this.Data["json"] = "OK"
	} else {
		this.Data["json"] = err.Error()
	}
	this.ServeJson()
}

// @Title Delete
// @Description delete the Photos
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (this *PhotosController) Delete() {
	idStr := this.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	if err := models.DeletePhotos(id); err == nil {
		this.Data["json"] = "OK"
	} else {
		this.Data["json"] = err.Error()
	}
	this.ServeJson()
}
