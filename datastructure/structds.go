// struct in golang

package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}
type persontwo struct {
	person
	feelwell bool
}

func main() {
	per2 := persontwo{
		person: person{
			first: "James",
			last:  "Bond",
			age:   54,
		},
		feelwell: true,
	}
	fmt.Println(per2)
	p1 := person{
		first: "Harry",
		last:  "Potter",
		age:   25,
	}
	p2 := person{
		first: "Tonny",
		last:  "Stark",
		age:   20,
	}
	fmt.Println(p1)
	fmt.Println(p2.first)

}
