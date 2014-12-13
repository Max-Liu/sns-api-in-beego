package controllers

import (
	"pet/models"
	"pet/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

// 意见反馈
type FeedbackController struct {
	beego.Controller
}

func (c *FeedbackController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title 反馈
// @Description create 反馈意见
// @Param	content		form 	String	true		"反馈内容"
// @Success 200 {int} models.Feedback.Id
// @Failure 403 body is empty
// @router / [post]
func (this *FeedbackController) Post() {

	var v models.Feedback
	valid := validation.Validation{}
	this.ParseForm(&v)

	passed, _ := valid.Valid(&v)
	if !passed {
		outPut := helper.Reponse(1, nil, valid.Errors[0].Key+" "+valid.Errors[0].Message)
		this.Data["json"] = outPut
	} else {
		currentUser := this.GetSession("user").(models.Users)
		v.UserId = currentUser.Id
		_, err := models.AddFeedback(&v)
		if err != nil {
			outPut := helper.Reponse(1, nil, err.Error())
			this.Data["json"] = outPut
			this.ServeJson()
			return
		}

		outPut := helper.Reponse(0, nil, "提交成功")
		this.Data["json"] = outPut
	}
	this.ServeJson()
}

// @Title Get
// @Description get Feedback by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Feedback
// @Failure 403 :id is empty
// @router /:id [get]

func (c *FeedbackController) GetOne() {

}

// @Title Get All
// @Description get Feedback
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Feedback
// @Failure 403
// @router / [get]

func (c *FeedbackController) GetAll() {

}

// @Title Update
// @Description update the Feedback
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Feedback	true		"body for Feedback content"
// @Success 200 {object} models.Feedback
// @Failure 403 :id is not int
// @router /:id [put]

func (c *FeedbackController) Put() {

}

// @Title Delete
// @Description delete the Feedback
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]

func (c *FeedbackController) Delete() {

}
