//bubble sort
package main

import (
	"fmt"
)

var data []int = []int{1, 3, 2, 4, 8, 6, 7, 2, 3, 0}

func BubbleSort(data []int) {
	for i := 0; i < len(data); i++ {
		for j := 1; j < len(data)-i; j++ {
			if data[j] < data[j-1] {
				data[j], data[j-1] = data[j-1], data[j]
			}
		}
	}
}
func main() {
	fmt.Println("Before sort\n", data)
	fmt.Println("***********************************")
	fmt.Printf("After sort\n")
	BubbleSort(data)
	fmt.Println(data)

}
