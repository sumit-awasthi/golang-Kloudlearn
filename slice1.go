package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("hi", s)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("show:", s[2])
	fmt.Println("len:", len(s))
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("added new elements:", s)
}
