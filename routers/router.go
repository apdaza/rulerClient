package routers

import (
	"rulerClient/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
		beego.Router("/listarDominios", &controllers.DominiosController{})
		beego.Router("/agregarDominios", &controllers.AddDominiosController{})
		beego.Router("/editarDominio/:id:int", &controllers.EditDominiosController{})
        beego.Router("/eliminarDominio/:id:int", &controllers.DeleteDominiosController{})
		beego.Router("/listarPredicados", &controllers.PredicadosController{})
		beego.Router("/agregarPredicados", &controllers.AddPredicadosController{})
		beego.Router("/editarPredicado/:id:int", &controllers.EditPredicadosController{})
        beego.Router("/eliminarPredicado/:id:int", &controllers.DeletePredicadosController{})
}
