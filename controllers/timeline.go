package controllers

import (
	"pet/models"
	"web"

	"github.com/astaxie/beego"
)

// oprations for Timeline
type TimelineController struct {
	beego.Controller
}

func (this *TimelineController) URLMapping() {
	this.Mapping("Post", this.Post)
	this.Mapping("GetOne", this.GetOne)
	this.Mapping("GetAll", this.GetAll)
	this.Mapping("Put", this.Put)
	this.Mapping("Delete", this.Delete)
}

// @Title Post
// @Description create Timeline
// @Param	body		body 	models.Timeline	true		"body for Timeline content"
// @Success 200 {int} models.Timeline.Id
// @Failure 403 body is empty
// @router / [post]

func (this *TimelineController) Post() {
	//var v models.Timeline
	//json.Unmarshal(this.Ctx.Input.RequestBody, &v)
	//if id, err := models.AddTimeline(&v); err == nil {
	//this.Data["json"] = map[string]int64{"id": id}
	//} else {
	//this.Data["json"] = err.Error()
	//}
	//this.ServeJson()
}

// @Title Get
// @Description get Timeline by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Timeline
// @Failure 403 :id is empty
// @router /:id [get]

func (this *TimelineController) GetOne() {
	//idStr := this.Ctx.Input.Params[":id"]
	//id, _ := strconv.Atoi(idStr)
	//v, err := models.GetTimelineById(id)
	//if err != nil {
	//this.Data["json"] = err.Error()
	//} else {
	//this.Data["json"] = v
	//}
	//this.ServeJson()
}

// @Title 获取照片timeline
// @Description get Timeline
// @Param	offset	query	string	false	"结果列表索引"
// @Success 200 {object} models.Timeline
// @Failure 403
// @router / [get]
func (this *TimelineController) GetAll() {

	if v, err := this.GetInt("offset"); err == nil {
		offset = v
	}

	userSession := this.GetSession("user").(models.Users)
	userId := userSession.Id

	l, err := models.GetFollowingPhotos(userId, offset)

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
// @Description update the Timeline
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Timeline	true		"body for Timeline content"
// @Success 200 {object} models.Timeline
// @Failure 403 :id is not int
// @router /:id [put]

func (this *TimelineController) Put() {
	//idStr := this.Ctx.Input.Params[":id"]
	//id, _ := strconv.Atoi(idStr)
	//v := models.Timeline{Id: id}
	//json.Unmarshal(this.Ctx.Input.RequestBody, &v)
	//if err := models.UpdateTimelineById(&v); err == nil {
	//this.Data["json"] = "OK"
	//} else {
	//this.Data["json"] = err.Error()
	//}
	//this.ServeJson()
}

// @Title Delete
// @Description delete the Timeline
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]

func (this *TimelineController) Delete() {
	//idStr := this.Ctx.Input.Params[":id"]
	//id, _ := strconv.Atoi(idStr)
	//if err := models.DeleteTimeline(id); err == nil {
	//this.Data["json"] = "OK"
	//} else {
	//this.Data["json"] = err.Error()
	//}
	//this.ServeJson()
}
