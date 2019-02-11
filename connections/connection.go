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
		cfg["DB_ADDRESS"],
		cfg["DB"],
		cfg["DB_USER"],
		cfg["DB_PASSWORD"],
		nil,
	}
	conn, err := dbox.NewConnection(cfg["DB_DRIVER"], &ci)
	if err != nil {
		tk.Println(err)
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
