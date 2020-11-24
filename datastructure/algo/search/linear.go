//Linear search
package main

import (
	"fmt"
)

func LinearSearch(elements []int, findElement int) bool {
	var element int
	for _, element = range elements {
		if element == findElement {
			return true
		}
	}
	return false
}
func main() {
	var elements []int
	elements = []int{15, 48, 26, 18, 41, 86, 29, 51, 20, 45, 49, 55}
	fmt.Println(elements)
	fmt.Println(LinearSearch(elements, 48))
	fmt.Printf("found element %d at index %d in %v\n", element, i, elements)
}
