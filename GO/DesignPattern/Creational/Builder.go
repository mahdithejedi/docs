package main

import "fmt"

type car struct {
	body    string
	wheels  string
	windows string
}

type IBuilder interface {
	setBody(name string) IBuilder
	setWheels(name string) IBuilder
	setWindows(name string) IBuilder
	getCar() car
	Describe() string
}

func (c *car) setBody(name string) IBuilder {
	c.body = name
	return c
}

func (c *car) setWheels(name string) IBuilder {
	c.wheels = name
	return c
}

func (c *car) setWindows(name string) IBuilder {
	c.windows = name
	return c
}

func (c *car) getCar() car {
	return *c
}

func Builder() IBuilder {
	return &car{}
}

func (p *car) Describe() string {
	return fmt.Sprintf("Car body is %s \n, Car wheels are %s \n, also car windows are %s \n",
		p.body, p.wheels, p.windows,
	)
}

func main() {
	// __source__ = 'https://gist.github.com/humamfauzi/57b568d52c26dc21addb56daab070e99'
	fmt.Println(
		Builder().setWheels("Pride Wheels").setBody("Pride Body").setWindows("Pride Windows").Describe(),
	)
}
