//value return
package main

import (
	"errors"
	"fmt"
)

func boom() error {
	return errors.New("error")
}

func main() {
	if err := boom(); err != nil {
		fmt.Println("An error occurred:", err)
		return
	}
	fmt.Println("go away!")
}
