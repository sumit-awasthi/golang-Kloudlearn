//recover in golang
package main 

import ( 
	"fmt"
) 
}
func handlepanic() { 

	if a := recover(); a != nil { 	
		fmt.Println("RECOVER", a) 
	} 
} 

// Function 
func entry(lang *string, aname *string) { 

	defer handlepanic() 
]	if lang == nil { 	
		panic("Error: Language cannot be nil") 
	} 
	
	if aname == nil { 
		panic("Error: Author name cannot be nil") 
	} 
	fmt.Printf("Author Language: %s \n Author Name: %s\n", *lang, *aname) 
	fmt.Printf("Return successfully from the entry function") 
} 

func main() { 

	A_lang := "GO Language"
	entry(&A_lang, nil) 
	fmt.Printf("Return successfully from the main function") 
} 
