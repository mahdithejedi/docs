package main

import (
	"fmt"
	"reflect"
)

func Incrs(a *int) {
	fmt.Print("****\n")
	fmt.Print(a)
	fmt.Print(reflect.TypeOf(a))
	fmt.Print("\n=======\n")
	fmt.Print(*a)
	fmt.Print(reflect.TypeOf(*a))
	*a++
}

func main() {
	// * Operator also termed as the dereferencing operator used to declare pointer variable and access the value stored in the address.
	// & operator termed as address operator used to returns the address of a variable or to access the address of a variable to a pointer.

	// Declare a pointer
	// var pointer_NAME *Data_TYPE

	var s *string // store only the memory addresses of string variables.
	s_variable := "hi this world is changin"
	s = &s_variable
	fmt.Println(*s)

	int_test := 12
	fmt.Println(&int_test)
	Incrs(&int_test)
	fmt.Printf("\n After Increase %v", int_test)

}
