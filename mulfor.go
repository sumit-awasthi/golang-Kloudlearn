package main

import "fmt"

func main() { //Dec multiple variable in for
	for i, j := 1, 1; i <= 10 && j <= 10; i, j = i+1, j+1 {
		fmt.Println(i, "*", j, "=", i*j)
	}
}
