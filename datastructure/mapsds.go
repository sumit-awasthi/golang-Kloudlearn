//maps in golang
package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"James":    32,
		"Bond":     27,
		"Tommy":    25,
		"Jeniffer": 22,
		"Gaurav":   20,
	}

	fmt.Println(m)
	fmt.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")
	fmt.Println(m["James"]) //fetching values using key values
	v := m["James"]
	fmt.Println(v)
	for k, v := range m { // range keyword for value and key element
		fmt.Println(k, v)
	}

	fmt.Println(m["Sumit"]) // prints 0 as there  is no element namely sumit in the map
	delete(m, "James") // to delete the element from  the map
	fmt.Println("###############################")
	fmt.Println(m)
}
