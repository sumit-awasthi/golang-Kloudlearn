//equal or not
package main
import (  
    "fmt"
)

type name struct {  
    firstName string
    lastName  string
}

func main() {  
    name1 := name{
        firstName: "Sumit",
        lastName:  "Awasthi",
    }
    name2 := name{
        firstName: "Sumit",
        lastName:  "Awasthi",
    }
    if name1 == name2 {
        fmt.Println("name1 and name2 are equal")
    } else {
        fmt.Println("name1 and name2 are not equal")
    }

    name3 := name{
        firstName: "Sumit",
        lastName:  "Awasthi",
    }
    name4 := name{
        firstName: "Sumit",
    }

    if name3 == name4 {
        fmt.Println("name3 and name4 are equal")
    } else {
        fmt.Println("name3 and name4 are not equal")
    }
}