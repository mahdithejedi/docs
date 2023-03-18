package db

import (
	"Notifier/helpers"
	"Notifier/listener"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"strconv"
)

var psqlConnection *sql.DB

func Connect(config *listener.Config) *sql.DB {
	port, err := strconv.Atoi(config.Port)
	helpers.PureCheckErr(err)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, port, config.Username, config.Password, config.DBName)

	psqlConnection, err := sql.Open("postgres", psqlInfo)
	helpers.CheckErr("Error while connecting postgresql DB", err)
	err = psqlConnection.Ping()
	helpers.CheckErr("Error while Ping DB", err)
	return psqlConnection
}

func Disconnect() {
	helpers.CheckErr("Error while disconnecting DB", psqlConnection.Close())
}
