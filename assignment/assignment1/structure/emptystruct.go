//empty struct
package main

import (
	"fmt"
)

type Employee struct {
	firstName string
	lastName  string
	age       int
	salary    int
}

func main() {
	var emp Employee //zero valued struct
	fmt.Println("First Name:", emp.firstName)
	fmt.Println("Last Name:", emp.lastName)
	fmt.Println("Age:", emp.age)
	fmt.Println("Salary:", emp.salary)
}
