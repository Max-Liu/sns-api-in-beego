package controllers

import (
	"github.com/astaxie/beego"
)

// oprations for Top10photo
type Top10photoController struct {
	beego.Controller
}

func (c *Top10photoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create Top10photo
// @Param	body		body 	models.Top10photo	true		"body for Top10photo content"
// @Success 200 {int} models.Top10photo.Id
// @Failure 403 body is empty
// @router / [post]
func (c *Top10photoController) Post() {

}

// @Title Get
// @Description get Top10photo by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Top10photo
// @Failure 403 :id is empty
// @router /:id [get]
func (c *Top10photoController) GetOne() {

}

// @Title Get All
// @Description get Top10photo
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Top10photo
// @Failure 403
// @router / [get]
func (c *Top10photoController) GetAll() {

}

// @Title Update
// @Description update the Top10photo
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Top10photo	true		"body for Top10photo content"
// @Success 200 {object} models.Top10photo
// @Failure 403 :id is not int
// @router /:id [put]
func (c *Top10photoController) Put() {

}

// @Title Delete
// @Description delete the Top10photo
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *Top10photoController) Delete() {

}
