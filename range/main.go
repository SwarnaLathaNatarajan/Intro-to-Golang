package main

import "fmt"

func main() {
	ids := []int{1, 2, 3, 4, 5}

	for i, id := range ids {
		fmt.Println(i, id)
	}

	for id := range ids {
		fmt.Println(id)
	}

	map1 := map[string]string{"a": "b", "c": "d"}
	for k, v := range map1 {
		fmt.Println(k, v)
	}
	for k := range map1 {
		fmt.Println(k)
	}
}
