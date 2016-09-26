package routers

import (
	"github.com/zhaoyao/showtime/backend/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
