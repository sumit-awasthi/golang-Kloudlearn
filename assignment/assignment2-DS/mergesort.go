//merge sort using go routing

package main

import (
	"fmt"
)

func main() {
	s := []int{11, 25, 9, 8, 10, 12, 18, 14, 7, 15}

	result := make(chan []int)
	go MergeSort(s, result)

	r := <-result
	for _, v := range r {
		fmt.Println("the elements are:", v)
	}
	close(result)
}
func Merge(leftdata []int, rightdata []int) (result []int) {
	result = make([]int, len(leftdata)+len(rightdata))
	lidx, ridx := 0, 0

	for i := 0; i < cap(result); i++ {
		switch {
		case lidx >= len(leftdata):
			result[i] = rightdata[ridx]
			ridx++
		case ridx >= len(rightdata):
			result[i] = leftdata[lidx]
			lidx++
		case leftdata[lidx] < rightdata[ridx]:
			result[i] = leftdata[lidx]
			lidx++
		default:
			result[i] = rightdata[ridx]
			ridx++
		}
	}

	return
}

func MergeSort(data []int, r chan []int) {
	if len(data) == 1 {
		r <- data
		return
	}

	leftChan := make(chan []int)
	rightChan := make(chan []int)
	middle := len(data) / 2

	go MergeSort(data[:middle], leftChan)
	go MergeSort(data[middle:], rightChan)

	leftdata := <-leftChan
	rightdata := <-rightChan

	close(leftChan)
	close(rightChan)
	r <- Merge(leftdata, rightdata)
	return
}
