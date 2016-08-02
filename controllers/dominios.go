package controllers

type DominiosController struct {
    MainController
}

type AddDominiosController struct {
    MainController
}

type EditDominiosController struct {
    MainController
}

type DeleteDominiosController struct {
    MainController
}

func (c *DominiosController) Get() {
    c.TplName = "listarDominios.html"
}

func (c *AddDominiosController) Get() {
    c.TplName = "agregarDominio.html"
}

func (c *EditDominiosController) Get() {
    c.Data["Id"] = c.Ctx.Input.Param(":id")
    c.TplName = "editarDominio.html"
}

func (c *DeleteDominiosController) Get() {
    c.Data["Id"] = c.Ctx.Input.Param(":id")
    c.TplName = "eliminarDominio.html"
}
