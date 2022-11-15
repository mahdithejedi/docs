package helpers

import (
	"database/sql"
	"strings"
)

var max_try int = 30
var current_try int

func Select(conn *sql.DB) (*sql.Rows, error) {
	return conn.Query("select name, unit from markets_coin")
}

func Insert(conn *sql.DB) (sql.Result, error) {
	coin := Coin()
	res, err := conn.Exec(
		"insert into markets_coin(created_at, updated_at, name, unit) values ($1, $2, $3, $4)",
		coin.created_at, coin.updated_at, coin.name, coin.unit)
	if err != nil && strings.Contains(err.Error(), "value too long for type character varying(20)") {
		if current_try > max_try {
			current_try = 0
			return nil, err
		}
		return Insert(conn)
	}
	current_try = 0
	return res, err
}
