package main

import (
	"fmt"
	"math"

	"github.com/SwarnaLathaNatarajan/Intro-to-Golang/packages/strutil"
)

func main() {

	fmt.Println("Hello World")

	var name = "Swarna"
	fmt.Println(name)

	var age = 22
	fmt.Println(age)

	fmt.Printf("%T", age)

	var isCool = true
	fmt.Printf("%T\n", isCool)

	//shorthand

	size := 12
	fmt.Println(size)

	name, email := "Swarna", "mail"
	fmt.Println(name, email)

	//Math Package

	fmt.Println(math.Floor(2.7), math.Ceil(2.7))

	//Calling function

	fmt.Println(strutil.Reverse("Kawshik"))
}
