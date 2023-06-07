package runners

import (
	"database/sql"
	"hash/fnv"
	"time"
)

type DBError struct {
	obj     error
	reason  string
	hash    uint32
	counter int
}

func (e *DBError) GetForRedis() (uint32, string, int) {
	return e.hash, e.reason, e.counter
}

var Errors map[string]int

type query func(*sql.DB) (sql.Result, error)

func InsertQueryRunner(connection *sql.DB, query query, running *bool, errors chan<- error, sleep int, runnerCount int) {
	for i := 0; i < runnerCount; i++ {
		//go runner(connection, query, running, errors, sleep)
		runner(connection, query, running, errors, sleep)
	}
}

func runner(connection *sql.DB, query query, running *bool, errors chan<- error, sleep int) {
	for *running {
		baseRunner(connection, query, errors)
		time.Sleep(time.Millisecond * time.Duration(sleep))
	}
}

func baseRunner(connection *sql.DB, query query, errors chan<- error) {
	_, err := query(connection)
	if err != nil {
		errors <- err
	}
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func CaptureErrors(errors map[uint32]DBError, error error) DBError {
	hashedError := hash(error.Error())
	if value, ok := errors[hashedError]; ok {
		updateError(&value)
		errors[hashedError] = value
	} else {
		errors[hashedError] = DBError{
			hash:    hashedError,
			reason:  error.Error(),
			obj:     error,
			counter: 1}
	}
	return errors[hashedError]
}

func updateError(er *DBError) {
	er.counter = er.counter + 1
}
