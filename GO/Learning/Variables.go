package main

import (
	"fmt"
	"math"
)

func main() {
	// you can assign a variable without type declartion
	var a = 2
	fmt.Println(a)

	// or with type declartion
	var b int = 32
	fmt.Println(b)

	var e uint
	e = 32
	fmt.Println(e)

	// or you can declare and assign a variable just by :=
	c := "this is a test"
	fmt.Println(c)

	sin_b := math.Sin(float64(b))

	fmt.Println(sin_b)

	// you can variable as a  constant
	const test_const = "test"

}
