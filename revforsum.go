package main

import "fmt"

func main() {

    sum := 0
    i := 9
    //reverse sum from 9 to 1
    for i > 0 {
        
        sum += i
        i--
    }

    fmt.Println(sum)
}