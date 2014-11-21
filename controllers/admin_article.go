package controllers

import (
	"pet/models"
	"time"

	"github.com/astaxie/beego"
)

// oprations for Admin_article
type Admin_articleController struct {
	beego.Controller
}

func (c *Admin_articleController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create Admin_article
// @Param	body		body 	models.Admin_article	true		"body for Admin_article content"
// @Success 200 {int} models.Admin_article.Id
// @Failure 403 body is empty
// @router / [post]
func (c *Admin_articleController) Post() {

}

// @Title Get
// @Description get Admin_article by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Admin_article
// @Failure 403 :id is empty
// @router /:id [get]
func (c *Admin_articleController) GetOne() {

}

// @Title Get All
// @Description get Admin_article
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Admin_article
// @Failure 403
// @router / [get]
func (c *Admin_articleController) GetAll() {
	var fields []string
	query := make(map[string]string)

	if v, err := c.GetInt("offset"); err == nil {
		offset = int64(v)
	}
	l, _ := models.GetAllArticles(query, fields, sortby, order, offset, limit)

	var articalDatas []models.Articles
	var artical models.Articles

	for _, v := range l {
		artical.Id = v["Id"].(int64)
		artical.Content = v["Content"].(string)
		artical.Title = v["Title"].(string)
		artical.CreatedAt = v["CreatedAt"].(time.Time)
		artical.TitleImage = v["TitleImage"].(string)
		articalDatas = append(articalDatas, artical)
	}
	c.Data["articles"] = &articalDatas
	c.Layout = "admin/layout.html"
	c.TplNames = "admin/article/index.tpl"
	c.Render()
}

// @Title Update
// @Description update the Admin_article
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Admin_article	true		"body for Admin_article content"
// @Success 200 {object} models.Admin_article
// @Failure 403 :id is not int
// @router /:id [put]
func (c *Admin_articleController) Put() {

}

// @Title Delete
// @Description delete the Admin_article
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *Admin_articleController) Delete() {

}
