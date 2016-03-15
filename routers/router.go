package routers

import (
	"WebScan/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/img/verifypic", &controllers.VerifyCodeController{})
}
