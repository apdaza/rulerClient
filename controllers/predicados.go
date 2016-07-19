package controllers

type PredicadosController struct {
    MainController
}

type AddPredicadosController struct {
    MainController
}

func (c *PredicadosController) Get() {
    c.TplName = "listarPredicados.html"
}

func (c *AddPredicadosController) Get() {
    c.TplName = "agregarPredicado.html"
}
