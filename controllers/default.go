package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Prepare() {
	c.Data["Title"] = "OAS Ruler"
	c.Data["Website"] = "udistrital.edu.co"
	c.Data["Email"] = "apdaza@gmail.com"
	c.Layout = "layout.tpl"

	c.Data["HeadStyles"] = []string{
		"/static/css/bootstrap.min.css",
		"/static/css/bootstrap-theme.min.css",
		"/static/libs/datatables/css/angular-datatables.css",
		"/static/css/app.css",
   }

   c.Data["HeadScripts"] = []string{
		"/static/js/jquery.min.js",
		"/static/js/jquery.dataTables.min.js",
		"/static/js/bootstrap.min.js",
		"/static/js/angular.min.js",
		"/static/libs/datatables/angular-datatables.min.js",
		"/static/js/main.js",
		"/static/js/dominios.js",
		"/static/js/predicados.js",
   }

}

func (c *MainController) Get() {
	c.TplName = "home.html"
}
