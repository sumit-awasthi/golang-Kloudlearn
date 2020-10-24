package main
import"fmt"
func main(){
    var a int=20
    var ip*int
    ip=&a
    fmt.Printf("address of a variableb %x\n",&a)
    fmt.Printf("address stored in ip variableb %x\n",ip)
    fmt.Printf("value stored in ip variableb %d\n",*ip)

}