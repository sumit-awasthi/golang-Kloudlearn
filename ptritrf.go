// Golang program to demonstrate the use of type inference in Pointer variables
package main

import "fmt"

func main() {

	var y = 458
	var p = &y

	fmt.Println("Value stored in y = ", y)
	fmt.Println("Address of y = ", &y)
	fmt.Println("Value stored in pointer variable p = ", p)
}
