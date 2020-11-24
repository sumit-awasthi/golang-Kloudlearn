//string search using sort per defined function
package main 

import ( 
	"fmt"
	"sort"
) 
func main() { 
	slice_1 := []string{"C", "Go", "Java", "C#", "Ruby"} 
	var f1, f2, f3 string 
	f1 = "Go"
	f2 = "C"
	f3 = "Java"
 
	sort.Strings(slice_1) 
	fmt.Println("Slice 1: ", slice_1) 
	res1 := sort.SearchStrings(slice_1, f1) 
	res2 := sort.SearchStrings(slice_1, f2)
	res3 := sort.SearchStrings(slice_1, f3)
	
	// Displaying the results 
	fmt.Println("Result 1: ", res1) 
	fmt.Println("Result 2: ", res2) 
	fmt.Println("Result 3: ", res3) 

} 
