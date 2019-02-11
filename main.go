package main

import (
	. "GOKNFRAMEWORK/connections"
	. "GOKNFRAMEWORK/controllers"
	. "GOKNFRAMEWORK/helpers"
	"log"
	"net/http"
	"os"

	_ "github.com/eaciit/dbox/dbc/mongo"
	"github.com/eaciit/knot/knot.v1"
)

var (
	wd = func() string {
		d, _ := os.Getwd()
		return d + "\\"
	}()
)

func main() {
	app := knot.GetApp("ext")
	if app == nil {
		return
	}
	routes := make(map[string]knot.FnContent, 1)
	routes["/"] = func(r *knot.WebContext) interface{} {
		http.Redirect(r.Writer, r.Request, "/home/default", http.StatusTemporaryRedirect)
		return true
	}

	knot.StartAppWithFn(app, "localhost:12345", routes)
}

func init() {
	cfg := ReadAppConfig()
	conn, err := GetConnection()
	if err != nil {
		log.Println(err)
	}
	controller := new(Controller)
	controller.APP_NAME = cfg["APP_NAME"]
	controller.BASE_URL = cfg["BASE_URL"]
	controller.DB = conn

	app := knot.NewApp("ext")
	app.ViewsPath = wd + "/views/"

	//Controller + AutoRouting
	app.Register(&HomeController{controller})

	app.Static("static", wd+"assets")
	app.LayoutTemplate = "layout.html"
	knot.RegisterApp(app)
}
