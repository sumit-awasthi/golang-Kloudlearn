package main

import "fmt"

func main() {
	var x, a, b, c, d = 3, 4, "sumit", 46, 15.15
	fmt.Println(x)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	//with type and value
	fmt.Printf("the value and data type is %v, %T\n", x, x)
	fmt.Printf("the value and data type is %v, %T\n", a, a)
	fmt.Printf("the value and data type is %v, %T\n", b, b)
	fmt.Printf("the value and data type is %v, %T\n", c, c)
	fmt.Printf("the value and data type is %v, %T\n", d, d)

}
