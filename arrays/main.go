package main

import "fmt"

func main() {
	var fruits [2]string
	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fmt.Println(fruits)

	names := [2]string{"swarna", "latha"}
	fmt.Println(names)

	slices := []string{"1", "2"}
	fmt.Println(slices, len(slices), slices[1:])
}
