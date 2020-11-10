//range exercise
package main

import "fmt"

func main() {
	var a []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//for i:=0;i< len(a);i++{  //normal printing of elements be like
	//fmt.Println(a[i])
	//	}
	for i, element := range a { //here i is the index and element is the stored item inside the brackets
		fmt.Printf("%d:%d\n", i, element) // to display content without the index value replace i with _(for _,elements:=range a     fmt.printf("%d\n",elements))
	}
}
