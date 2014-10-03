package main

import (
	"encoding/gob"
	"encoding/json"
	_ "pet/docs"
	"pet/models"
	_ "pet/routers"
	"web"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/mysql"
	_ "github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
)

var dbAddress string = "root:38143195@tcp(192.168.33.11:3306)/pet?charset=utf8"

func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", dbAddress)
	orm.Debug = true
	beego.SetLevel(beego.LevelInformational)
	beego.SessionOn = true
	beego.SessionProvider = "file"
	beego.SessionSavePath = "./tmp"

	gob.Register(models.Users{})

}
func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/static/swagger"] = "swagger"
	}

	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)
	beego.Run()
}

var FilterUser = func(ctx *context.Context) {
	user := ctx.Input.Session("user")

	if user == nil && ctx.Request.URL.Path != "/v1/users/login" {
		outPut := helper.Reponse(1, nil, "请先登录")
		b, _ := json.Marshal(outPut)
		ctx.WriteString(string(b))
	}
}
