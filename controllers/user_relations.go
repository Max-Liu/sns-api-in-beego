package controllers

import (
	"pet/models"
	"pet/utils"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/beego/redigo/redis"
)

// 用户关系相关
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

// @Title 关注
// @Description 关注某用户
// @Param	following		form 	String	true		"所关注用户的Id"
// @Success 200 {int} models.UserRelations.Id
// @Failure 403 body is empty
// @router / [post]
func (this *UserRelationsController) Post() {
	var v models.UserRelations

	valid := validation.Validation{}
	this.ParseForm(&v)

	//当前登陆用户ID
	follower := this.GetSession("user").(models.Users)
	v.Follower = &follower

	//关注目标的用户ID
	followingIdStr := this.GetString("following")
	followingId, _ := strconv.ParseInt(followingIdStr, 10, 0)

	v.Following, err = models.GetUsersById(followingId)
	if err != nil {
		outPut := helper.Reponse(1, nil, err.Error())
		this.Data["json"] = outPut
		this.ServeJson()
		return
	}
	if follower.Id == v.Following.Id {
		outPut := helper.Reponse(1, nil, "无效的关注对象")
		this.Data["json"] = outPut
		this.ServeJson()
		return
	}

	passed, _ := valid.Valid(&v)
	if !passed {
		outPut := helper.Reponse(1, nil, valid.Errors[0].Key+" "+valid.Errors[0].Message)
		this.Data["json"] = outPut
	} else {

		v.CreatedAt = time.Now()
		v.UpdatedAt = time.Now()

		redisAddress, _ := beego.Config("String", "redisServer", "")
		c, err := redis.Dial("tcp", redisAddress.(string))
		defer c.Close()
		if err != nil {
			beego.Error(err.Error())
		}

		if _, err := models.AddUserRelations(&v); err == nil {
			//add relations in redis
			if err != nil {
				beego.Error(err.Error())
			}
			followerIdStr := strconv.FormatInt(follower.Id, 10)
			result, err := c.Do("ZADD", "following:"+followerIdStr, time.Now().Unix(), followingIdStr)
			if err != nil {
				beego.Error(err.Error())
			}

			result, err = c.Do("ZCARD", "following:"+followerIdStr)
			beego.Debug(result)
			if err != nil {
				beego.Error(err.Error())
			}
			v.Follower.Following = result.(int64)
			models.UpdateUsersById(v.Follower)

			result, err = c.Do("ZADD", "follower:"+followingIdStr, time.Now().Unix(), followerIdStr)
			beego.Debug(result)
			if err != nil {
				beego.Error(err.Error())
			}

			result, err = c.Do("ZCARD", "follower:"+followingIdStr)
			beego.Debug(result)
			if err != nil {
				beego.Error(err.Error())
			}

			v.Following.Follower = result.(int64)
			models.UpdateUsersById(v.Following)
			data := models.ConverToUserRelationsApiStruct(&v)

			outPut := helper.Reponse(0, data, "创建成功")
			this.Data["json"] = outPut
		} else {
			outPut := helper.Reponse(1, nil, err.Error())
			this.Data["json"] = outPut
		}
	}

	this.ServeJson()
}

// @Title Get
// @Description get UserRelations by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.UserRelations
// @Failure 403 :id is empty
// @router /:id [get]

func (this *UserRelationsController) GetOne() {
	//idStr := this.Ctx.Input.Params[":id"]
	//id, _ := strconv.Atoi(idStr)
	//v, err := models.GetUserRelationsById(id)
	//if err != nil {
	//this.Data["json"] = err.Error()
	//} else {
	//this.Data["json"] = v
	//}
	//this.ServeJson()
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

	//l, err := models.GetAllUserRelations(query, fields, sortby, order, offset, limit)
	//if err != nil {
	//this.Data["json"] = err.Error()
	//} else {
	//this.Data["json"] = l
	//}
	//this.ServeJson()
}

// @Title Update
// @Description update the UserRelations
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.UserRelations	true		"body for UserRelations content"
// @Success 200 {object} models.UserRelations
// @Failure 403 :id is not int
// @router /:id [put]

func (this *UserRelationsController) Put() {
	//idStr := this.Ctx.Input.Params[":id"]
	//id, _ := strconv.Atoi(idStr)
	//v := models.UserRelations{Id: id}
	//json.Unmarshal(this.Ctx.Input.RequestBody, &v)
	//if err := models.UpdateUserRelationsById(&v); err == nil {
	//this.Data["json"] = "OK"
	//} else {
	//this.Data["json"] = err.Error()
	//}
	//this.ServeJson()
}

// @Title 取消关注
// @Description 取消关注
// @Param	id		path 	string	true		"所要取消关注的用户ID"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (this *UserRelationsController) Delete() {
	var v models.UserRelations

	followingIdStr := this.Ctx.Input.Params[":id"]
	followingId, _ := strconv.ParseInt(followingIdStr, 10, 0)

	valid := validation.Validation{}
	this.ParseForm(&v)
	follower := this.GetSession("user").(models.Users)
	v.Follower = &follower

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

			if _, err := models.DeleteUserRelationsByUsers(v.Follower.Id, followingId); err == nil {

				//delete relations in redis
				redisAddress, _ := beego.Config("String", "redisServer", "")
				c, err := redis.Dial("tcp", redisAddress.(string))
				defer c.Close()
				if err != nil {
					beego.Error(err.Error())
				}
				followerIdStr := strconv.FormatInt(follower.Id, 10)
				result, err := c.Do("ZREM", "following:"+followerIdStr, followingIdStr)
				beego.Debug(result)
				if err != nil {
					beego.Error(err.Error())
				}
				result, err = c.Do("ZREM", "follower:"+followingIdStr, followerIdStr)
				beego.Debug(result)
				if err != nil {
					beego.Error(err.Error())
				}

				where := make(map[string]string)
				where["id"] = strconv.FormatInt(v.Follower.Id, 10)
				helper.MinusOne("users", "following", where)
				where["id"] = strconv.FormatInt(followingId, 10)
				helper.MinusOne("users", "follower", where)

				outPut := helper.Reponse(0, nil, "取消关注成功")
				this.Data["json"] = outPut
			} else {
				outPut := helper.Reponse(1, nil, err.Error())
				this.Data["json"] = outPut
			}
		}
	}
	this.ServeJson()
}

// @Title 粉丝列表
// @Description 获取粉丝列表
// @Param	user_id	query	string	false	"目标用户id，为空代表获取当前用户"
// @Param	offset	query	string	false	"结果索引"
// @Success 200 {object} models.UserRelations
// @Failure 403
// @router /follower [get]
func (this *UserRelationsController) Follower() {
	if v, err := this.GetInt("offset"); err == nil {
		offset = int64(v)
	}
	query := make(map[string]string)
	if userIdStr := this.GetString("user_id"); userIdStr != "" {
		query["following"] = userIdStr
	} else {
		userSession := this.GetSession("user").(models.Users)
		userIdStr := strconv.FormatInt(userSession.Id, 10)
		query["following"] = userIdStr
	}
	fields := []string{"CreatedAt", "follower__id"}

	l, err := models.GetAllUserRelations(query, fields, sortby, order, offset, limit)

	var UserRelationsApiDatas []*models.UserRelationsFollowerApi
	var UserRelation models.UserRelations

	for _, v := range l {
		UserRelation.CreatedAt = v["CreatedAt"].(time.Time)
		UserRelation.Follower, _ = models.GetUsersById(v["Follower__Id"].(int64))
		UserRelationApiData := models.ConverToUserRelationsFollowerApirStruct(&UserRelation)
		UserRelationsApiDatas = append(UserRelationsApiDatas, UserRelationApiData)

	}

	if err != nil {
		outPut := helper.Reponse(1, nil, err.Error())
		this.Data["json"] = outPut
	} else {
		outPut := helper.Reponse(0, UserRelationsApiDatas, "")
		this.Data["json"] = outPut
	}
	this.ServeJson()
}

// @Title 关注列表
// @Description 获取关注列表
// @Param	user_id	query	string	false	"目标用户id,为空代表获取当前用户"
// @Param	offset	query	string	false	"结果索引"
// @Success 200 {object} models.UserRelations
// @Failure 403
// @router /following [get]
func (this *UserRelationsController) Following() {

	if v, err := this.GetInt("offset"); err == nil {
		offset = int64(v)
	}

	query := make(map[string]string)
	if userIdStr := this.GetString("user_id"); userIdStr != "" {
		query["follower"] = userIdStr
	} else {
		userSession := this.GetSession("user").(models.Users)
		userIdStr := strconv.FormatInt(userSession.Id, 10)
		query["follower"] = userIdStr
	}
	fields := []string{"CreatedAt", "following__name", "following__id"}

	l, err := models.GetAllUserRelations(query, fields, sortby, order, offset, limit)

	var UserRelationsApiDatas []*models.UserRelationsFollowingApi
	var UserRelation models.UserRelations

	for _, v := range l {
		UserRelation.CreatedAt = v["CreatedAt"].(time.Time)
		UserRelation.Following, _ = models.GetUsersById(v["Following__Id"].(int64))
		UserRelationApiData := models.ConverToUserRelationsFollowingApiStruct(&UserRelation)
		UserRelationsApiDatas = append(UserRelationsApiDatas, UserRelationApiData)
	}
	if err != nil {
		outPut := helper.Reponse(1, nil, err.Error())
		this.Data["json"] = outPut
	} else {
		outPut := helper.Reponse(0, UserRelationsApiDatas, "")
		this.Data["json"] = outPut
	}
	this.ServeJson()
}
