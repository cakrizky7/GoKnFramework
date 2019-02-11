package controllers

import (
	"github.com/eaciit/dbox"
)

type Controller struct {
	APP_NAME string
	BASE_URL string
	DB       dbox.IConnection
}
