package main

import "fmt"

func Calculate(input int) (result int) {
	result = input + 2
	return result
}

func main() {
	fmt.Println("Hello World")
	result := Calculate(2)
	fmt.Println(result)
}
