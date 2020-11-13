//password matching using regular expression
package main

import (
	"fmt"
	"regexp"
)

func main() {
	var pass string
	regex := "^[a-zA-Z0-9]{10}$"
	fmt.Println("Enter your password(must have atleast 10 character(without special symbols)): ")
	fmt.Scanln(&pass)
	match1, _ := regexp.MatchString(regex, pass)
	// Test the result.
	if match1 {
		fmt.Println("Valid Password")
	} else {
		fmt.Println("Please enter password according to requirement")
	}
}
