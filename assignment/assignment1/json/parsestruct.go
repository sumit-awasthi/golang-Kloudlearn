 // concept of parsing json to struct 
package main 

import ( 
	"encoding/json"
	"fmt"
) 
type Country struct { 
	Name	 string
	Capital string
	Continent string
} 
func main() { 
	var country1 Country 
	Data := []byte(`{ 
		"Name": "India", 
		"Capital": "New Delhi", 
		"Continent": "Asia"
	}`) 
	err := json.Unmarshal(Data, &country1) 

	if err != nil { 
		fmt.Println(err) 
	} 
	fmt.Println("Struct is:", country1) 
	fmt.Printf("%s's capital is %s and it is in %s.\n", country1.Name, 
								country1.Capital, country1.Continent) 
} 
