package main

import (
	"Filler/helpers"
	"Filler/runners"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var connection *sql.DB

func checkErrorWithMsg(err error, msg string) {
	if err != nil {
		log.Fatal(fmt.Sprintf("Error %s, reason %s", msg, err.Error()))
		//panic(err)
	}
}

func CheckError(err error) {
	if err != nil {
		checkErrorWithMsg(err, "")
	}
}

func init() {
	connStr := "host=localhost user=bot password=bot dbname=bot"
	_connection, err := sql.Open("postgres", connStr)
	CheckError(err)
	connection = _connection
}

func main() {
	errStream := make(chan error)
	ErrorMap := make(map[uint32]runners.DBError)
	running := true
	go func() {
		var err error
		for true {
			err = <-errStream
			helpers.Set(runners.CaptureErrors(ErrorMap, err))
			fmt.Println("error is", err.Error())
		}
	}()
	runners.InsertQueryRunner(connection, helpers.Insert, &running, errStream, 100, 10)

}
