package DB

import (
	"context"
	"fmt"
	_ "github.com/lib/pq"
	"gopkg.in/ini.v1"
	"swapcore/Core/DB/ent"
	"swapcore/Core/lib"
)

var Client *ent.Client

func getStringConnection() string {
	cfg, err := ini.Load("./config.ini")
	lib.CheckError(err)
	config := cfg.Section("postgres")
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		config.Key("Host").String(),
		config.Key("Port").String(),
		config.Key("Username").String(),
		config.Key("DB_Name").String(),
		config.Key("Password").String(),
	)
}

func init() {
	print("BEFORE")
	var err error
	Client, err = ent.Open("postgres", getStringConnection(), func(config *interface{}) {
		config.debug = true
	})
	print("AFTER")
	lib.CheckError(err)
	lib.CheckError(Client.Schema.Create(context.Background()))
}
