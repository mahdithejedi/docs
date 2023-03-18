package listener

import (
	"Notifier/helpers"
	"database/sql"
	"fmt"
)

func Run(db *sql.DB) {
	fmt.Println("Going to run")
	row, err := db.Query("SELECT NOW();")
	helpers.PureCheckErr(err)
	var res string
	row.Next()
	row.Scan(&res)
	fmt.Println(res)
}
