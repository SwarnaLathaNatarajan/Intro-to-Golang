package main

import "fmt"

func main() {
	//Define map
	emails := make(map[string]string)
	map2 := map[string]string{"a": "b", "c": "d", "e": "f"}
	fmt.Println(map2)

	//Assign values
	emails["a"] = "b"
	emails["c"] = "d"
	emails["e"] = "f"
	fmt.Println(emails, len(emails), emails["a"])

	//delete
	delete(emails, "e")
	fmt.Println(emails, len(emails), emails["a"])

}
