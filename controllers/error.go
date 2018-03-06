package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error403() {
	c.Data["content"] = "403 Forbidden"
	c.TplName = "error/error.html"
}

func (c *ErrorController) Error404() {
	c.Data["content"] = "404, page not found"
	c.TplName = "error/error.html"
}

func (c *ErrorController) Error500() {
	c.Data["content"] = "500, internal server error"
	c.TplName = "error/error.html"
}

func (c *ErrorController) ErrorDb() {
	c.Data["content"] = "database is now down"
	c.TplName = "error/error.html"
}