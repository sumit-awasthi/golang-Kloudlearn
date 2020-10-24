package main
import"fmt"
func main(){
    var a int
    var ptr *int 
    var pptr **int
    a=100
    ptr=&a
    pptr=&ptr
    fmt.Printf("the value of a is  %d\n",a)
    fmt.Printf("the value of a is  %d\n",*ptr)
    fmt.Printf("the value of a is  %d\n",**pptr)



}