package main

import (
	"fmt"
	"time"
)

func display(str string, until int) {
	for w := 0; w < until; w++ {
		fmt.Println(str)
		time.Sleep(5)
	}
}

func main() {
	go display("inside Gorouting!", 7)
	display("Outside", 7)

}
