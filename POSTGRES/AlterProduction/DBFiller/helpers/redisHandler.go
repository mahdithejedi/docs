package helpers

import (
	"Filler/runners"
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"strconv"
)

var rdb *redis.Client
var ctx = context.Background()

func init() {
	_rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       2,  // use default DB
	})
	rdb = _rdb
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err.Error())
		panic(err.Error())
	}
}

func Set(errDB runners.DBError) {
	hash, reason, counter := errDB.GetForRedis()
	err := rdb.HMSet(ctx, strconv.FormatUint(uint64(hash), 10), reason, counter)
	checkError(err.Err())
}
