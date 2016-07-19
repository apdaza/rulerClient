package routers

import (
	"rulerClient/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
		beego.Router("/listarDominios", &controllers.DominiosController{})
		beego.Router("/agregarDominios", &controllers.AddDominiosController{})
		beego.Router("/listarPredicados", &controllers.PredicadosController{})
		beego.Router("/agregarPredicados", &controllers.AddPredicadosController{})
}
