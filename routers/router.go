// @APIVersion 1.0.0
// @Title Pet Rest API
package routers

import (
	"pet/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/articles",
			beego.NSInclude(
				&controllers.ArticlesController{},
			),
		),

		beego.NSNamespace("/likes",
			beego.NSRouter("/users", &controllers.LikesController{}, "get:UsersList"),
			beego.NSInclude(
				&controllers.LikesController{},
			),
		),

		beego.NSNamespace("/comments",
			beego.NSInclude(
				&controllers.PhotoCommentsController{},
			),
		),

		beego.NSNamespace("/photos",
			beego.NSRouter("/following", &controllers.PhotosController{}, "get:GetFollowingPhotosTimeline"),
			beego.NSRouter("/top10", &controllers.PhotosController{}, "get:GetTop10"),
			beego.NSInclude(
				&controllers.PhotosController{},
			),
		),

		beego.NSNamespace("/ul",
			beego.NSRouter("/follower", &controllers.UserRelationsController{}, "get:Follower"),
			beego.NSRouter("/following", &controllers.UserRelationsController{}, "get:Following"),
			beego.NSInclude(
				&controllers.UserRelationsController{},
			),
		),

		beego.NSNamespace("/msg",
			beego.NSInclude(
				&controllers.MsgController{},
			),
		),

		beego.NSNamespace("/feedback",
			beego.NSInclude(
				&controllers.FeedbackController{},
			),
		),
		beego.NSNamespace("/users",
			beego.NSRouter("/login", &controllers.UsersController{}, "get:Login"),
			beego.NSRouter("/logout", &controllers.UsersController{}, "get:Logout"),
			beego.NSRouter("/register", &controllers.UsersController{}, "post:Register"),
			beego.NSRouter("/send_position", &controllers.UsersController{}, "post:CurrentPostion"),
			beego.NSInclude(
				&controllers.UsersController{},
			),
		),
	)
	//admin := beego.NewNamespace("/admin",

	//beego.NSNamespace("/articles",
	//beego.NSInclude(
	////&controllers.Admin_articleController{},
	//),
	//),
	//)

	//beego.AddNamespace(admin)
	beego.AddNamespace(ns)
}
