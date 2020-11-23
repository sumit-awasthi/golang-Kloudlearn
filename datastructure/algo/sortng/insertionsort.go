//Insertion sort
package main

import (
	"fmt"
)

var data []int = []int{1, 3, 2, 4, 8, 6, 7, 2, 3, 0}

func InsertionSort(data []int) {
	for i := 1; i < len(data); i++ {
		if data[i] < data[i-1] {
			j := i - 1
			temp := data[i]
			for j >= 0 && data[j] > temp {
				data[j+1] = data[j]
				j--
			}
			data[j+1] = temp
		}
	}
}
func main() {
	fmt.Println("Before sort:-\t", data)
	fmt.Println("***********************************")
	fmt.Printf("After sort:-")
	InsertionSort(data)
	fmt.Println("\t", data)

}
