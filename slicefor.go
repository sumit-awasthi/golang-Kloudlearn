// Golang program for the iterating over a slice using for loop
package main

import "fmt"

func main() {
	myslice := []string{"This", "is", "the", "practice", "of", "Go", "by Sumit"}
	for e := 0; e < len(myslice); e++ {
		fmt.Println(myslice[e])
	}
}
