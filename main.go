package main

import (
	"encoding/gob"
	"encoding/json"
	_ "pet/docs"
	"pet/models"
	_ "pet/routers"
	"pet/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/mysql"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	dbAddress, _ := beego.Config("String", "DbAddress", "")
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", dbAddress.(string))
	orm.Debug = true

	beego.EnableAdmin = true
	beego.SetLevel(beego.LevelDebug)
	beego.SetLogFuncCall(true)
	gob.Register(models.Users{})
	gob.Register(models.UserPosition{})

}
func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.SetStaticPath("/doc", "static/swagger")
	}
	beego.SetStaticPath("/s", "static/source")

	beego.EnableDocs = true
	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)
	beego.Run()
}

var FilterUser = func(ctx *context.Context) {
	user := ctx.Input.Session("user")
	if user == nil && ctx.Request.URL.Path != "/v1/users/login" && ctx.Request.URL.Path[:4] != "/doc" && ctx.Request.URL.Path != "/v1/users/register" {
		outPut := helper.Reponse(1, nil, "请先登录")
		b, _ := json.Marshal(outPut)
		ctx.Output.Header("Access-Control-Allow-Origin", "*")
		ctx.WriteString(string(b))
	}
	ctx.Output.Header("Access-Control-Allow-Origin", "*")
}
