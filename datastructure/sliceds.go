//Slices in golang

package main

import (
	"fmt"
)

func main() {
	su := []string{"Sumit", "awasthi", "@8990"} //slice1
	fmt.Println(su)
	aw := []string{"Harry", "Potter", "@87870"} // slice2
	fmt.Println(aw)

	ap := [][]string{su, aw} // append two slices(slice1 +slice2)
	fmt.Println(ap)
	fmt.Println("****************************************")
	xi := []int{4, 5, 6, 4, 8, 7, 123}
	for i, v := range xi { // slice to display the index and values of the element of the slice
		fmt.Println(i, v)
		fmt.Println("######################################")
	}
	xi = append(xi, 120, 84, 78, 46, 54)
	fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")
	fmt.Println("The values appended to the slice :-", xi)
}
