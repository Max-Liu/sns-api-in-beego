package controllers

import (
	"os"
	"pet/models"
	"pet/utils"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

// 用户信息
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

// @Title 注册
// @Description 用户注册
// @Param	email		form 	String	true		"注册邮箱"
// @Param	phone		form 	String	true		"注册手机"
// @Param	name		form 	String	true		"用户名"
// @Param	password	form 	String	true		"密码"
// @Success 200 {int} models.Users.Id
// @Failure 403 body is empty
// @router /register [post]
func (this *UsersController) Register() {
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
			v.Id = id
			data := models.ConverToUserApiStruct(&v)
			outPut := helper.Reponse(0, data, "创建成功")
			this.Data["json"] = outPut
		} else {
			outPut := helper.Reponse(1, nil, err.Error())
			this.Data["json"] = outPut
		}
	}
	this.ServeJson()
}

// @Title 更新个人信息
// @Description 更新个人信息
// @Param	id			path 	string	true		"用户ID"
// @Param	gender		form	Int		true		"用户性别"
// @Param	birthday	form	String	true		"生日"
// @Param	head		form	File	true		"头像图片"
// @Success 200 {int} models.Users.Id
// @Failure 403 body is empty
// @router /:id [put]
func (this *UsersController) Put() {

	idStr := this.Ctx.Input.Params[":id"]
	id, _ := strconv.ParseInt(idStr, 10, 0)

	userSession := this.GetSession("user")
	user := userSession.(models.Users)

	if userSession.(models.Users).Id == id {
		var v models.Users
		valid := validation.Validation{}
		v = user

		passed, _ := valid.Valid(&v)
		if !passed {
			outPut := helper.Reponse(1, nil, valid.Errors[0].Key+" "+valid.Errors[0].Message)
			this.Data["json"] = outPut
		} else {

			v, _ := models.GetUsersById(id)
			v.Birthday = this.GetString("birthday")
			genderStr := this.GetString("gender")
			gender, _ := strconv.Atoi(genderStr)
			v.Gender = gender

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
			err := this.SaveToFile("head", imagePath)

			if err != nil {
				outPut := helper.Reponse(1, nil, err.Error())
				this.Data["json"] = outPut
			} else {
				v.Head = imagePath
				if err := models.UpdateUsersById(v); err == nil {
					data := models.ConverToUserApiStruct(v)
					outPut := helper.Reponse(0, data, "更新成功")
					this.Data["json"] = outPut
				} else {
					outPut := helper.Reponse(1, nil, err.Error())
					this.Data["json"] = outPut
				}
			}
		}
	} else {
		outPut := helper.Reponse(1, nil, "没有权限更新其它用户")
		this.Data["json"] = outPut
	}

	this.ServeJson()

}

//idStr := this.Ctx.Input.Params[":id"]
//id, _ := strconv.Atoi(idStr)
//json.Unmarshal(this.Ctx.Input.RequestBody, &v)
//this.ServeJson()
//}

// @Title 获取单个用户信息
// @Description 获取单个用户信息
// @Param	id		path 	string	true		"用户ID"
// @Success 200 {object} models.Users
// @Failure 403 :id is empty
// @router /:id [get]
func (this *UsersController) GetOne() {
	idStr := this.Ctx.Input.Params[":id"]
	id, _ := strconv.ParseInt(idStr, 10, 0)
	v, err := models.GetUsersById(id)
	if err != nil {
		outPut := helper.Reponse(1, nil, "")
		this.Data["json"] = outPut
	} else {
		data := models.ConverToUserApiStruct(v)
		outPut := helper.Reponse(0, data, "")
		this.Data["json"] = outPut
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
	//idStr := this.Ctx.Input.Params[":id"]
	//id, _ := strconv.Atoi(idStr)
	//if err := models.DeleteUsers(id); err == nil {
	//this.Data["json"] = "OK"
	//} else {
	//this.Data["json"] = err.Error()
	//}
	//this.ServeJson()
}

// @Title 登陆
// @Description 登陆
// @Param	info		query	String	true		"用户名或手机或邮箱"
// @Param	password	query	String	true		"密码"
// @Success 200 {int} models.Users.Id
// @Failure 403 body is empty
// @router /login
func (this *UsersController) Login() {
	if user := this.GetSession("user"); user != nil {
		outPut := helper.Reponse(0, &user, "登陆成功")
		this.Data["json"] = outPut
	} else {
		info := this.GetString("info")
		password := this.GetString("password")

		email, phone, name := info, info, info

		user, err := models.GetUserByLoginfo(password, email, phone, name)

		if err == nil {
			data := models.ConverToUserApiStruct(&user)
			outPut := helper.Reponse(0, data, "登陆成功")
			this.SetSession("user", user)
			this.Data["json"] = outPut
		} else {
			outPut := helper.Reponse(1, nil, "用户信息不存在")
			this.Data["json"] = outPut
		}
	}
	this.ServeJson()
}

// @Title 注销
// @Description 注销
// @Success 200 {int} models.Users.Id
// @Failure 403 body is empty
// @router /logout
func (this *UsersController) Logout() {
	this.DelSession("user")
	outPut := helper.Reponse(0, nil, "登出成功")
	this.Data["json"] = outPut
	this.ServeJson()
}

// @Title 发送地理位置
// @Description 发送地理位置
// @Success 200 {int} models.Users.Id
// @Failure 403 body is empty
// @Param	x	query	string	false	"坐标x"
// @Param	y	query	string	false	"坐标y"
// @Param	current_time	query	string	false	"当前时间"
// @router /send_postion [post]
func (this *UsersController) CurrentPostion() {
	x, _ := this.GetFloat("x")
	y, _ := this.GetFloat("y")
	currentTime, _ := this.GetInt64("current_time")
	position := new(models.UserPosition)
	position.X = x
	position.Y = y
	position.CurrentTime = currentTime

	this.SetSession("userPosition", *position)
	outPut := helper.Reponse(0, nil, "")
	this.Data["json"] = outPut
	this.ServeJson()
}
