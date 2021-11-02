package main

import "fmt"

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

}
