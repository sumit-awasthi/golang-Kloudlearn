package main

import "fmt"

func main() {
	//var a [5]int
	var a = [5]int{1, 5, 10, 26, 28}
	a[4] = 100
	fmt.Println("get:", a[4])
	fmt.Println("emp:", a)
	fmt.Println("len:", len(a))
}
