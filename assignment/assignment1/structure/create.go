//creating a struct
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

    //creating struct specifying field names
    emp1 := Employee{
        firstName: "Sumit",
        age:       21,
        salary:    11000,
        lastName:  "Awasthi",
    }

    //creating struct without specifying field names
    emp2 := Employee{"Tonny", "Stark", 29, 10800}

    fmt.Println("Employee 1", emp1)
    fmt.Println("Employee 2", emp2)
}