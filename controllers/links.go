package controllers

import (
	"pet/utils"

	"github.com/astaxie/beego"
)

// 获取动态链接
type LinksController struct {
	beego.Controller
}

func (c *LinksController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Appstore链接
// @Description 获取appstore链接
// @Success 200 {object} models.Links
// @Failure 403 :id is empty
// @router /appstore
func (this *LinksController) AppStore() {

	resp := make(map[string]interface{})
	resp["link"] = "http://apple.com.cn"
	outPut := helper.Reponse(0, resp, "")
	this.Data["json"] = outPut
	this.ServeJson()
}

// @Title Post
// @Description create Links
// @Param	body		body 	models.Links	true		"body for Links content"
// @Success 200 {int} models.Links.Id
// @Failure 403 body is empty
// @router / [post]

func (c *LinksController) Post() {

}

// @Title Get
// @Description get Links by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Links
// @Failure 403 :id is empty
// @router /:id [get]

func (c *LinksController) GetOne() {

}

// @Title Get All
// @Description get Links
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Links
// @Failure 403
// @router / [get]

func (c *LinksController) GetAll() {

}

// @Title Update
// @Description update the Links
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Links	true		"body for Links content"
// @Success 200 {object} models.Links
// @Failure 403 :id is not int
// @router /:id [put]

func (c *LinksController) Put() {

}

// @Title Delete
// @Description delete the Links
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]

func (c *LinksController) Delete() {

}
