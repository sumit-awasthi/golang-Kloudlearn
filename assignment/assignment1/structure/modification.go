//modification in struct
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
    emp := Employee{
        firstName: "Sumit",
        lastName:  "Awasthi",
        age:       21,
        salary:    11000,
    }
    fmt.Println("First Name:", emp.firstName)
    fmt.Println("Last Name:", emp.lastName)
    fmt.Println("Age:", emp.age)
    fmt.Printf("Salary: $%d\n", emp.salary)
    emp.salary = 16500
    fmt.Printf("New Salary: $%d", emp.salary)
}