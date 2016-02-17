package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["backendAPI/controllers:ObjectController"] = append(beego.GlobalControllerRouter["backendAPI/controllers:ObjectController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["backendAPI/controllers:ObjectController"] = append(beego.GlobalControllerRouter["backendAPI/controllers:ObjectController"],
		beego.ControllerComments{
			"Get",
			`/:objectId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["backendAPI/controllers:ObjectController"] = append(beego.GlobalControllerRouter["backendAPI/controllers:ObjectController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["backendAPI/controllers:ObjectController"] = append(beego.GlobalControllerRouter["backendAPI/controllers:ObjectController"],
		beego.ControllerComments{
			"Put",
			`/:objectId`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["backendAPI/controllers:ObjectController"] = append(beego.GlobalControllerRouter["backendAPI/controllers:ObjectController"],
		beego.ControllerComments{
			"Delete",
			`/:objectId`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["backendAPI/controllers:TaskController"] = append(beego.GlobalControllerRouter["backendAPI/controllers:TaskController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["backendAPI/controllers:TaskController"] = append(beego.GlobalControllerRouter["backendAPI/controllers:TaskController"],
		beego.ControllerComments{
			"Get",
			`/:taskId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["backendAPI/controllers:TaskController"] = append(beego.GlobalControllerRouter["backendAPI/controllers:TaskController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["backendAPI/controllers:TaskController"] = append(beego.GlobalControllerRouter["backendAPI/controllers:TaskController"],
		beego.ControllerComments{
			"Delete",
			`/:taskId`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["backendAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["backendAPI/controllers:UserController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["backendAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["backendAPI/controllers:UserController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["backendAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["backendAPI/controllers:UserController"],
		beego.ControllerComments{
			"Get",
			`/:uid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["backendAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["backendAPI/controllers:UserController"],
		beego.ControllerComments{
			"Put",
			`/:uid`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["backendAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["backendAPI/controllers:UserController"],
		beego.ControllerComments{
			"Delete",
			`/:uid`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["backendAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["backendAPI/controllers:UserController"],
		beego.ControllerComments{
			"Login",
			`/login`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["backendAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["backendAPI/controllers:UserController"],
		beego.ControllerComments{
			"Logout",
			`/logout`,
			[]string{"get"},
			nil})

}
