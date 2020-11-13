package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {

	p := person{name: name}
	p.age = 26
	return &p
}

func main() {

	fmt.Println(person{"Sumit", 21})

	fmt.Println(& person{name: "Suraj", age: 25})

	fmt.Println(newPerson("Dinesh"))

	s := person{name: "Sanna", age: 23}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 21
	fmt.Println(sp.age)
}