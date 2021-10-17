package main

import (
	"fmt"
	"math"
	"reflect"
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

	// ********
	// ARRAYS
	// ********
	var array_a [5]int
	array_a[4] = 3
	fmt.Println('\n', a)
	// or you can defind one line variable
	array_b := [5]int{1, 2, 3, 4, 5}
	fmt.Println('\n', array_b[0])

	// ********
	// SLICES
	// ********
	// slices is array but more like python way
	slice_a := make([]string, 3)
	slice_a[0] = "a"
	slice_a[1] = "b"
	slice_a[2] = "c"
	slice_a = append(slice_a, "d")
	slice_a = append(slice_a, "e", "f")

	slice_b := make([]string, len(slice_a))
	copy(slice_b, slice_a)
	fmt.Println("\n\n slice b", slice_b)
	fmt.Println(slice_b[:3], slice_a[1:3])

	// fill slice in one line
	slice_c := []string{"a", "b", "c"}
	slice_c = append(slice_c, "END!")
	fmt.Println("\n\n\n slice c", slice_c)
	slice_d := append(slice_c, "==-")
	fmt.Println("\n\n\n slice d", slice_d)

	for index, slice_value := range slice_d {
		fmt.Println("\n", "index is ", index, " value is ", slice_value)
	}

	// ********
	// MAPS
	// ********
	// you can make maps with `make`
	map_a := make(map[string]float32)
	map_a["Test1"] = 87.6
	map_a["Test2"] = 980.78
	delete(map_a, "Test2")

	var_exists, _ := map_a["Test3"]
	fmt.Println("\n var 3 existance", var_exists, reflect.TypeOf(var_exists))

	if var_exists == float32(0) {
		fmt.Println("not exists!")
	}

	// declare map in one line!
	map_b := map[int]string{1: "Tester", 2: "Testing"}
	fmt.Println("\n\n\n", map_b)

	for key, value := range map_b {
		fmt.Println("\n", "key is ", key, "value is ", value)
	}
	// or you can just loop over keys
	for kys := range map_b {
		fmt.Println("\n", "kys is ", kys)
	}

}
