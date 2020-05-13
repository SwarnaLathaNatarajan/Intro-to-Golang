package main

import "fmt"

func main() {

	//for loop - long method

	i := 1
	for i < 10 {
		fmt.Println(i)
		// i = i + 1
		i++
	}

	//for loop - short method

	for i := 1; i < 10; i++ {
		fmt.Println(i % 2)
	}

}
