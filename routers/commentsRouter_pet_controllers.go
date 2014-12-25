package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["pet/controllers:ArticlesController"] = append(beego.GlobalControllerRouter["pet/controllers:ArticlesController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:ArticlesController"] = append(beego.GlobalControllerRouter["pet/controllers:ArticlesController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:ArticlesController"] = append(beego.GlobalControllerRouter["pet/controllers:ArticlesController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:MsgController"] = append(beego.GlobalControllerRouter["pet/controllers:MsgController"],
		beego.ControllerComments{
			"Me",
			`/me`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:MsgController"] = append(beego.GlobalControllerRouter["pet/controllers:MsgController"],
		beego.ControllerComments{
			"GetFollowingPhotosTimeline",
			`/following`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:MsgController"] = append(beego.GlobalControllerRouter["pet/controllers:MsgController"],
		beego.ControllerComments{
			"HasNotice",
			`/has_notice`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:LinksController"] = append(beego.GlobalControllerRouter["pet/controllers:LinksController"],
		beego.ControllerComments{
			"AppStore",
			`/appstore`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:Top10photoController"] = append(beego.GlobalControllerRouter["pet/controllers:Top10photoController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:Top10photoController"] = append(beego.GlobalControllerRouter["pet/controllers:Top10photoController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:Top10photoController"] = append(beego.GlobalControllerRouter["pet/controllers:Top10photoController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:Top10photoController"] = append(beego.GlobalControllerRouter["pet/controllers:Top10photoController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:Top10photoController"] = append(beego.GlobalControllerRouter["pet/controllers:Top10photoController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:UserRelationsController"] = append(beego.GlobalControllerRouter["pet/controllers:UserRelationsController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:UserRelationsController"] = append(beego.GlobalControllerRouter["pet/controllers:UserRelationsController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:UserRelationsController"] = append(beego.GlobalControllerRouter["pet/controllers:UserRelationsController"],
		beego.ControllerComments{
			"Follower",
			`/follower`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:UserRelationsController"] = append(beego.GlobalControllerRouter["pet/controllers:UserRelationsController"],
		beego.ControllerComments{
			"Following",
			`/following`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:UserRelationsController"] = append(beego.GlobalControllerRouter["pet/controllers:UserRelationsController"],
		beego.ControllerComments{
			"HasFollowed",
			`/has_followed`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:UsersController"] = append(beego.GlobalControllerRouter["pet/controllers:UsersController"],
		beego.ControllerComments{
			"Register",
			`/register`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:UsersController"] = append(beego.GlobalControllerRouter["pet/controllers:UsersController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:UsersController"] = append(beego.GlobalControllerRouter["pet/controllers:UsersController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:UsersController"] = append(beego.GlobalControllerRouter["pet/controllers:UsersController"],
		beego.ControllerComments{
			"Login",
			`/login`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:UsersController"] = append(beego.GlobalControllerRouter["pet/controllers:UsersController"],
		beego.ControllerComments{
			"Logout",
			`/logout`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:UsersController"] = append(beego.GlobalControllerRouter["pet/controllers:UsersController"],
		beego.ControllerComments{
			"CurrentPostion",
			`/send_postion`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:FeedbackController"] = append(beego.GlobalControllerRouter["pet/controllers:FeedbackController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:LikesController"] = append(beego.GlobalControllerRouter["pet/controllers:LikesController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:LikesController"] = append(beego.GlobalControllerRouter["pet/controllers:LikesController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:LikesController"] = append(beego.GlobalControllerRouter["pet/controllers:LikesController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:LikesController"] = append(beego.GlobalControllerRouter["pet/controllers:LikesController"],
		beego.ControllerComments{
			"UsersList",
			`/users`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:PhotoCommentsController"] = append(beego.GlobalControllerRouter["pet/controllers:PhotoCommentsController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:PhotoCommentsController"] = append(beego.GlobalControllerRouter["pet/controllers:PhotoCommentsController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:PhotosController"] = append(beego.GlobalControllerRouter["pet/controllers:PhotosController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:PhotosController"] = append(beego.GlobalControllerRouter["pet/controllers:PhotosController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:PhotosController"] = append(beego.GlobalControllerRouter["pet/controllers:PhotosController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:PhotosController"] = append(beego.GlobalControllerRouter["pet/controllers:PhotosController"],
		beego.ControllerComments{
			"GetTop10",
			`/top10`,
			[]string{"get"},
			nil})

}
