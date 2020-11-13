// decoding using JSON 
package main 
import ( 
	"fmt"
	"encoding/json"
) 
type Human struct{ 
    Name string 
	Address string 
	Age int
} 
 
func main() { 
	
	var human1 Human 
	
	Data := []byte(`{ 
		"Name": "Sumit", 
		"Address": "Uttar Pradesh", 
		"Age": 21 
	}`) 
	
	err := json.Unmarshal(Data, &human1) 
	
	if err != nil { 
		fmt.Println(err) 
	} 
	// decoded data 
	fmt.Println("Struct is:", human1) 
	fmt.Printf("%s lives in %s.\n", human1.Name, human1.Address) 
	
	var human2 []Human 
		Data2 := []byte(` 
	[ 
		{"Name": "Vani", "Address": "Delhi", "Age": 21}, 
		{"Name": "Rashi", "Address": "Noida", "Age": 24}, 
		{"Name": "Rahul", "Address": "Pune", "Age": 25} 
	]`) 
	err2 := json.Unmarshal(Data2, &human2) 
	
	if err2 != nil { 
			fmt.Println(err2) 
		} 
	for i := range human2{ 
	
		fmt.Println(human2[i]) 
	} 
} 
