//Selection  sort
package main

import (
	"fmt"
)

var data []int = []int{1, 3, 2, 4, 8, 6, 7, 2, 3, 0}

func SelectionSort(data []int) {
	length := len(data)
	for i := 0; i < length; i++ {
		maxIndex := 0
		for j := 1; j < length-i; j++ {
			if data[j] > data[maxIndex] {
				maxIndex = j
			}
		}
		data[length-i-1], data[maxIndex] = data[maxIndex], data[length-i-1]
	}
}
func main() {
	fmt.Println("Before sort:-\t",data)
	fmt.Println("***********************************")
	fmt.Printf("After sort:-")
	SelectionSort(data)
	fmt.Println("\t",data)
	
}
