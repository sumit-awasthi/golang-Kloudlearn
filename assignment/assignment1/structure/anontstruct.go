//anonymous struct field
package main

import (
	"fmt"
)

type Address struct {
	city  string
	state string
}
type Person struct {
	name string
	age  int
	Address
}

func main() {
	p := Person{
		name: "Sumit",
		age:  21,
		Address: Address{
			city:  "Lucknow",
			state: "Uttar Pradesh",
		},
	}

	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("City:", p.city)   //city is promoted field
	fmt.Println("State:", p.state) //state is promoted field
}
