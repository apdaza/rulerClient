package controllers

type PredicadosController struct {
    MainController
}

type AddPredicadosController struct {
    MainController
}

type EditPredicadosController struct {
    MainController
}

type DeletePredicadosController struct {
    MainController
}

func (c *PredicadosController) Get() {
    c.TplName = "listarPredicados.html"
}

func (c *AddPredicadosController) Get() {
    c.TplName = "agregarPredicado.html"
}

func (c *EditPredicadosController) Get() {
    c.Data["Id"] = c.Ctx.Input.Param(":id")
    c.TplName = "editarPredicado.html"
}

func (c *DeletePredicadosController) Get() {
    c.Data["Id"] = c.Ctx.Input.Param(":id")
    c.TplName = "eliminarPredicado.html"
}
