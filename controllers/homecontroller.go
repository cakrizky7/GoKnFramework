package controllers

import (
	"github.com/eaciit/knot/knot.v1"
	tk "github.com/eaciit/toolkit"
)

type HomeController struct {
	*Controller
}

func (a *HomeController) Default(k *knot.WebContext) interface{} {
	k.Config.OutputType = knot.OutputTemplate
	if k.Request.Method == "GET" {
		k.Config.ViewName = "home/default.html"
		return tk.M{}.Set("BASE_URL", a.BASE_URL)
	}
	return "Wrong Method"

}
