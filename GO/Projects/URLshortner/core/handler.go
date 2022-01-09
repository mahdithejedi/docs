package core

import (
	"fmt"
	"net"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func remote_dsn(name string, password string, host net.IP, port int) string {
	return fmt.Sprintf("%s:%s@tcp(%s,%d)", name, password, host, port)
}

func initiate_DB(name string, password string, host string, port int) *gorm.DB {
	fmt.Print("Go to initiate")
	dsn := remote_dsn(name, password, net.ParseIP(host), port)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("db err (Init) err:%s", err)
	}
	DB = db
	return DB
}
