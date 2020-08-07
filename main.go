package main

import (
	"github.com/astaxie/beego/pkg"

	"github.com/flycash/hello/controllers"
	_ "github.com/flycash/hello/routers"
)

func main() {
	// beego.InsertFilter("*", beego.BeforeRouter, filter.CORS())
	beego.Include(&controllers.CMSController{})
	// admin := beego.NewNamespace("/Admin", beego.NSInclude(&controllers.AdminController{}))
	// beego.AddNamespace(admin)
	beego.Run()
}
