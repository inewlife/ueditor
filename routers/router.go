package routers

import (
	"zueditor/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/ueditor", &controllers.UeditorController{})
	beego.Router("/ueditor/action", &controllers.UeditorController{}, "*:Action")
}
