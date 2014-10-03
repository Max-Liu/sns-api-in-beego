package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["pet/controllers:UserRelationsController"] = append(beego.GlobalControllerRouter["pet/controllers:UserRelationsController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:UserRelationsController"] = append(beego.GlobalControllerRouter["pet/controllers:UserRelationsController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:UserRelationsController"] = append(beego.GlobalControllerRouter["pet/controllers:UserRelationsController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:UserRelationsController"] = append(beego.GlobalControllerRouter["pet/controllers:UserRelationsController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:UserRelationsController"] = append(beego.GlobalControllerRouter["pet/controllers:UserRelationsController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:UsersController"] = append(beego.GlobalControllerRouter["pet/controllers:UsersController"],
		beego.ControllerComments{
			"Post",
			`/`,
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
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:UsersController"] = append(beego.GlobalControllerRouter["pet/controllers:UsersController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:ObjectController"] = append(beego.GlobalControllerRouter["pet/controllers:ObjectController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:ObjectController"] = append(beego.GlobalControllerRouter["pet/controllers:ObjectController"],
		beego.ControllerComments{
			"Get",
			`/:objectId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:ObjectController"] = append(beego.GlobalControllerRouter["pet/controllers:ObjectController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:ObjectController"] = append(beego.GlobalControllerRouter["pet/controllers:ObjectController"],
		beego.ControllerComments{
			"Put",
			`/:objectId`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:ObjectController"] = append(beego.GlobalControllerRouter["pet/controllers:ObjectController"],
		beego.ControllerComments{
			"Delete",
			`/:objectId`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:TimelineController"] = append(beego.GlobalControllerRouter["pet/controllers:TimelineController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:TimelineController"] = append(beego.GlobalControllerRouter["pet/controllers:TimelineController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:TimelineController"] = append(beego.GlobalControllerRouter["pet/controllers:TimelineController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:TimelineController"] = append(beego.GlobalControllerRouter["pet/controllers:TimelineController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:TimelineController"] = append(beego.GlobalControllerRouter["pet/controllers:TimelineController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:LikesController"] = append(beego.GlobalControllerRouter["pet/controllers:LikesController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:LikesController"] = append(beego.GlobalControllerRouter["pet/controllers:LikesController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:LikesController"] = append(beego.GlobalControllerRouter["pet/controllers:LikesController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:LikesController"] = append(beego.GlobalControllerRouter["pet/controllers:LikesController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:LikesController"] = append(beego.GlobalControllerRouter["pet/controllers:LikesController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
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
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:PhotosController"] = append(beego.GlobalControllerRouter["pet/controllers:PhotosController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

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

	beego.GlobalControllerRouter["pet/controllers:ArticlesController"] = append(beego.GlobalControllerRouter["pet/controllers:ArticlesController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:ArticlesController"] = append(beego.GlobalControllerRouter["pet/controllers:ArticlesController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:ImagesController"] = append(beego.GlobalControllerRouter["pet/controllers:ImagesController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:ImagesController"] = append(beego.GlobalControllerRouter["pet/controllers:ImagesController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:ImagesController"] = append(beego.GlobalControllerRouter["pet/controllers:ImagesController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:ImagesController"] = append(beego.GlobalControllerRouter["pet/controllers:ImagesController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:ImagesController"] = append(beego.GlobalControllerRouter["pet/controllers:ImagesController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:PhotoCommentsController"] = append(beego.GlobalControllerRouter["pet/controllers:PhotoCommentsController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:PhotoCommentsController"] = append(beego.GlobalControllerRouter["pet/controllers:PhotoCommentsController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:PhotoCommentsController"] = append(beego.GlobalControllerRouter["pet/controllers:PhotoCommentsController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:PhotoCommentsController"] = append(beego.GlobalControllerRouter["pet/controllers:PhotoCommentsController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:PhotoCommentsController"] = append(beego.GlobalControllerRouter["pet/controllers:PhotoCommentsController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:PhotoController"] = append(beego.GlobalControllerRouter["pet/controllers:PhotoController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:PhotoController"] = append(beego.GlobalControllerRouter["pet/controllers:PhotoController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:PhotoController"] = append(beego.GlobalControllerRouter["pet/controllers:PhotoController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:PhotoController"] = append(beego.GlobalControllerRouter["pet/controllers:PhotoController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["pet/controllers:PhotoController"] = append(beego.GlobalControllerRouter["pet/controllers:PhotoController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

}
