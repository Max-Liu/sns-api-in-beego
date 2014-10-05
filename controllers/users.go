package controllers

import (
	"pet/models"
	"strconv"
	"time"
	"web"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

// oprations for Users
type UsersController struct {
	beego.Controller
}

func (this *UsersController) URLMapping() {
	this.Mapping("Post", this.Post)
	this.Mapping("GetOne", this.GetOne)
	this.Mapping("GetAll", this.GetAll)
	this.Mapping("Put", this.Put)
	this.Mapping("Delete", this.Delete)
}

// @Title Post
// @Description create Users
// @Param	body		body 	models.Users	true		"body for Users content"
// @Success 200 {int} models.Users.Id
// @Failure 403 body is empty
// @router / [post]
func (this *UsersController) Post() {
	var v models.Users
	valid := validation.Validation{}
	this.ParseForm(&v)
	passed, _ := valid.Valid(&v)
	if !passed {
		outPut := helper.Reponse(1, nil, valid.Errors[0].Key+" "+valid.Errors[0].Message)
		this.Data["json"] = outPut
	} else {
		v.CreatedAt = time.Now()
		v.UpdatedAt = time.Now()

		if id, err := models.AddUsers(&v); err == nil {
			v.Id = int(id)
			outPut := helper.Reponse(0, v, "创建成功")
			this.Data["json"] = outPut
		} else {
			outPut := helper.Reponse(1, nil, err.Error())
			this.Data["json"] = outPut
		}
	}
	this.ServeJson()
}

// @Title Update
// @Description update the Users
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Users	true		"body for Users content"
// @Success 200 {object} models.Users
// @Failure 403 :id is not int
// @router /:id [put]
func (this *UsersController) Put() {

	idStr := this.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)

	userSession := this.GetSession("user")

	if userSession.(models.Users).Id == id {
		v := models.Users{Id: id}
		this.ParseForm(&v)
		valid := validation.Validation{}
		passed, _ := valid.Valid(&v)

		if !passed {
			outPut := helper.Reponse(1, nil, valid.Errors[0].Key+" "+valid.Errors[0].Message)
			this.Data["json"] = outPut
		} else {
			if err := models.UpdateUsersById(&v); err == nil {
				outPut := helper.Reponse(0, v, "更新成功")
				this.Data["json"] = outPut
			} else {
				outPut := helper.Reponse(1, nil, err.Error())
				this.Data["json"] = outPut
			}
		}
	} else {
		outPut := helper.Reponse(1, nil, "no access")
		this.Data["json"] = outPut
	}

	this.ServeJson()

}

//idStr := this.Ctx.Input.Params[":id"]
//id, _ := strconv.Atoi(idStr)
//json.Unmarshal(this.Ctx.Input.RequestBody, &v)
//this.ServeJson()
//}

// @Title Get
// @Description get Users by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Users
// @Failure 403 :id is empty
// @router /:id [get]
func (this *UsersController) GetOne() {
	idStr := this.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetUsersById(id)
	if err != nil {
		this.Data["json"] = err.Error()
	} else {
		this.Data["json"] = v
	}
	this.ServeJson()
}

// @Title Get All
// @Description get Users
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Users
// @Failure 403
// @router / [get]
func (this *UsersController) GetAll() {
	user := this.GetSession("user")
	outPut := helper.Reponse(0, user, "")
	this.Data["json"] = outPut
	//var fields []string
	//var sortby []string
	//var order []string
	//var query map[string]string = make(map[string]string)
	//var limit int64 = 10
	//var offset int64 = 0

	//// fields: col1,col2,entity.col3
	//if v := this.GetString("fields"); v != "" {
	//fields = strings.Split(v, ",")
	//}
	//// limit: 10 (default is 10)
	//if v, err := this.GetInt("limit"); err == nil {
	//limit = v
	//}
	//// offset: 0 (default is 0)
	//if v, err := this.GetInt("offset"); err == nil {
	//offset = v
	//}
	//// sortby: col1,col2
	//if v := this.GetString("sortby"); v != "" {
	//sortby = strings.Split(v, ",")
	//}
	//// order: desc,asc
	//if v := this.GetString("order"); v != "" {
	//order = strings.Split(v, ",")
	//}
	//// query: k:v,k:v
	//if v := this.GetString("query"); v != "" {
	//for _, cond := range strings.Split(v, ",") {
	//kv := strings.Split(cond, ":")
	//if len(kv) != 2 {
	//this.Data["json"] = errors.New("Error: invalid query key/value pair")
	//this.ServeJson()
	//return
	//}
	//k, v := kv[0], kv[1]
	//query[k] = v
	//}
	//}

	//l, err := models.GetAllUsers(query, fields, sortby, order, offset, limit)
	//if err != nil {
	//this.Data["json"] = err.Error()
	//} else {
	//this.Data["json"] = l
	//}
	this.ServeJson()
}

// @Title Delete
// @Description delete the Users
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (this *UsersController) Delete() {
	idStr := this.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteUsers(id); err == nil {
		this.Data["json"] = "OK"
	} else {
		this.Data["json"] = err.Error()
	}
	this.ServeJson()
}

func (this *UsersController) Login() {
	if user := this.GetSession("user"); user != nil {
		outPut := helper.Reponse(0, &user, "登陆成功")
		this.Data["json"] = outPut
	} else {

		email := this.GetString("email")
		password := this.GetString("password")
		phone := this.GetString("phone")
		name := this.GetString("name")
		user, err := models.GetUserByLoginfo(password, email, phone, name)

		if err == nil {
			outPut := helper.Reponse(0, &user, "登陆成功")
			this.SetSession("user", user)
			this.Data["json"] = outPut
		} else {
			outPut := helper.Reponse(1, nil, "用户信息不存在")
			this.Data["json"] = outPut
		}
	}
	this.ServeJson()
}

func (this *UsersController) Logout() {
	this.DelSession("user")
	outPut := helper.Reponse(0, nil, "登出成功")
	this.Data["json"] = outPut
	this.ServeJson()
}
