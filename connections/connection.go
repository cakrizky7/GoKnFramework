package connection

import (
	. "GOKNFRAMEWORK/helpers"

	"github.com/eaciit/dbox"
	_ "github.com/eaciit/dbox/dbc/mongo"
	tk "github.com/eaciit/toolkit"
)

func GetConnection() (dbox.IConnection, error) {
	cfg := ReadAppConfig()
	ci := dbox.ConnectionInfo{
		tk.ToString(cfg["DB_ADDRESS"]),
		tk.ToString(cfg["DB"]),
		tk.ToString(cfg["DB_USER"]),
		tk.ToString(cfg["DB_PASSWORD"]),
		nil,
	}
	conn, err := dbox.NewConnection("mongo", &ci)
	if err != nil {
		panic("Connect Failed") // Change with your error handling
	}
	err = conn.Connect()
	if err != nil {
		panic("Connect Failed") // Change with your error handling
	}
	if err != nil {
		return nil, err
	}
	return conn, nil
}
