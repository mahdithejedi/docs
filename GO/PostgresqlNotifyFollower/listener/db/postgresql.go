package db

import (
	"Notifier/helpers"
	"Notifier/listener"
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"strconv"
)

var psqlConnection *sql.DB

func Connect(config *listener.Config) *pq.Listener {
	port, err := strconv.Atoi(config.Port)
	helpers.PureCheckErr(err)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, port, config.Username, config.Password, config.DBName)
	checkConnect(psqlInfo)
	return subscribe(config, psqlInfo)
}

func checkConnect(psqlInfo string) {
	psqlConnection, err := sql.Open("postgres", psqlInfo)
	helpers.CheckErr("Error while connecting postgresql DB", err)
	err = psqlConnection.Ping()
	helpers.CheckErr("Error while Ping DB", err)
	defer psqlConnection.Close()
}

func Disconnect() {
	helpers.CheckErr("Error while disconnecting DB", psqlConnection.Close())
}

func subscribe(config *listener.Config, psql string) *pq.Listener {
	return pq.NewListener(psql, config.MinReconnect, config.MaxReconnect, func(event pq.ListenerEventType, err error) {
		helpers.CheckErr("Error in listener", err)
	})
}
