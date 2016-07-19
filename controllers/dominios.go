package controllers

type DominiosController struct {
    MainController
}

type AddDominiosController struct {
    MainController
}

func (c *DominiosController) Get() {
    c.TplName = "listarDominios.html"
}

func (c *AddDominiosController) Get() {
    c.TplName = "agregarDominio.html"
}
