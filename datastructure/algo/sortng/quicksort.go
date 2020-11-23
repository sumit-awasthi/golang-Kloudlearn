//Quick sort using recursion
package main

import (
	"fmt"
)

var data []int = []int{1, 3, 2, 4, 8, 6, 7, 2, 3, 0}

func quickSort(nums []int) {
	recursionSort(nums, 0, len(nums)-1)
}

func recursionSort(data []int, left int, right int) {
	if left < right {
		pivot := partition(data, left, right)
		recursionSort(data, left, pivot-1)
		recursionSort(data, pivot+1, right)
	}
}

func partition(data []int, left int, right int) int {
	for left < right {
		for left < right && data[left] <= data[right] {
			right--
		}
		if left < right {
			data[left], data[right] = data[right], data[left]
			left++
		}

		for left < right && data[left] <= data[right] {
			left++
		}
		if left < right {
			data[left], data[right] = data[right], data[left]
			right--
		}
	}
	return left
}
func main() {
	fmt.Println("Before sort:-\t", data)
	fmt.Println("***********************************")
	fmt.Printf("After sort:-")
	quickSort(data)
	fmt.Println("\t", data)

}
