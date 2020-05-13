package main

import "fmt"

func main() {

	//If else
	x, y := 15, 10
	if x < y {
		fmt.Printf("%d is < %d\n", x, y)
	} else {
		fmt.Println("Meh")
	}

	//else if
	color := "red"
	if color == "red" {
		fmt.Println("red")
	} else if color == "blue" {
		fmt.Println("blue")
	}

	//Switch
	switch color {
	case "red":
		fmt.Println("RRRed")
	case "blue":
		fmt.Println("BBBlue")
	default:
		fmt.Println("Meh")
	}
}
