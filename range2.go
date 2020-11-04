package main

import "fmt"

func main() {
	var a []int = []int{1, 2, 3, 4, 5, 6, 7, 8}

	for _, element := range a {
		for _, element2 := range a {
			if element == element2 {
				fmt.Println(element)
			}
		}
	}
}
