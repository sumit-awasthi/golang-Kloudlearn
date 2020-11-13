//basic stucture
package main 

import "fmt"
type Address struct { 
	Name string 
	city string 
	Pincode int
} 

func main() { 
	var a Address 
	fmt.Println(a) 
	a1 := Address{"Sumit", "Lucknow", 226012} 
	fmt.Println("Address1: ", a1) 
	
    a2 := Address{Name: "Tonny ", city: "Ballia", 
								Pincode: 277001} 

	fmt.Println("Address2: ", a2) 
} 
