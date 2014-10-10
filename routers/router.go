// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
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

		beego.NSNamespace("/images",
			beego.NSInclude(
				&controllers.ImagesController{},
			),
		),

		beego.NSNamespace("/likes",
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
			beego.NSInclude(
				&controllers.PhotosController{},
			),
		),

		beego.NSNamespace("/timeline",
			beego.NSInclude(
				&controllers.TimelineController{},
			),
		),

		beego.NSNamespace("/ul",
			beego.NSRouter("/follower", &controllers.UserRelationsController{}, "get:Follower"),
			beego.NSRouter("/following", &controllers.UserRelationsController{}, "get:Following"),
			beego.NSInclude(
				&controllers.UserRelationsController{},
			),
		),

		beego.NSNamespace("/users",
			beego.NSRouter("/login", &controllers.UsersController{}, "get:Login"),
			beego.NSRouter("/logout", &controllers.UsersController{}, "get:Logout"),
			beego.NSInclude(
				&controllers.UsersController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
