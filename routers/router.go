package routers

import (
	"github.com/GuoLizhi/read-community/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
