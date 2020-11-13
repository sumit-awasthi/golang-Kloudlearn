    //encoding using JSON 
package main 

import ( 
	"fmt"
	"encoding/json"
) 

type Human struct{  
	Name string 
	Age int
	Address string 
} 

func main() {  
	human1 := Human{"Sumit", 22, "Uttar Pradesh"} 
	
		human_enc, err := json.Marshal(human1) 
	
	if err != nil { 
		
		fmt.Println(err) 
	} 
	
	fmt.Println(string(human_enc)) 
	human2 := []Human{ 
		{Name: "Akshit", Age: 22, Address: "New Delhi"}, 
		{Name: "Priyanshi", Age: 20, Address: "Pune"}, 
		{Name: "Akash", Age: 24, Address: "Bangalore"}, 
	}  
	human2_enc, err := json.Marshal(human2) 
		
		if err != nil { 
			fmt.Println(err) 
		} 
		fmt.Println() 
		fmt.Println(string(human2_enc)) 
} 
