package controllers

import (
	"pet/models"
	"time"

	"github.com/astaxie/beego"
)

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
func (c *Admin_articleController) Post() {

}

func (c *Admin_articleController) GetOne() {

}

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

func (c *Admin_articleController) Put() {

}

func (c *Admin_articleController) Delete() {

}
