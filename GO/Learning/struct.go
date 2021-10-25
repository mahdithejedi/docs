package main

import (
	"fmt"
	"time"
)

type person struct {
	name      string
	family    string
	age       int
	birthyear int
}

func new_person(name string, family string, age int) *person {
	p := person{name: name, family: family, age: age}
	return &p
}

// you can add method to a struct
func (p *person) set_birthyear() {
	current_time := time.Now()
	p.birthyear = current_time.Year() - p.age
}

func main() {
	mahdi := person{name: "Mahdi", family: "Moosavi", age: 22}
	fmt.Println("mahdi is ", mahdi)

	new_mahdi := &mahdi
	fmt.Println("new_mahdi is ", new_mahdi, "and *new_mahdi is ", *new_mahdi, "&new_mahdi", &new_mahdi)

	new_mahdi_moosavi := new_person("mahdi", "moosavi", 12)
	new_mahdi_moosavi.age = 22
	new_mahdi_moosavi.set_birthyear()
	fmt.Println("new_mahdi_moosavi is ", new_mahdi_moosavi, "and *new_mahdi_moosavi is ", *new_mahdi_moosavi, "&new_mahdi_moosavi", &new_mahdi_moosavi)

}
