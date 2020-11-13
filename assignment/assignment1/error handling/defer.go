//defer in golang
package main

import "fmt"

func add(a1, a2 int) int {
	res := a1 + a2
	fmt.Println("Result: ", res)
	return 0
}
func main() {

	fmt.Println("Start")
	defer fmt.Println("End")
	defer add(34, 56)
	defer add(10, 10)
}
