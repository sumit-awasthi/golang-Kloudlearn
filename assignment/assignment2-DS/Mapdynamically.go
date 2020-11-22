//maps Dynamically
package main
import ("fmt"
)
func main() {
	// Create map of string slices.
	m := map[string][]string{
		"Shirt":   {"Maroon", "Blue"},
		"Jeans":   {"black", "Blue"},
		"Tshirts": {"Black", "White"},
	}
	// Add new items .
	m["Tshirts"] = []string{"orange", "red"}
	// Print slice at key.
	fmt.Println(m["Shirt"])
	fmt.Println(m["Jeans"])
	fmt.Println(m["Tshirts"])
	
	//iterate full map
	for i, value := range m {
		fmt.Println(i, ":", value)
	}
	// Loop over string slice at key.
	for i := range m["Shirt"] {
		fmt.Println("Shirts:-",i, m["Shirt"][i])
	}
	for i := range m["Jeans"] {
		fmt.Println("Jeans:-",i, m["Jeans"][i])
	}
	for i := range m["Tshirts"] {
		fmt.Println("Tshirts:-",i, m["Tshirts"][i])
	}
}
