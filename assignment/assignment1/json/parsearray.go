// concept of parsing JSON to an array 
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
	var country []Country 
	Data := []byte(` 
	[ 
		{"Name": "India", "Capital": "New Delhi", "Continent": "Asia"}, 
		{"Name": "Germany", "Capital": "Berlin", "Continent": "Europe"}, 
		{"Name": "Greece", "Capital": "Athens", "Continent": "Europe"}, 
		{"Name": "Israel", "Capital": "Jerusalem", "Continent": "Asia"} 
	]`) 
	err := json.Unmarshal(Data, &country) 

	if err != nil { 
		fmt.Println(err) 
	} 
	for i := range country { 
		fmt.Println(country[i].Name + " - " + country[i].Capital + 
									" - " + country[i].Continent) 
	} 
} 
