// array data structure
package main

import (
	"fmt"
)

func main() {
	x := [7]int{1, 2, 5, 3, 4, 68, 4}
	fmt.Println("the value stored in the array:", x)
	y := [4]string{"Red", "Green", "Blue", "Orange"}
	fmt.Println("the value stored in array of string :", y)
	
	
	for k, v := range x {
		fmt.Println("the index & value is", k, v)
	}
}
