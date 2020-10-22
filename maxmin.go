package main

import "fmt"

func main() {
	var a int=100
	var b int =200
	var res int
	res=max(a,b)
	fmt.Printf("max value is %d\n",res)
}
func max(num1,num2 int)int{
	var result int 
	if(num1>num2){
		result=num1
	}else{
		result=num2
	}
	return result
}