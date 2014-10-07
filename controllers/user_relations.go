package controllers

import (
	"encoding/json"
	"errors"
	"pet/models"
	"strconv"
	"strings"
	"time"
	"web"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

// oprations for UserRelations
type UserRelationsController struct {
	beego.Controller
}

func (this *UserRelationsController) URLMapping() {
	this.Mapping("Post", this.Post)
	this.Mapping("GetOne", this.GetOne)
	this.Mapping("GetAll", this.GetAll)
	this.Mapping("Put", this.Put)
	this.Mapping("Delete", this.Delete)
}

// @Title Post
// @Description create UserRelations
// @Param	body		body 	models.UserRelations	true		"body for UserRelations content"
// @Success 200 {int} models.UserRelations.Id
// @Failure 403 body is empty
// @router / [post]
func (this *UserRelationsController) Post() {
	var v models.UserRelations

	valid := validation.Validation{}
	this.ParseForm(&v)
	v.Follower = this.GetSession("user").(models.Users).Id
	passed, _ := valid.Valid(&v)
	if !passed {
		outPut := helper.Reponse(1, nil, valid.Errors[0].Key+" "+valid.Errors[0].Message)
		this.Data["json"] = outPut
	} else {

		if v.Follower == v.Following {
			outPut := helper.Reponse(1, nil, "无效的关系")
			this.Data["json"] = outPut
		} else {
			v.CreatedAt = time.Now()
			v.UpdatedAt = time.Now()

			if id, err := models.AddUserRelations(&v); err == nil {
				v.Id = int(id)
				where := make(map[string]string)
				where["id"] = strconv.Itoa(v.Follower)
				helper.AddOne("users", "following", where)
				where["id"] = strconv.Itoa(v.Following)
				helper.AddOne("users", "follower", where)

				outPut := helper.Reponse(0, v, "创建成功")
				this.Data["json"] = outPut
			} else {
				outPut := helper.Reponse(1, nil, err.Error())
				this.Data["json"] = outPut
			}
		}
	}

	//json.Unmarshal(this.Ctx.Input.RequestBody, &v)
	//if id, err := models.AddUserRelations(&v); err == nil {
	//this.Data["json"] = map[string]int64{"id": id}
	//} else {
	//this.Data["json"] = err.Error()
	//}
	this.ServeJson()
}

// @Title Get
// @Description get UserRelations by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.UserRelations
// @Failure 403 :id is empty
// @router /:id [get]
func (this *UserRelationsController) GetOne() {
	idStr := this.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetUserRelationsById(id)
	if err != nil {
		this.Data["json"] = err.Error()
	} else {
		this.Data["json"] = v
	}
	this.ServeJson()
}

// @Title Get All
// @Description get UserRelations
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.UserRelations
// @Failure 403
// @router / [get]
func (this *UserRelationsController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query map[string]string = make(map[string]string)
	var limit int64 = 10
	var offset int64 = 0

	// fields: col1,col2,entity.col3
	if v := this.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := this.GetInt("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := this.GetInt("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := this.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := this.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := this.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.Split(cond, ":")
			if len(kv) != 2 {
				this.Data["json"] = errors.New("Error: invalid query key/value pair")
				this.ServeJson()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllUserRelations(query, fields, sortby, order, offset, limit)
	if err != nil {
		this.Data["json"] = err.Error()
	} else {
		this.Data["json"] = l
	}
	this.ServeJson()
}

// @Title Update
// @Description update the UserRelations
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.UserRelations	true		"body for UserRelations content"
// @Success 200 {object} models.UserRelations
// @Failure 403 :id is not int
// @router /:id [put]
func (this *UserRelationsController) Put() {
	idStr := this.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	v := models.UserRelations{Id: id}
	json.Unmarshal(this.Ctx.Input.RequestBody, &v)
	if err := models.UpdateUserRelationsById(&v); err == nil {
		this.Data["json"] = "OK"
	} else {
		this.Data["json"] = err.Error()
	}
	this.ServeJson()
}

// @Title Delete
// @Description delete the UserRelations
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (this *UserRelationsController) Delete() {
	var v models.UserRelations

	followingIdStr := this.Ctx.Input.Params[":id"]
	followingId, _ := strconv.Atoi(followingIdStr)

	valid := validation.Validation{}
	this.ParseForm(&v)
	v.Follower = this.GetSession("user").(models.Users).Id
	v.Following = followingId

	passed, _ := valid.Valid(&v)
	if !passed {
		outPut := helper.Reponse(1, nil, valid.Errors[0].Key+" "+valid.Errors[0].Message)
		this.Data["json"] = outPut
	} else {
		if v.Follower == v.Following {
			outPut := helper.Reponse(1, nil, "无效的关系")
			this.Data["json"] = outPut
		} else {
			v.CreatedAt = time.Now()
			v.UpdatedAt = time.Now()

			if err := models.DeleteUserRelationsByUsers(v.Follower, followingId); err == nil {
				where := make(map[string]string)
				where["id"] = strconv.Itoa(v.Follower)
				helper.MinusOne("users", "following", where)
				where["id"] = strconv.Itoa(v.Following)
				helper.MinusOne("users", "follower", where)

				outPut := helper.Reponse(0, v, "取消关注成功")
				this.Data["json"] = outPut
			} else {
				outPut := helper.Reponse(1, nil, err.Error())
				this.Data["json"] = outPut
			}
		}
	}
	this.ServeJson()
}
