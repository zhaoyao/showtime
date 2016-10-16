package routers

import (
	"github.com/astaxie/beego"
	"github.com/zhaoyao/showtime/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/search", &controllers.SearchController{})
	beego.Router("/resources/:id", &controllers.ResourceController{})
}
