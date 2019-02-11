package controllers

import (
	"github.com/eaciit/knot/knot.v1"
)

type HomeController struct {
	*Controller
}

func (a *HomeController) Default(k *knot.WebContext) interface{} {
	k.Config.OutputType = knot.OutputTemplate
	if k.Request.Method == "GET" {
		k.Config.ViewName = "home/default.html"
		return ""
	}
	return "Wrong Method"

}
