package main

import (
	"URLShortner/core"
	"fmt"
)

func main() {
	db := core.Initiate_DB("go_learn", "go", "172.16.4.2", 5432)
	fmt.Print(db)
}
