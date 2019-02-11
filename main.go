package main

import (
	. "TUTORIAL/controllers"
	"net/http"
	"os"

	"github.com/eaciit/dbox"
	_ "github.com/eaciit/dbox/dbc/mongo"
	"github.com/eaciit/knot/knot.v1"
	tk "github.com/eaciit/toolkit"
)

var (
	wd = func() string {
		d, _ := os.Getwd()
		return d + "\\"
	}()
)
var APP_NAME = "TUTORIAL"
var BASE_URL = "localhost"

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
	ci := dbox.ConnectionInfo{
		"127.0.0.1",
		"tutorial",
		"",
		"",
		nil,
	}
	conn, err := dbox.NewConnection("mongo", &ci)
	if err != nil {
		tk.Println(err)
		panic("Connect Failed") // Change with your error handling
	}
	err = conn.Connect()
	if err != nil {
		panic("Connect Failed") // Change with your error handling
	}

	controller := new(Controller)
	controller.APP_NAME = APP_NAME
	controller.BASE_URL = BASE_URL
	controller.DB = conn

	app := knot.NewApp("ext")
	app.ViewsPath = wd + "/views/"
	app.Register(&HomeController{controller})
	app.Static("static", wd+"assets")
	app.LayoutTemplate = "layout.html"
	knot.RegisterApp(app)
}
