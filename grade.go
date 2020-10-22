package main

import "fmt"

func main() {
	grade := 'A'
	fmt.Printf("grade %d is ", grade)
	switch grade {
	case 'A':
		fmt.Println("Good")
	case 'B':
		fmt.Println("Ok")
	case 'C':
		fmt.Println("Average")
	case 'D':
		fmt.Println("Poor")
	case 'E':
		fmt.Println("Very Poor")
	case 'F':
		fmt.Println("Fail")
	default:
		fmt.Println("Result not out yet")

	}
}

