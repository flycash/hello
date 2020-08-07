package routers

import (
	"github.com/astaxie/beego/pkg"
	"github.com/flycash/hello/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
