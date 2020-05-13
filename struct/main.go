package main

import (
	"fmt"
	"strconv"
)

//Define Person struct
type Person struct {
	name         string
	age          int
	gender, city string
}

//value reciever method
func (p Person) greet() string {
	return "Hello " + p.name + " " + strconv.Itoa(p.age)
}

//pointer reciever method
func (p *Person) hasBirthday() {
	p.age++
}

func main() {

	//Initialize person using struct
	p1 := Person{name: "swarna", age: 22}
	p2 := Person{"latha", 25, "F", "Boulder"}
	fmt.Println(p1, p2)
	fmt.Println(p1.name)
	fmt.Println(p1.greet())
	p2.hasBirthday()
	fmt.Println(p2.greet())
}
